from enum import Enum as _Enum


class ImageType(_Enum):
    passport = "passport"
    snils = "snils"


def detect_type() -> ImageType:
    ...
