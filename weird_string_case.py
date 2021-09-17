def to_weird_case(string: str) -> str:
    cap = True
    word = ""
    
    for letter in string:

        print(letter, cap, word)
        if letter == " ":
            word += letter
            cap = True
        elif cap:
            word += letter.upper()
            cap = False
        else:
            word += letter.lower()
            cap = True

    return word


def main():
    string = "This is a test"
    expected = "ThIs Is A TeSt"
    got = to_weird_case(string)
    print(got == expected)
    print(got, expected)

if __name__ == "__main__":
    main()