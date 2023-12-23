__all__ = (
    "process_passport",
)

from cv2 import INTER_AREA as _AREA
from cv2 import imshow
from cv2 import rectangle as _rectangle
from cv2 import resize as _resize
from cv2 import waitKey
from cv2.typing import MatLike as _MatLike


def process_passport(img: _MatLike) -> None:
    """
    """

    img = _resize(img, (750, 1000), interpolation=_AREA)

    # Hidding picture
    _rectangle(
        img,
        (40, 600),
        (240, 850),
        (0, 0, 0,),
        -1,
    )

    # Hidding name, part one
    _rectangle(
        img,
        (300, 555),
        (650, 580),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (300, 590),
        (650, 620),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (300, 630),
        (650, 660),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (300, 670),
        (650, 700),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (270, 710),
        (350, 730),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (440, 710),
        (650, 730),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (310, 750),
        (650, 770),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (270, 790),
        (650, 810),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (270, 830),
        (650, 850),
        (0, 0, 0),
        -1
    )


def process_snils(img: _MatLike):
    img = _resize(img, (1000, 750), interpolation=_AREA)

    _rectangle(
        img,
        (170, 260),
        (600, 410),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (400, 400),
        (900, 470),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (100, 500),
        (700, 600),
        (0, 0, 0),
        -1
    )

    _rectangle(
        img,
        (100, 640),
        (400, 700),
        (0, 0, 0),
        -1
    )

    imshow("", img)
    waitKey(0)
