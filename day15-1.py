import re
from functools import partial
from math import inf
from multiprocessing import Pool


def parse_input(input_str):
    sensors = []
    file = open('day15.txt', 'r')
    lines = file.readlines()
    for index, line in enumerate(lines):
        #x, y, bx, by = map(int, re.findall(r'(-?\d+)', line))
        a = line.split(" ")
        b = a[2].split("=")
        c = b[1].replace(",","")
        d = a[3].split("=")
        e = d[1].replace(",","")
        e = e.replace(":", "")
        f = a[8].split("=")
        g = f[1].replace(",","")
        h = a[9].split("=")
        i = h[1].replace("/n", "")
        r = abs(int(c) - int(g)) + abs(int(e) - int(i))
        sensors.append((int(c), int(e), r))
    return sensors


def solve_first(sensors):
    # assuming no beacons or signal gaps on y = 2_000_000
    min_x, max_x = inf, -inf
    print(inf)
    for x, y, r in sensors:
        dy = abs(2000000 - y)
        if dy > r:
            continue
        dx = r - dy
        min_x = min(x - dx, min_x)
        max_x = max(x + dx, max_x)
    return max_x - min_x


def solve(sensors, y):
    ranges = []
    for sx, sy, r in sensors:
        dy = abs(y - sy)
        if dy > r:
            continue
        dx = r - dy
        ranges.append((sx - dx, sx + dx))
    ranges.sort()
    prev_x2 = ranges[0][1]
    for x1, x2 in ranges[1:]:
        if x1 > prev_x2:
            return prev_x2 + 1, y
        prev_x2 = max(x2, prev_x2)
    return False


def solve_second(sensors):
    with Pool() as p:
        results = p.map(partial(solve, sensors), range(4000000))
        x, y = next(filter(bool, results))
        return x * 4000000 + y


if __name__ == '__main__':
    sensors = parse_input("day15.txt")
    print(solve_first(sensors))
    print(solve_second(sensors))