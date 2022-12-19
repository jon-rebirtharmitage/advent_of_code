grid = {}

if __name__ == "__main__":
    file = open('2021-day25.txt', 'r')
    lines = file.readlines()


    i = 0
    j = 0
    maxX = 0
    maxY = 0
    for index, line in enumerate(lines):
        j = 0
        for n in line:
            if n == "\n":
                i += 1
            else:
                grid[(j, i)] = n
            if i == 0:
                maxX += 1
            j += 1
    maxY = i + 1
    maxX = maxX
    a = 0
    b = 0
    print("========= Grid ============")
    while a < maxY:
        t = []
        b = 0
        while b < maxX:
            t.append(grid[(b,a)])
            b += 1
        print(t)
        a += 1
    print("======================================")


    print(maxX, maxY)

    moving = 0
    currentMoving = 0
    moved = True
    steps = 0
    rmove = {} 
    dmove = {}

    while moved:
        steps += 1
        currentMoving = moving
        a = 0
        b = 0

        while a < maxY:
            b = 0
            while b < maxX:
                if grid[(b,a)] == ">":
                    d = a
                    if (b + 1) == (maxX):
                        c = 0
                    else:
                        c = b+1
                    if grid[c,d] == ".":
                        rmove[b,a] = (c,d) 
                        moving += 1
                b += 1  
            a += 1
        for k,v in rmove:
            grid[k,v] = "."
            grid[(rmove[k,v][0], rmove[k,v][1])] = ">"
        rmove.clear()
        print("Right Done")

        q = 0
        w = 0
        while q < maxX:
            w = 0
            while w < maxY:
                if grid[q,w] == "v":
                    #print("Down Move Considered : ", q,w)
                    e = q
                    if (w + 1) == (maxY):
                        r = 0
                    else:
                        r = w+1
                    if grid[e,r] == ".":
                        dmove[q,w] = (e,r)
                        moving += 1
                    
                w += 1
            q += 1
        for k,v in dmove:
            grid[k,v] = "."
            grid[(dmove[k,v][0], dmove[k,v][1])] = "v"
        dmove.clear()
        print("Down Done")

        if moving == currentMoving:
            moved = False
        
        a = 0
        b = 0
        # print("========= Step ", steps, " ============")
        # while a < maxY:
        #     t = []
        #     b = 0
        #     while b < maxX:
        #         t.append(grid[(b, a)])
        #         b += 1
        #     print(t)
        #     a += 1
        # print("======================================")


print("Total Steps to stop moving : ", steps)
    




