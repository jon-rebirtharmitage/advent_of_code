
vents = []
masterVents = []
overlap = {}

def checkOverlap():
    ct = 0
    for i in masterVents:
        if masterVents.count(i) > 1:
            ct += 1
            overlap[i] = 1
    print(ct/2)

if __name__ == "__main__":
    file = open('2021-day5.txt', 'r')
    lines = file.read().splitlines()

    for index, line in enumerate(lines):
        a = line.split(" -> ")
        b = a[0].split(",")
        c = a[1].split(",")
        xa = int(b[0])
        ya = int(b[1])
        xb = int(c[0])
        yb = int(c[1])
        if xa == xb:
            tvent = []
            #line vertical
            if ya < yb:
                s = ya
                while s <= yb:
                    tvent.append((xa,s))
                    s += 1
                vents.append(tvent)
            if ya > yb:
                s = yb
                while s <= ya:
                    tvent.append((xa,s))
                    s += 1
                vents.append(tvent)
        if ya == yb:
            tvent = []
            #line horizontal
            if xa < xb:
                s = xa
                while s <= xb:
                    tvent.append((s,ya))
                    s += 1
                vents.append(tvent)
            if xa > xb:
                s = xb
                while s <= xa:
                    tvent.append((s,ya))
                    s += 1
                vents.append(tvent)
        if ya < yb and xa < xb:
            tvent = []
            a = ya
            b = xa
            while a <= yb and b <= xb:
                tvent.append((b,a))
                b += 1
                a += 1
            vents.append(tvent)
        if ya > yb and xa > xb:
            tvent = []
            a = yb
            b = xb
            while a <= ya and b <= xa:
                tvent.append((b,a))
                b += 1
                a += 1
            vents.append(tvent)
        if ya > yb and xa < xb:
            tvent = []
            a = yb
            b = xb
            while a <= ya and b >= xa:
                tvent.append((b,a))
                b -= 1
                a += 1
            vents.append(tvent)
        if ya < yb and xa > xb:
            tvent = []
            a = yb
            b = xb
            while a >= ya and b <= xa:
                tvent.append((b,a))
                b += 1
                a -= 1
            vents.append(tvent)
    for i in vents:
        for j in i:
            masterVents.append(j)
    checkOverlap()
    print(len(overlap))
