from click import argument, command, group, option

from .img import process_image


@group
def main() -> int:
    return 0


@command(name="process")
@argument("file")
def process(file: str) -> None:
    if file.endswith(".pdf"):
        try:
            # cont = extract_data_from_pdf(file)
            # cont = anonimize_data(cont)
            # build_pdf_from_data(f"{file}.test", cont)
            ...

        except OSError:
            return

    elif file.endswith(".jpg") or \
            file.endswith(".jpeg") or \
            file.endswith(".png"):
        return process_image(file)


main.add_command(process)


if __name__ == "__main__":
    exit(main())
