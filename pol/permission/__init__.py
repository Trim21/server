from typing import Optional

from pol.models import Permission


class Role:
    permission: Optional[Permission]

    def allow_nsfw(self) -> bool:
        """if this user can see nsfw contents"""
        raise NotImplementedError()

    def get_user_id(self) -> Optional[int]:
        raise NotImplementedError()
