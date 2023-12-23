from click import argument
from click import command
from click import group
from click import option

from .pdf import anonimize_pdf


@group
def main() -> int:
    return 0


@command(name="process")
@argument("file")
def process(file: str) -> None:
    if file.endswith(".pdf"):
        try:
            anonimize_pdf(file)
        except OSError:
            return


main.add_command(process)


if __name__ == "__main__":
    exit(main())

