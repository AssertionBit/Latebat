__all__ = (
    "PdfPage",
    "PdfImage",
    "PdfText",
    "PdfDocument"
)

from abc import ABC as _ABC
from abc import abstractmethod as _abstractmethod
from dataclasses import dataclass as _structure
from typing import List as _List
from typing import Union as _Union

from pypdf import PageObject as _PdfPageObject
from pypdf.generic import TextStringObject as _TextObject
from reportlab.lib.units import cm as _cm
from reportlab.pdfgen.canvas import Canvas as _PdfCanvas


class Offset:
    start_x: int
    start_y: int
    max_x:   int
    max_y:   int


class PdfObject(_ABC):
    @_abstractmethod
    def to_pdf_object(self,
                      canvas: _PdfCanvas,
                      offset_y: int) -> None:
        ...


class PdfText(PdfObject):
    __content__: str

    def __init__(self, text: str) -> None:
        self.__content__ = text

    def to_pdf_object(self, canvas: _PdfCanvas, offset_y: int) -> None:
        ...


class PdfImage(PdfObject):
    __img__: str

    def __init__(self, img_path: str) -> None:
        self.__img__ = img_path

    def to_pdf_object(self, canvas: _PdfCanvas, offset_y: int) -> None:
        ...


class PdfPage(object):
    """
    @brief Page information object
    @field __page_number__ Number of page
    @field __page_text__ Text content of page
    """

    __page_number__: int = -1
    __page_objects__: _List[_Union[PdfText, PdfImage]] = []

    def __init__(self, 
                 number: int,
                 text: str) -> None:
        """
        @brief Constructor for page information
        @param text Raw text encoded with utf-8
        """

        self.__page_text__ = text
        self.__page_number__ = number

    def to_pdf_page(self, canvas: _PdfCanvas) -> None:
        """
        @brief Draw content of this file to pdf canvas
        @param canvas New PDF document to work with
        @return Insertable and processable PDF page
        """

        for object in self.__page_objects__:
            object.to_pdf_object(canvas)


class PdfDocument(object):
    """
    @brief Representation of PDF document which could be readed
        and converted to pdf and back
    """

    __doc_name__: str
    __doc_path__: str
    __doc_pages__: _List[PdfPage]

    def __init__(self, path: str) -> None:
        ...

    def build(self, out: str) -> None:
        """
        @brief Compiles back object into PDF
        """

    def __add__(self, o: PdfPage) -> None:
        ...

    def __iter__(self) -> "PdfDocument":
        self.__iterration_count__ = 0
        return self

    def __next__(self) -> PdfPage:
        self.__iterration_count__ += 1

        if len(self.__doc_pages__) == self.__iterration_count__:
            raise StopIteration

        return self.__doc_pages__[self.__iterration_count__]

