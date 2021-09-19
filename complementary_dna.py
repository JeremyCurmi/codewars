def DNA_strand(dna: str) -> str:
    result: str = ""
    for letter in dna:
        if letter == 'A':
            result += 'T'
        elif letter == 'T':
            result += 'A'
        elif letter == 'C':
            result += 'G'
        elif letter == 'G':
            result += 'C'
    return result

def best_practice(dna):
    pairs = {'A':'T','T':'A','C':'G','G':'C'}
    return ''.join(pairs[letter] for letter in dna)


def main():
    test = {
        "AAAA":"TTTT",
        "ATTGC": "TAACG",
        "GTAT":"CATA",
    }

    for par, expected in test.items():
        got = DNA_strand(par)
        try:
            assert(got == expected)
        except AssertionError:
            print(got, expected)

    for par, expected in test.items():
        got = best_practice(par)
        try:
            assert(got == expected)
        except AssertionError:
            print(got, expected)
    print("Tested !")

if __name__ == "__main__":
    main()
