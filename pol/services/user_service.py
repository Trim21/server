from typing import Optional
from datetime import datetime

from loguru import logger
from fastapi import Depends
from power_cache import TTLCache
from sqlalchemy.ext.asyncio import AsyncSession

from pol import sa
from pol.models import User, Permission, PublicUser
from pol.depends import get_db, get_redis
from pol.db.tables import ChiiMember, ChiiUserGroup, ChiiOauthAccessToken
from pol.curd.exceptions import NotFoundError
from pol.redis.json_cache import JSONRedis


class UserNotFound(NotFoundError):
    pass


cache: TTLCache[int, Permission] = TTLCache(capacity=15, ttl=60)


class UserService:
    __slots__ = ("_db",)
    _db: AsyncSession
    NotFoundError = UserNotFound

    @classmethod
    async def new(
        cls,
        session: AsyncSession = Depends(get_db),
        redis: JSONRedis = Depends(get_redis),
    ):
        return cls(session, redis)

    def __init__(self, db: AsyncSession, redis: JSONRedis):
        self._db = db

    async def get_permission(self, group_id: int):
        """从数据库读取当前的权限规则，在app中缓存60s"""
        if permission := cache.get(group_id):
            return permission
        p: ChiiUserGroup = await self._db.get(ChiiUserGroup, group_id)
        permission = Permission.parse_obj(p.usr_grp_perm)
        cache.set(group_id, permission)
        return permission

    async def get_by_name(self, username: str) -> PublicUser:
        """return a public readable user with limited information"""
        u: Optional[ChiiMember] = await self._db.scalar(
            sa.get(ChiiMember, ChiiMember.username == username)
        )

        if not u:
            raise self.NotFoundError

        return PublicUser(
            id=u.uid,
            username=u.username,
            nickname=u.nickname,
        )

    async def get_by_access_token(self, access_token: str) -> User:
        """return a authorized user"""
        access: Optional[ChiiOauthAccessToken] = await self._db.scalar(
            sa.get(
                ChiiOauthAccessToken,
                ChiiOauthAccessToken.access_token == access_token,
                ChiiOauthAccessToken.expires > datetime.now(),
            )
        )

        if not access:
            raise self.NotFoundError

        member: ChiiMember = await self._db.get(ChiiMember, int(access.user_id))

        if not member:
            # 有access token又没有对应的user不太可能发生，如果发生的话打个 log 当作验证失败
            logger.error(
                "can't find user {user_id} for access token", user_id=access.user_id
            )
            raise self.NotFoundError

        return User(
            id=member.uid,
            group_id=member.groupid,
            username=member.username,
            nickname=member.nickname,
            registration_time=member.regdate,
            sign=member.sign,
            avatar=member.avatar,
            permission=await self.get_permission(member.groupid),
        )
