from copy import deepcopy

if __name__ == "__main__":
    file = open('2021-day3.txt', 'r')
    lines = file.read().splitlines()
    gamma = ""
    epsilon = ""
    f = []
    z = []
    for index, line in enumerate(lines):
        i = 0
        subf = []
        while i < len(line):
            subf.append(line[i])
            i += 1
        f.append(subf)
    i = 0
    d = []
    zeta = deepcopy(f)
    aj = len(f[0])
    while i < aj:
        mj = len(f)
        j = 0
        subd = []
        while j < mj:
            subd.append(f[j].pop(0))
            j += 1
        #print(subd)
        d.append(subd)
        i += 1

    print(len(d[0]))
    i = 0
    while i < len(d):
        if d[i].count("0") > d[i].count("1"):
            gamma += "0"
            epsilon += "1"
        else:
            gamma += "1"
            epsilon += "0"
        i += 1

    i = 0
    sGamma = deepcopy(zeta)
    sEpsilon = deepcopy(zeta)
    i = 0
    tGamma = []
    tEpsilon = []
    while len(sGamma) != 1:
        ts = []
        for q in sGamma:
            ts.append(q[i])

        if ts.count("0") > ts.count("1"):
            tq = "0"
        elif ts.count("0") == ts.count("1"):
            tq = "1"
        else:
            tq = "1"
        
        for m in sGamma:
            if m[i] == tq:
                tGamma.append(m)
        sGamma = deepcopy(tGamma)
        tGamma = []
        i += 1

    tEpsilon = []
    i = 0
    while len(sEpsilon) != 1:
        ts = []
        for q in sEpsilon:
            ts.append(q[i])

        if ts.count("0") > ts.count("1"):
            tq = "1"
        elif ts.count("0") == ts.count("1"):
            tq = "0"
        else:
            tq = "0"
        
        for m in sEpsilon:
            if m[i] == tq:
                tEpsilon.append(m)
        sEpsilon = deepcopy(tEpsilon)
        tEpsilon = []
        i += 1

    print(sGamma)
    print(sEpsilon)
    fGamma = ""
    fEpsilon = ""
    for a in sGamma:
        for b in a:
            fGamma += b
    for a in sEpsilon:
        for b in a:
            fEpsilon += b
    print(gamma)
    print(epsilon)
    print(int(gamma, 2) * int(epsilon, 2))
    print(int(fGamma, 2) * int(fEpsilon, 2))



        