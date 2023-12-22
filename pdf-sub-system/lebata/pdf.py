from io import BytesIO, StringIO
from os import makedirs as _mkdirs
from os.path import abspath as _abspath
from os.path import join as _path_join
from os.path import isfile as _isfile

from pypdf.generic import ContentStream as _ContentStream
from pypdf import PdfReader as _PdfReader
from pypdf import PdfWriter as _PdfWritter
from pypdf import PageObject as _PdfPage

from .common import typed as _tp
from .logger import getLogger as _get_logger


def process_pdf_file(file: str) -> None:
    logger = _get_logger()
    logger.info(f"Processing file {file}")

    if not _isfile(file):
        raise OSError("'file' is not a file", file)

    try:
        pdf_file_raw = open(file, "r+b")
        pdf_file = _PdfReader(pdf_file_raw)
        pdf_writter = _PdfWritter()

        if pdf_file.is_encrypted:
            logger.error("File is encrypted, using simple password")
            pdf_file.decrypt("")

        logger.info(f"Pages count: {len(pdf_file.pages)}")

        for i, page in enumerate(pdf_file.pages):
            logger.info(f"Processing page: {i}")
            if len(page.images) != 0:
                logger.info(f"Images detected: {len(page.images)}")

            text = page.extract_text()
            _tp.replace_personal_data(text)

            replaced_page = _PdfPage.create_blank_page(None, page.mediabox[2], page.mediabox[3])

            pdf_writter.add_page(replaced_page)

        logger.info("Processed file normally")
        with open(f"{file.replace('.pdf', '')}.result.pdf", "w+b") as file:
            pdf_writter.write(file)

    except StopIteration:
        ...

    finally:
        logger.info("Processing finished")
        pdf_file_raw.close() # type: ignore

