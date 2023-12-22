from click import argument
from click import command
from click import group
from click import option

from .pdf import process_pdf_file


@group
def main() -> int:
    return 0


@command(name="process")
@argument("file")
def process(file: str) -> None:
    if file.endswith(".pdf"):
        try:
            process_pdf_file(file)
        except OSError:
            return


main.add_command(process)


if __name__ == "__main__":
    exit(main())

