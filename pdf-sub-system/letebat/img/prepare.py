__all__ = (
    "read_image",
    "process_image"
)

from re import MULTILINE as _MULTILINE
from re import findall as _findall
from typing import Optional as _Optional

from cv2 import COLOR_BGR2GRAY
from cv2.dnn_superres import DnnSuperResImpl as _SuperResImpl
from cv2 import cvtColor as _conver_color
from cv2 import filter2D as _filter
from cv2 import imread as _cv_read
from cv2 import imshow
from cv2 import imwrite as _imwrite
from cv2 import waitKey as _wait
from cv2.typing import MatLike as _MatLike
from numpy import array as _np_array
from pytesseract import image_to_string

from .anonimize import process_passport as _passport_proc
from .anonimize import process_snils as _process_snils
from .type import ImageType as _Type

def is_passport(text: str) -> bool:
    markers = 0

    patterns = (
        r"отделом[\sа-яА-Я0-9]+",
        r"[0-9а-яА-Яa-zA-Z<\$]+$"
    )

    if text.count("паспорт выдан") != 0:
       markers += 2

    for pattern in patterns:
        print(pattern)
        if _findall(pattern, text, _MULTILINE) is not None:
            markers += 1

    return markers 


def is_snils(text: str):
    m = 0

    if text.count("страховое") != 0 or text.count("свидетельство") != 0:
        m += 10

    patterns = (
        r"([\-0-9]+)",
    )

    for pattern in patterns:
        if _findall(pattern, text, _MULTILINE) != 0:
            m += 1

    return m


def detect_type(text: str) -> _Type:
    results = {}

    for t, c in {_Type.snils: is_snils, _Type.passport: is_passport}.items():
        if (m := c(text)) != 0:
            print(f"{t}: {m}")
            results[m] = t

    max = -1
    max_i = 0
    for i, c in enumerate(results.keys()):
        if max < c:
            max = c
            max_i = i

    return results[max]


def read_image(file: str) -> _MatLike:
    cv_f = _cv_read(file)

    return cv_f


def process_image(file: str, type: _Optional[_Type] = None) -> None:
    cv_f = read_image(file)

    sr = _SuperResImpl.create()
    sr.readModel("letebat-data/FSRCNN_x3.pb")
    sr.setModel("fsrcnn", 3)

    cv_f = sr.upsample(cv_f)

    rgb_cv_f = _conver_color(cv_f, COLOR_BGR2GRAY)

    # Procss text
    text: str = image_to_string(rgb_cv_f, lang='rus')
    text = text.lower()
    print(text)

    imshow("", rgb_cv_f)
    _wait(0)

    if type is None:
        type = detect_type(text)
    print(type)

    if type == _Type.passport:
        rgb_cv_f = _passport_proc(rgb_cv_f)

    if type == _Type.snils:
        rgb_cv_f = _process_snils(rgb_cv_f)
    
    _imwrite(file, rgb_cv_f)
