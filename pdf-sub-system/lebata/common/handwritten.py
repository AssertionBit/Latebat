__all__ = (
    "has_signature",
    "remove_signature",
    "remove_person_name",
    "remove_person_address"
)

from cv2 import imread as _cv_read
from pypdf import PageObject as _Page


def has_signature(page: _Page) -> bool:
    ...


def remove_signature(page: _Page) -> None:
    if not has_signature(page):
        return


def remove_person_name(page: _Page) -> None:
    ...


def remove_person_address(page: _Page) -> None:
    ...

