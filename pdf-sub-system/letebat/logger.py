from logging import DEBUG as _DEBUG_LVL
from logging import ERROR as _ERROR_LVL
from logging import INFO as _INFO_LVL
from logging import Logger as _Logger
from logging import StreamHandler as _StreamHandler
from os import environ as _environ
from sys import stderr as _stderr
from sys import stdout as _stdou


def getLogger() -> _Logger:
    logger = _Logger("lebata")

    def_stream = _StreamHandler(_stdou)
    def_stream.setLevel(_INFO_LVL)

    # err_stream = _StreamHandler(_stderr)
    # err_stream.setLevel(_ERROR_LVL)

    logger.addHandler(def_stream)
    # logger.addHandler(err_stream)
    logger.setLevel(_INFO_LVL)

    return logger
