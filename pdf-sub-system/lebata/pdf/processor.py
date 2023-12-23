from logging import Logger as _Logger
from os.path import isfile as _isfile

from pypdf import PdfReader as _PdfReader

from lebata.logger import getLogger as _getLogger

from . import extractor as _ex


def anonimize_text(cont: str) -> None:
    ...


def anonimize_images(cont) -> None:
    ...


def anonimize_pdf(file: str) -> None:
    logger = _getLogger()

    if not _isfile(file):
        logger.error(f"'File' is not file", file)
        return

    try:
        origin = open(file, "rb")
        reader = _PdfReader(origin)
        for i in range(len(reader.pages)):
            styles = _ex.extract_styles(logger, reader.pages[i])
            text = _ex.extract_text_data(logger, reader.pages[i])
            images = _ex.extract_img_data(logger, reader.pages[i])

    except Exception as e:
        logger.error(e)

    finally:
        logger.info("Processing finished")
        origin.close() # type: ignore

