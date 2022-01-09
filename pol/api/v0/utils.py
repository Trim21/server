from typing import Dict, List, Optional

from pol.db.tables import ChiiPerson


def person_images(s: Optional[str]) -> Optional[Dict[str, str]]:
    if not s:
        return None

    return {
        "large": "https://lain.bgm.tv/pic/crt/l/" + s,
        "medium": "https://lain.bgm.tv/pic/crt/m/" + s,
        "small": "https://lain.bgm.tv/pic/crt/s/" + s,
        "grid": "https://lain.bgm.tv/pic/crt/g/" + s,
    }


def get_career(p: ChiiPerson) -> List[str]:
    s = []
    if p.prsn_producer:
        s.append("producer")
    if p.prsn_mangaka:
        s.append("mangaka")
    if p.prsn_artist:
        s.append("artist")
    if p.prsn_seiyu:
        s.append("seiyu")
    if p.prsn_writer:
        s.append("writer")
    if p.prsn_illustrator:
        s.append("illustrator")
    if p.prsn_actor:
        s.append("actor")
    return s


def short_description(s: str):
    return s[:80]


def user_avatar(s: str):
    if not s:
        return {
            "large": "https://lain.bgm.tv/pic/user/l/icon.jpg",
            "medium": "https://lain.bgm.tv/pic/user/m/icon.jpg",
            "small": "https://lain.bgm.tv/pic/user/s/icon.jpg",
            "grid": "https://lain.bgm.tv/pic/user/g/icon.jpg",
        }
    return {
        "large": "https://lain.bgm.tv/pic/user/l/" + s,
        "medium": "https://lain.bgm.tv/pic/user/m/" + s,
        "small": "https://lain.bgm.tv/pic/user/s/" + s,
        "grid": "https://lain.bgm.tv/pic/user/g/" + s,
    }
