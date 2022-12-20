laterns = []
lat = []

if __name__ == "__main__":
    file = open('2021-day6.txt', 'r')
    lines = file.read().splitlines()

    for index, line in enumerate(lines):
        a = line.split(",")
        for b in a:
            laterns.append(int(b))
    i = 0
    while i < 9:
        lat.append(laterns.count(i))
        i += 1

    j = 0
    while j < 256:
        g = lat.pop(0)
        lat[6] = lat[6] + g
        lat.append(g)
        j += 1

    s = 0
    for n in lat:
        s += n

    print(s)