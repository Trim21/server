from typing import Dict

from fastapi import Path, Depends, APIRouter

from pol.models import Avatar, Subject
from pol.router import ErrorCatchRoute
from pol.permission import Role
from pol.api.v0.models import Paged, Pager
from pol.api.v0.depends import get_readable_subject
from pol.api.v0.topic.res import Topic, TopicDetail
from pol.api.v0.depends.auth import optional_user
from pol.services.topic_service import TopicType, TopicService

router = APIRouter(
    tags=["社区"],
    route_class=ErrorCatchRoute,
    redirect_slashes=False,
)


@router.get(
    "/subjects/{subject_id}/topics",
    summary="获取条目讨论帖",
    response_model=Paged[Topic],
)
async def get_topics(
    page: Pager = Depends(),
    user: Role = Depends(optional_user),
    service: TopicService = Depends(TopicService.new),
    subject: Subject = Depends(get_readable_subject),
):
    total = await service.count(
        TopicType.subject,
        subject.id,
        permission=user.permission,
    )

    page.check(total)

    data = await service.list(
        TopicType.subject,
        subject.id,
        limit=page.limit,
        offset=page.offset,
        permission=user.permission,
    )

    for i in data:
        i.creator.avatar = Avatar.from_db_record(i.creator.avatar).large

    return {
        "total": total,
        **page.dict(),
        "data": data,
    }


@router.get(
    "/subjects/{subject_id}/topics/{topic_id}",
    summary="获取条目讨论帖",
    response_model=TopicDetail,
)
async def get_topic(
    topic_id: int = Path(..., gt=0),
    user: Role = Depends(optional_user),
    service: TopicService = Depends(TopicService.new),
    _: Subject = Depends(get_readable_subject),
):
    await service.get_topic(
        TopicType.subject,
        topic_id,
        permission=user.permission,
    )

    data = await service.get_topic(
        TopicType.subject,
        topic_id,
        permission=user.permission,
    )

    for i in data.replies:
        i.creator.avatar = Avatar.from_db_record(i.creator.avatar).large

    data.creator.avatar = Avatar.from_db_record(data.creator.avatar).large

    c: Dict[int, dict] = {x.id: x.dict() for x in data.replies if x.parent == 0}
    for i in data.replies:
        if i.parent:
            c[i.parent]["replies"].append(i.dict())

    data = data.dict()
    data["replies"] = list(c.values())
    data["reply_count"] = len(data["replies"])
    return data
