import enum
from typing import List, Tuple, Iterator

from fastapi import Depends
from sqlalchemy.ext.asyncio import AsyncSession

from pol import sa
from pol.models import Creator, UserGroup, Permission
from pol.depends import get_db
from pol.db.tables import ChiiSubjectTopic
from pol.curd.exceptions import NotFoundError
from pol.models.community import Topic, TopicDisplayType


class TopicNotFound(NotFoundError):
    pass


class TopicType(enum.Enum):
    group = "group"
    subject = "subject"


class TopicService:
    __slots__ = ("_db",)
    _db: AsyncSession
    NotFoundError = TopicNotFound

    @classmethod
    async def new(cls, session: AsyncSession = Depends(get_db)):
        return cls(session)

    def __init__(self, db: AsyncSession):
        self._db = db

    async def get(self, type: TopicType, id: int):
        pass

    async def list(
        self,
        type: TopicType,
        group_id: int,
        limit: int,
        offset: int,
        permission: Permission,
    ) -> Tuple[int, List[Topic]]:
        """
        caller should check existence of parent entity (subject, group)

        :param limit:
        :param offset:
        :param type: group type, `subject` or `group`
        :param group_id: subject_id or group_id.
        :param permission: user permission rule
        """
        if type == TopicType.subject:
            return await self._list_subject_topics(group_id, limit, offset, permission)
        elif type == TopicType.group:
            return await self._list_group_topics(group_id, limit, offset)
        raise ValueError(type)

    async def _list_subject_topics(
        self,
        subject_id: int,
        limit: int,
        offset: int,
        permission: Permission,
    ) -> Tuple[int, List[Topic]]:
        where = self._permission(permission)

        total = await self._db.scalar(
            sa.select(sa.count(1))
            .where(ChiiSubjectTopic.sbj_tpc_subject_id == subject_id, *where)
            .options(sa.joinedload(ChiiSubjectTopic.creator))
        )

        results: Iterator[ChiiSubjectTopic] = await self._db.scalars(
            sa.select(ChiiSubjectTopic)
            .where(ChiiSubjectTopic.sbj_tpc_subject_id == subject_id, *where)
            .limit(limit)
            .offset(offset)
            .order_by(ChiiSubjectTopic.sbj_tpc_lastpost.desc())
            .options(sa.joinedload(ChiiSubjectTopic.creator))
        )

        return total, [
            Topic(
                id=x.sbj_tpc_id,
                title=x.sbj_tpc_title,
                created_at=x.sbj_tpc_dateline,
                updated_at=x.sbj_tpc_lastpost,
                creator=Creator.parse_obj(x.creator.as_creator()),
                reply_count=x.sbj_tpc_replies,
            )
            for x in results
        ]

    @staticmethod
    def _permission(p: Permission):
        default = [ChiiSubjectTopic.sbj_tpc_display == TopicDisplayType.normal]
        if not p:
            return default
        if p.user_group in (UserGroup.bangumi_admin, UserGroup.admin):
            return []
        return default

    async def _list_group_topics(
        self,
        subject_id: int,
        limit: int,
        offset: int,
    ) -> Tuple[int, List[Topic]]:
        raise NotImplementedError()
