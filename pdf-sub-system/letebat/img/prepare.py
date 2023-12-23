__all__ = (
    "read_image",
    "process_image"
)

from typing import Optional as _Optional

from cv2 import COLOR_BGR2GRAY
from cv2 import cvtColor as _conver_color
from cv2 import imread as _cv_read
from cv2.typing import MatLike as _MatLike
from pytesseract import image_to_string

from .anonimize import process_passport as _passport_proc
from .anonimize import process_snils as _process_snils
from .type import ImageType as _Type


def is_passport(text: str) -> bool:
    ...  # РМВИЗЬЕСМЕСО<<МЕАРТЗАУ<МТКОЕАЕУТЗ<<<<<<<<<<
    # 601767612481$6403120м<<<<<<<9191211610069<46

    return False


def is_snils(text: str) -> bool:

    return True


def detect_type(text: str) -> _Type:
    for t, c in {_Type.passport: is_passport, _Type.snils: is_snils}.items():
        if c(text):
            return t

    return _Type.passport


def read_image(file: str) -> _MatLike:
    cv_f = _cv_read(file)

    return cv_f


def process_image(file: str, type: _Optional[_Type] = None) -> None:
    cv_f = read_image(file)

    rgb_cv_f = _conver_color(cv_f, COLOR_BGR2GRAY)
    text = image_to_string(rgb_cv_f, lang='rus')
    print(text)

    if type is None:
        type = detect_type(text)

    if type == _Type.passport:
        return _passport_proc(cv_f)

    if type == _Type.snils:
        return _process_snils(cv_f)
