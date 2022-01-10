from typing import Optional
from datetime import datetime

from fastapi import Depends
from power_cache import TTLCache
from sqlalchemy.ext.asyncio import AsyncSession

from pol import sa
from pol.models import Subject, Permission
from pol.depends import get_db
from pol.db.const import SubjectType
from pol.db.tables import ChiiSubject
from pol.curd.exceptions import NotFoundError


class SubjectNotFound(NotFoundError):
    pass


cache: TTLCache[int, Permission] = TTLCache(capacity=15, ttl=60)


class SubjectService:
    __slots__ = ("_db",)
    _db: AsyncSession
    NotFoundError = SubjectNotFound

    @classmethod
    async def new(cls, session: AsyncSession = Depends(get_db)):
        return cls(session)

    def __init__(self, db: AsyncSession):
        self._db = db

    async def get_by_id(self, subject_id: int) -> Subject:
        s: Optional[ChiiSubject] = await self._db.get(
            ChiiSubject, subject_id, options=[sa.joinedload(ChiiSubject.fields)]
        )

        if not s:
            raise self.NotFoundError

        date = None
        v = s.fields.field_date
        if isinstance(v, datetime):
            date = f"{v.year:04d}-{v.month:02d}-{v.day:02d}"

        return Subject(
            id=s.subject_id,
            type=SubjectType(s.subject_type_id),
            name=s.subject_name,
            name_cn=s.subject_name_cn,
            nsfw=bool(s.subject_nsfw),
            summary=s.field_summary,
            infobox=s.field_infobox,
            ban=s.subject_ban,
            date=date,
            image=s.subject_image,
            platform=s.subject_platform,
        )
