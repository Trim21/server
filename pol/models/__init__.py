import enum
from datetime import datetime, timedelta

from pydantic import BaseModel

__all__ = ["PublicUser", "User", "Permission", "Creator"]


class Creator(BaseModel):
    id: int
    username: str
    nickname: str
    avatar: str
    sign: str


class PublicUser(BaseModel):
    id: int
    username: str
    nickname: str


class Permission(BaseModel):
    app_erase: int = 0
    manage_app: int = 0
    user_list: int = 0
    manage_user_group: int = 0
    manage_user_photo: int = 0
    manage_topic_state: int = 0
    manage_report: int = 0
    user_ban: int = 0
    manage_user: int = 0
    user_group: int = 0
    user_wiki_approve: int = 0
    subject_edit: int = 0
    subject_lock: int = 0
    subject_refresh: int = 0
    subject_related: int = 0
    subject_merge: int = 0
    subject_erase: int = 0
    subject_cover_lock: int = 0
    subject_cover_erase: int = 0
    mono_edit: int = 0
    mono_lock: int = 0
    mono_merge: int = 0
    mono_erase: int = 0
    ep_edit: int = 0
    ep_move: int = 0
    ep_merge: int = 0
    ep_lock: int = 0
    ep_erase: int = 0
    report: int = 0
    doujin_subject_erase: int = 0
    doujin_subject_lock: int = 0


class UserGroup(enum.IntEnum):
    admin = 1  # 管理员
    bangumi_admin = 2  # Bangumi 管理猿
    window_admin = 3  # 天窗管理猿
    quite_user = 4  # 禁言用户
    banned_user = 5  # 禁止访问用户
    character_admin = 8  # 人物管理猿
    wiki_admin = 9  # 维基条目管理猿
    normal_user = 10  # 用户
    wiki = 11  # 维基人


class User(BaseModel):
    id: int
    username: str
    nickname: str
    group_id: UserGroup
    registration_time: datetime
    sign: str
    avatar: str
    permission: Permission

    def allow_nsfw(self) -> bool:
        allow_date = self.registration_time + timedelta(days=60)
        return datetime.utcnow().astimezone() > allow_date

    def get_user_id(self) -> int:
        return self.id
