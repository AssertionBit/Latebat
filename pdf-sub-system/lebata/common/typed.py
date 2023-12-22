__all__ = (
    "replace_person_address",
    "replace_person_name",
    "replace_personal_data",
)


from json import load as _json_load


def replace_person_address(text: str) -> None:
    ...


def replace_person_name(text: str) -> None:
    ...


def replace_personal_data(text: str) -> str:
    print(text.find("Типовая форма"))
    text.replace("Типовая форма", "meow")
    replace_person_name(text)
    replace_person_address(text)
    return text

