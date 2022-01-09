import enum
from typing import List, Tuple, Iterator

from fastapi import Depends
from sqlalchemy.ext.asyncio import AsyncSession

from pol import sa
from pol.models import Creator
from pol.depends import get_db
from pol.db.tables import ChiiSubjectTopic
from pol.curd.exceptions import NotFoundError
from pol.models.community import Topic


class UserNotFound(NotFoundError):
    pass


class Category(enum.Enum):
    public = 1
    logged = 2
    admin = 3


class TopicType(enum.Enum):
    group = "group"
    subject = "subject"


class TopicService:
    _db: AsyncSession
    NotFoundError = UserNotFound

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
    ) -> Tuple[int, List[Topic]]:
        """
        caller should check existence of parent entity (subject, group)

        :param limit:
        :param offset:
        :param type: group type, `subject` or `group`
        :param group_id: subject_id or group_id.
        """
        if type == TopicType.subject:
            return await self._list_subject_topics(group_id, limit, offset)
        elif type == TopicType.group:
            return await self._list_group_topics(group_id, limit, offset)
        raise ValueError(type)

    async def _list_subject_topics(
        self,
        subject_id: int,
        limit: int,
        offset: int,
    ) -> Tuple[int, List[Topic]]:
        total = await self._db.scalar(
            sa.select(sa.count(1))
            .where(ChiiSubjectTopic.sbj_tpc_subject_id == subject_id)
            .options(sa.joinedload(ChiiSubjectTopic.creator))
        )

        results: Iterator[ChiiSubjectTopic] = await self._db.scalars(
            sa.select(ChiiSubjectTopic)
            .where(ChiiSubjectTopic.sbj_tpc_subject_id == subject_id)
            .limit(limit)
            .offset(offset)
            .order_by(ChiiSubjectTopic.sbj_tpc_lastpost)
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

    async def _list_group_topics(
        self,
        subject_id: int,
        limit: int,
        offset: int,
    ) -> Tuple[int, List[Topic]]:
        raise NotImplementedError()
