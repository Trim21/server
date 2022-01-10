import enum
from typing import Optional
from datetime import datetime, timedelta

from pydantic import BaseModel

from pol.db.const import SubjectType

__all__ = ["PublicUser", "User", "Permission", "Creator", "Subject", "UserGroup"]


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
    user_group: UserGroup = UserGroup.normal_user
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


class Subject(BaseModel):
    id: int
    type: SubjectType
    infobox: str
    name: str
    name_cn: str
    summary: str
    nsfw: bool

    date: Optional[str]  # air date in `YYYY-MM-DD` format
    platform: int  # TV, Web, 欧美剧, PS4...
    image: str

    ban: int

    # volumes: int = Field(description="书籍条目的册数，由旧服务端从wiki中解析")
    # eps: int = Field(description="由旧服务端从wiki中解析，对于书籍条目为`话数`")
    # total_episodes: int = Field(description="数据库中的章节数量")

    @property
    def banned(self) -> bool:
        """redirected, not visible"""
        return self.ban == 1

    @property
    def locked(self) -> bool:
        """visible but not editable"""
        return self.ban == 2

    class Config:
        use_enum_values = True
