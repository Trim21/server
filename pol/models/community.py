import datetime

from pydantic import BaseModel

from pol.models import Creator
from pol.db.const import IntEnum


class Topic(BaseModel):
    id: int
    title: str
    creator: Creator
    created_at: datetime.datetime
    updated_at: datetime.datetime
    reply_count: int


class TopicStateType(IntEnum):
    """
    topic和post共用同一套state定义
    column: {tbl}_tpc_state
    """

    normal = 0
    closed = 1
    reopen = 2
    pin = 3
    merge = 4
    silent = 5
    delete = 6  # used by post
    private = 7  # used by post


class TopicDisplayType(IntEnum):
    """
    column: {tbl}_tpc_display
    """

    all = -1
    ban = 0
    normal = 1
    review = 2
