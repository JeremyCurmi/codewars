def parse(data: str) -> list:
    result: list = []
    value = 0

    for letter in data:
        if letter == "i":
            value += 1
        elif letter == "d":
            value -= 1
        elif letter == "s":
            value *= value
        elif letter == "o":
            result.append(value)

    return result

def main():
    tests = ["iiisdoso", "ooo", "ioioio", "codewars", "isoisoiso", "ssdizbsnsdkkincuybidds"]
    expected = [[8,64], [0,0,0], [1,2,3], [0], [1,4,25], None] 
    for i, par in enumerate(tests):
        got = parse(par)
        print(got, expected[i])
        print(got == expected[i])

if __name__ == "__main__":
    main()