Plane = {}
Targets = {}
Nums = []

if __name__ == "__main__":
    file = open('2023-day3.txt', 'r')
    lines = file.readlines()

    SpecialAll = ["!","@","#","$","%","^","&","*","/","+","=","-"]
    Special = ["*"]
    Numbers = ["0","1","2","3","4","5","6","7","8","9"]

    x = 0
    y = 0

    for index, line in enumerate(lines):
        for element in line:
            Plane[str(x)+":"+str(y)] = element.strip()
            x += 1
        x = 0
        y += -1

    #Get numbers targets
    hold = ""
    vhold = ""
    for index, value in Plane.items():
        if value in Numbers:
            hold += index + ";"
            vhold += value
        elif value not in Numbers:
            if len(hold) > 0:
                Targets[hold] = vhold
            hold = ""
            vhold = ""
    GearCheck = {}
    Gears = []
    for index, value in Targets.items():
        a = index.split(";")
        hasBeenAdded = False
        for b in a:
            c = b.split(":")
            if len(c) > 1:
                for i in range(int(c[0]) - 1, int(c[0]) + 2):
                    for j in range(int(c[1]) - 1, int(c[1]) + 2):
                        if (i, j) != (int(c[0]), int(c[1])) and not hasBeenAdded:
                            try:
                                w = str(i)+":"+str(j)
                                if Plane[w] in Special:
                                    if w not in GearCheck:
                                        GearCheck[w] = value
                                    else:
                                        Gears.append(int(value) * int(GearCheck[w]))
                                if Plane[w] in SpecialAll:
                                    Nums.append(value)
                                    hasBeenAdded = True
                            except:
                                #do nothing
                                w = str(i)+":"+str(j)
    su = 0
    so = 0
    for p in Nums:
        su += int(p)
    print("Part 1: ", su)
    for r in Gears:
        so += int(r)
    print("Part 2: ", so)            
                        
