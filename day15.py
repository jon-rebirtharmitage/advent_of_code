import numpy as np
import scipy.spatial as spatial

def manhattanDistance(x, y, z, a):
    return abs(x-z) + abs(y-z)

if __name__ == "__main__":
    file = open('day15.txt', 'r')
    lines = file.readlines()

    gridSensor = dict()
    gridBeacon = dict()
    grid = dict()
    minX = 0
    minY = 0
    maxX = 0
    maxY = 0

    for index, line in enumerate(lines):
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
        j = (int(c),int(e))
        k = (int(g),int(i))
        man = manhattanDistance(int(c), int(e), int(g), int(i))
        if int(c) - man < minX:
            minX = int(c) - man
        if int(c) + man > maxX:
            maxX = int(c) + man
        if int(g) - man < minX:
            minX = int(g)
        if int(g) + man > maxX:
            maxX = int(g)
        if int(e) - man < minY:
            minY = int(e)  - man
        if int(e) + man > maxY:
            maxY = int(e) + man
        if int(i) < minY:
            minY = int(i) 
        if int(i) > maxY:
            maxY = int(i)
        gridSensor[j] = (int(g), int(i), "S")
        gridBeacon[k] = "B"
    # print(gridSensor)
    # print(gridBeacon)
    # print(minX, maxX)
    # print(minY,maxY)
    xDimension = ((abs(minX)+abs(maxX)))
    yDimension = ((abs(minY)+abs(maxY)))
    cx = minX
    cy = minY
    ct = minY
    cr = minX
    numArray = []
    while ct <= maxY:
        cr = minX
        while cr <= maxX:
            #if (cr,ct) in gridBeacon:
                #grid[(cr,ct)] = (cr, ct, "B")
            if (cr,ct) in gridSensor:
                grid[(cr,ct)] = (gridSensor[(cr,ct)][0], gridSensor[(cr, ct)][1], "S")
                numArray.append((cr, ct))
            #else:
                #grid[(cr,ct)] = (cr, ct, ".")
            
            cr += 1
        ct += 1
    count = 0
    points = np.array(numArray)

    for al in grid:
        point_tree = spatial.cKDTree(points)
        if grid[al][2] == "S":
            #print(al , grid[al])
            #find all coordinates within distance
            a = al[0]
            b = al[1]
            c = grid[al][0]
            d = grid[al][1]
            md = manhattanDistance(a, b, c, d)
            amax = (a + man, b)
            amin = (a - man, b)
            bmax = (a, b + man)
            bmin = (a, b - man)
            #point_tree = spatial.cKDTree([al[0], al[1]], md)
            # This finds the index of all points within distance 1 of [1.5,2.5].
            mapThis = point_tree.data[point_tree.query_ball_point([al[0], al[1]], md)]
            v = 0
            while v < len(mapThis):
                #print(int(mapThis[v][0]), int(mapThis[v][1]))
                popGrid = grid[int(mapThis[v][0]), int(mapThis[v][1])] 
                if popGrid[2] != "S" or popGrid[2] != "B":
                    grid[int(mapThis[v][0]), int(mapThis[v][1])] = (popGrid[0], popGrid[1], "#")
                v += 1
    final = 2000000
    finalCounter = 0
    finalStart = minX
    finalEnd = maxX
    while finalStart <= finalEnd:
        if grid[finalStart, final][2] == "#":
            finalCounter += 1
        finalStart += 1
            
    print(finalCounter)
        




    
        
        