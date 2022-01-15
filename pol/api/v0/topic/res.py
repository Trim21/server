import datetime
from typing import List

from pydantic import Field, BaseModel

from pol.api.v0.models.creator import Creator


class Topic(BaseModel):
    id: int
    title: str
    creator: Creator
    updated_at: datetime.datetime
    created_at: datetime.datetime
    reply_count: int = Field(description="回复数量，包含二级回复")


class BaseReply(BaseModel):
    id: int
    creator: Creator
    content: str
    created_at: datetime.datetime


class SubReply(BaseReply):
    parent: int


class Reply(BaseReply):
    replies: List[SubReply] = []


class TopicDetail(Topic):
    replies: List[Reply] = []
