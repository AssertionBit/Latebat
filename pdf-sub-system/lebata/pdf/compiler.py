from logging import Logger as _Logger

from pypdf import PdfWriter as _PdfWritter


def compile_pdf(logger: _Logger) -> None:
    """
    @brief Compile pdf back from source files
    @param logger Logger to build message
    """

    try:
        writter = _PdfWritter()

    except:
        ...

    finally:
        ...

