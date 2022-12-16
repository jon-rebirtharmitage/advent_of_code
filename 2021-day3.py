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
    while i < len(d):
        j = 0
        if d[i].count("0") > d[i].count("1"):
            while j < len(sGamma) and len(sGamma) != 1:
                if sGamma[i][j]
                j+=1
        print(sGamma)
        i += 1
    print(sGamma)
    print(gamma)
    print(epsilon)
    print(int(gamma, 2) * int(epsilon, 2))



        