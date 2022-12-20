laterns = []

if __name__ == "__main__":
    file = open('2021-day6.txt', 'r')
    lines = file.read().splitlines()

    for index, line in enumerate(lines):
        a = line.split(",")
        for b in a:
            laterns.append(int(b))
    i = 0
    while i < 256:
        q = 0
        while q < len(laterns):
            if laterns[q] == 0:
                laterns.append(9)
                laterns[q] = 6
            else:
                laterns[q] -= 1
            q += 1
        i += 1
    print(len(laterns))