from pol.db.const import SubjectType
from .user import User, Avatar, UserGroup, Permission, PublicUser
from .subject import Subject

__all__ = [
    "PublicUser",
    "User",
    "Permission",
    "Subject",
    "UserGroup",
    "PublicUser",
    "User",
    "Avatar",
    "Subject",
]

#
# class Subject(BaseModel):
#     id: int
#     type: SubjectType
#     infobox: str
#     name: str
#     name_cn: str
#     summary: str
#     nsfw: bool
#
#     date: Optional[str]  # air date in `YYYY-MM-DD` format
#     platform: int  # TV, Web, 欧美剧, PS4...
#     image: str
#
#     ban: int
#
#     # volumes: int = Field(description="书籍条目的册数，由旧服务端从wiki中解析")
#     # eps: int = Field(description="由旧服务端从wiki中解析，对于书籍条目为`话数`")
#     # total_episodes: int = Field(description="数据库中的章节数量")
#
#     @property
#     def banned(self) -> bool:
#         """redirected, not visible"""
#         return self.ban == 1
#
#     @property
#     def locked(self) -> bool:
#         """visible but not editable"""
#         return self.ban == 2
#
#     class Config:
#         use_enum_values = True
