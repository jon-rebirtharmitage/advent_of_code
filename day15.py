import numpy as np
import scipy.spatial as spatial

def manhattanDistance(x, y, z, a):
    return abs(x-z) + abs(y-a)

if __name__ == "__main__":
    file = open('day15.txt', 'r')
    lines = file.readlines()
    final = 2000000
    gridSensor = dict()
    gridBeacon = dict()
    grid = dict()
    minX = 0
    minY = 0
    maxX = 0
    maxY = 0
    numArray = []
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
        if int(c) < minX:
            minX = int(c)
        if int(c) > maxX:
            maxX = int(c)
        if int(g) < minX:
            minX = int(g)
        if int(g) > maxX:
            maxX = int(g)
        if int(e) < minY:
            minY = int(e)
        if int(e) > maxY:
            maxY = int(e)
        if int(i) < minY:
            minY = int(i) 
        if int(i) > maxY:
            maxY = int(i)
        gridSensor[j] = (int(g), int(i), "S")
        #numArray.append((int(g), int(i)))
        #gridBeacon[k] = "B"
    # print(gridSensor)
    # print(gridBeacon)
    # print(minX, maxX)
    # print(minY,maxY)
    xDimension = ((abs(minX)+abs(maxX)))
    yDimension = ((abs(minY)+abs(maxY)))
    # cx = minX
    # cy = minY
    ct = final
    cr = minX

    while ct == final:
        cr = minX
        while cr <= maxX:
            #if (cr,ct) in gridBeacon:
                #grid[(cr,ct)] = (cr, ct, "B")
            # if (cr,ct) in gridSensor:
            #     grid[(cr,ct)] = (gridSensor[(cr,ct)][0], gridSensor[(cr, ct)][1], "S")
            #else:
                #grid[(cr,ct)] = (cr, ct, ".")
            numArray.append((cr, ct))
            cr += 1
        ct += 1

    count = 0
    points = np.array(numArray)
    finalSAdd = 0
    gridCorrect = dict()
    for al in gridSensor:
        point_tree = spatial.cKDTree(points)
        if gridSensor[al][2] == "S":
            #print(al , gridSensor[al])
            #find all coordinates within distance
            a = al[0]
            b = al[1]
            c = gridSensor[al][0]
            d = gridSensor[al][1]
            md = manhattanDistance(a, b, c, d)
            #print(md)
            amax = (a + man, b)
            amin = (a - man, b)
            bmax = (a, b + man)
            bmin = (a, b - man)
            #point_tree = spatial.cKDTree([al[0], al[1]], md)
            # This finds the index of all points within distance 1 of [1.5,2.5].
            mapThis = point_tree.data[point_tree.query_ball_point([al[0], al[1]], md-1)]
            #print(mapThis)
            v = 0
            #print(mapThis)
            while v < len(mapThis):
                if mapThis[v][1] == final:
                    if mapThis[v][0] < minX or mapThis[v][1] < minY or mapThis[v][0] > maxX or mapThis[v][1] > maxY:
                        print("Fuck You")
                    else:
                    #print("In the target row", mapThis[v][0], mapThis[v][1])
                        grid[(mapThis[v][0], mapThis[v][1])] = "#"
                        finalSAdd += 1
                #print(int(mapThis[v][0]), int(mapThis[v][1]))
                # popGrid = grid[int(mapThis[v][0]), int(mapThis[v][1])] 
                # if popGrid[2] != "S" or popGrid[2] != "B":
                #     grid[int(mapThis[v][0]), int(mapThis[v][1])] = (popGrid[0], popGrid[1], "#")
                v += 1
            # if al[1] == final:
            #     grid.pop[(al[0],al[1])]
            if gridSensor[al][1] == final:
                gridCorrect[(c,d)] = "B"
 
    finalCounter = 0
    finalStart = minX
    finalEnd = maxX

    finalCounter = yDimension - len(grid)
            
    print(finalSAdd, finalCounter)
    print((grid))
    print(len(gridCorrect))
    print(len(grid)-len(gridCorrect))
    #print(grid)
        




    
        
        