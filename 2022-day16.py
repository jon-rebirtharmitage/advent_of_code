Valves = {}
Paths = {}
currentLocation = "AA"
countdown = 30
pressue = 0

def findPath(a, c):
    dest = 0
    while c > 0:
        for d in Paths[a]:
            
        print("Turn check: ",c)
        c = c-1
    return True


if __name__ == "__main__":
    file = open('2022-day16.txt', 'r')
    lines = file.readlines()

    for index, line in enumerate(lines):
        a = line.split(";")
        b = a[0].split("=")
        Valves[a[0][6:8]] = b[1]
        if "," in a[1]:
            c = a[1].split("valves")
            d = c[1].split(",")
            for i in range(len(d)):
                d[i] = d[i].strip()
            #d[len(d) - 1] = d[len(d) - 1].strip()
            Paths[a[0][6:8]] = d
        else:
            c = a[1].split("valve")
            d.append(c[1].strip())
            Paths[a[0][6:8]] = d
        d = []
    print(Valves)
    print(Paths)
    while countdown >= 0:
        m = findPath(currentLocation, countdown)
        print(m)
        countdown = countdown - 1