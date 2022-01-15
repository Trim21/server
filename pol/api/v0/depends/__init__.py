from fastapi import Path, Depends

from pol import res
from pol.config import CACHE_KEY_PREFIX
from pol.models import PublicUser
from pol.depends import get_redis
from pol.permission import Role
from pol.redis.json_cache import JSONRedis
from pol.api.v0.depends.auth import optional_user
from pol.services.user_service import UserService
from pol.services.subject_service import Subject, SubjectService


async def get_public_user(
    username: str,
    user_service: UserService = Depends(UserService.new),
    not_found: res.HTTPException = Depends(res.not_found_exception),
) -> PublicUser:
    """
    get the user for `username` like `/user/{username}/collections`.
    UID is not working.
    """
    try:
        return await user_service.get_by_name(username)
    except UserService.NotFoundError:
        raise not_found


async def _get_subject(
    not_found: res.HTTPException = Depends(res.not_found_exception),
    subject_id: int = Path(..., gt=0),
    service: SubjectService = Depends(SubjectService.new),
    redis: JSONRedis = Depends(get_redis),
) -> Subject:
    """get a basic subject without any check"""
    cache_key = CACHE_KEY_PREFIX + f"subject:{subject_id}"
    subject = await redis.get_with_model(cache_key, Subject)
    if not subject:
        try:
            subject = await service.get_by_id(subject_id)
        except SubjectService.NotFoundError:
            raise not_found
        await redis.set_json(cache_key, value=subject.dict(), ex=300)

    if not subject:
        raise not_found

    return subject


async def get_readable_subject(
    subject: Subject = Depends(_get_subject),
    not_found: res.HTTPException = Depends(res.not_found_exception),
    user: Role = Depends(optional_user),
) -> Subject:
    """make sure subject visible for user"""
    if subject.nsfw and not user.allow_nsfw():
        raise not_found

    return subject


async def get_subject(
    subject_id: int = Path(..., gt=0),
    subject_service: SubjectService = Depends(SubjectService.new),
    user: Role = Depends(optional_user),
    not_found: res.HTTPException = Depends(res.not_found_exception),
) -> Subject:
    """
    make sure current subject is visible for current user
    also omit merged subject
    """
    try:
        return await subject_service.get_by_id(
            subject_id, user.allow_nsfw(), include_redirect=False
        )
    except SubjectService.NotFoundError:
        raise not_found
