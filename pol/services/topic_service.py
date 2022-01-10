import enum
import datetime
from typing import List, Iterator

from fastapi import Depends
from pydantic import BaseModel
from sqlalchemy.ext.asyncio import AsyncSession

from pol import sa
from pol.models import Creator, UserGroup, Permission
from pol.depends import get_db
from pol.db.tables import ChiiSubjectPost, ChiiSubjectTopic
from pol.curd.exceptions import NotFoundError
from pol.models.community import Topic, TopicDisplayType


class TopicNotFound(NotFoundError):
    pass


class TopicType(enum.Enum):
    group = "group"
    subject = "subject"


class Reply(BaseModel):
    id: int
    creator: Creator
    parent: int
    content: str
    state: int
    created_at: datetime.datetime
    replies: list = []


class TopicDetail(BaseModel):
    id: int
    title: str
    creator: Creator
    created_at: datetime.datetime
    updated_at: datetime.datetime
    reply_count: int

    replies: List[Reply] = []


class TopicService:
    __slots__ = ("_db",)
    _db: AsyncSession
    NotFoundError = TopicNotFound

    @classmethod
    async def new(cls, session: AsyncSession = Depends(get_db)):
        return cls(session)

    def __init__(self, db: AsyncSession):
        self._db = db

    async def get_topic(
        self, type: TopicType, id: int, permission: Permission
    ) -> TopicDetail:
        if type == TopicType.subject:
            return await self._get_subject_topic(id, permission)
        # elif type == TopicType.group:
        #     return await self._list_group_topics(group_id, limit, offset)
        raise ValueError(type)

    async def _get_subject_topic(self, id: int, permission: Permission) -> TopicDetail:
        topic: ChiiSubjectTopic = await self._db.get(
            ChiiSubjectTopic,
            id,
            options=[
                sa.joinedload(ChiiSubjectTopic.creator),
                sa.joinedload(ChiiSubjectTopic.replies).joinedload(
                    ChiiSubjectPost.creator
                ),
            ],
        )

        data = TopicDetail(
            id=topic.sbj_tpc_id,
            title=topic.sbj_tpc_title,
            creator=topic.creator.as_creator(),
            created_at=topic.sbj_tpc_dateline,
            updated_at=topic.sbj_tpc_lastpost,
            reply_count=0,
            replies=[
                Reply(
                    id=x.sbj_pst_id,
                    creator=x.creator.as_creator(),
                    parent=x.sbj_pst_related,
                    content=x.sbj_pst_content,
                    created_at=x.sbj_pst_dateline,
                    state=x.sbj_pst_state,
                )
                for x in topic.replies[1:]
            ],
        )

        return data

    async def count(
        self,
        type: TopicType,
        group_id: int,
        permission: Permission,
    ) -> int:
        """
        caller should check existence of parent entity (subject, group)

        :param limit:
        :param offset:
        :param type: group type, `subject` or `group`
        :param group_id: subject_id or group_id.
        :param permission: user permission rule
        """
        if type == TopicType.subject:
            return await self._count_subject_topics(group_id, permission)
        raise ValueError(type)

    async def list(
        self,
        type: TopicType,
        group_id: int,
        limit: int,
        offset: int,
        permission: Permission,
    ) -> List[Topic]:
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

    async def _count_subject_topics(
        self,
        subject_id: int,
        permission: Permission,
    ) -> int:
        where = self._permission(permission)

        return await self._db.scalar(
            sa.select(sa.count(1))
            .where(ChiiSubjectTopic.sbj_tpc_subject_id == subject_id, *where)
            .options(sa.joinedload(ChiiSubjectTopic.creator))
        )

    async def _list_subject_topics(
        self,
        subject_id: int,
        limit: int,
        offset: int,
        permission: Permission,
    ) -> List[Topic]:
        where = self._permission(permission)
        results: Iterator[ChiiSubjectTopic] = await self._db.scalars(
            sa.select(ChiiSubjectTopic)
            .where(ChiiSubjectTopic.sbj_tpc_subject_id == subject_id, *where)
            .limit(limit)
            .offset(offset)
            .order_by(ChiiSubjectTopic.sbj_tpc_lastpost.desc())
            .options(sa.joinedload(ChiiSubjectTopic.creator))
        )

        return [
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
    ) -> List[Topic]:
        raise NotImplementedError()
