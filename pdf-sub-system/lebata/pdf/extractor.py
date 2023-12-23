from os import makedirs as _mkdirs
from os.path import isdir as _isdir
from logging import Logger as _Logger
from shutil import rmtree as _rmtree

from pypdf import PageObject as _PageObject


def extract_text_data(logger: _Logger, page: _PageObject) -> None:
    logger.info("Extracting text data")
    logger.info(page.extract_text())
    logger.info("Text extracted")


def extract_img_data(logger: _Logger, page: _PageObject) -> None:
    logger.info("Extracting images")
    logger.info(page.images)
    logger.info("Images extracted")


def extract_styles(logger: _Logger, page: _PageObject) -> None:
    logger.info("Extracting styles")
    for pi_key, pi_val in page.items():
        logger.info(f"{pi_key}: {pi_val}")
    logger.info("Extracted styles")

