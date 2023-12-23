__all__ = (
    "PageInfo",
)

from dataclasses import dataclass as _structure

from pypdf import PageObject as _PdfPageObject
from pypdf.generic import TextStringObject as _TextObject


@_structure
class Metadata(object):
    height: int
    width:  int


class PageInfo(object):
    """
    @brief Page information object
    @field __page_number__ Number of page
    @field __page_text__ Text content of page
    """

    __page_metadata__: Metadata
    __page_number__: int
    __page_text__:   str

    def __init__(self, 
                 meta: Metadata, 
                 number: int,
                 text: str) -> None:
        """
        @brief Constructor for page information
        @param text Raw text encoded with utf-8
        """

        self.__page_text__ = text
        self.__page_number__ = number
        self.__page_metadata__ = meta

    def to_pdf_page(self) -> _PdfPageObject:
        """
        @brief Convert this object to native PyPDF object
        @return Insertable and processable PDF page
        """

        page = _PdfPageObject.create_blank_page(None, 
                                                self.__page_metadata__.width, 
                                                self.__page_metadata__.height)

        return page

