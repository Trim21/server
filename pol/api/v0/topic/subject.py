import datetime

from fastapi import Depends, APIRouter
from pydantic import BaseModel

from pol.router import ErrorCatchRoute
from pol.api.v0.utils import user_avatar
from pol.api.v0.models import Paged, Pager
from pol.api.v0.models.creator import Creator
from pol.services.topic_service import TopicType, TopicService

router = APIRouter(
    tags=["社区"],
    route_class=ErrorCatchRoute,
    redirect_slashes=False,
)


class Topic(BaseModel):
    id: int
    title: str
    creator: Creator
    updated_at: datetime.datetime
    created_at: datetime.datetime
    reply_count: int


@router.get(
    "/subjects/{subject_id}/topics",
    summary="获取条目讨论帖",
    response_model=Paged[Topic],
)
async def get_topics(
    subject_id: int,
    page: Pager = Depends(),
    service: TopicService = Depends(TopicService.new),
):
    total, data = await service.list(
        TopicType.subject, subject_id, limit=page.limit, offset=page.offset
    )

    for i in data:
        i.creator.avatar = user_avatar(i.creator.avatar)["large"]

    return {
        "total": total,
        **page.dict(),
        "data": data,
    }
