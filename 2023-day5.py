
targets = []
soil = []
fertilizer = []
water = []
light = []
temp = []
humidity = []
location = []

t = ""

def findSoil(c):
    current = c
    for j in soil:
        if int(j[1]) + (int(j[2].strip())-1) >= c and c >= int(j[1]):
            z = int(c - int(j[2].strip()))
            a = int(int(j[1]) - int(j[0]))
            current = int(j[0]) + z + a
    return current

def findFertilizer(c):
    current = c
    for j in fertilizer:
        if int(j[1]) + (int(j[2].strip())) >= c and c >= int(j[1]):
            current = (int(j[0]) - int(j[1])) + c
            return current
    return current

def findWater(c):
    current = c
    for j in water:
        if int(j[1]) + (int(j[2].strip())) >= c and c >= int(j[1]):
            current = (int(j[0]) - int(j[1])) + c
            return current
    return current

def findLight(c):
    current = c
    for j in light:
        if int(j[1]) + (int(j[2].strip())) >= c and c >= int(j[1]):
            current = (int(j[0]) - int(j[1])) + c
            return current
    return current

def findTemp(c):
    current = c
    for j in temp:
        if int(j[1]) + (int(j[2].strip())) >= c and c >= int(j[1]):
            current = (int(j[0]) - int(j[1])) + c
            return current
    return current

def findHumidity(c):
    current = c
    for j in humidity:
        if int(j[1]) + (int(j[2].strip())) >= c and c >= int(j[1]):
            current = (int(j[0]) - int(j[1])) + c
            return current
    return current

def locateMe(c):
    current = c
    for j in location:
        if int(j[1]) + (int(j[2].strip())) >= c and c >= int(j[1]):
            current = (int(j[0]) - int(j[1])) + c
            return current
    return current
    
def partTwo():
    current = 0
    solution = 0
    j = 0
    i = 0
    while j <= (len(targets)-2):
        print(targets[j], j, targets[j+1], j+1)
        i = int(targets[j])
        while i <= int(targets[j]) + int(targets[j+1]):
            #print(i)
            current = findSoil(int(i))
            #print("Soil : ", current)
            current = findFertilizer(current)
            #print("fertilizer", current)
            current = findWater(current)
            #print("water", current)
            current = findLight(current)
            #print("light", current)
            current = findTemp(current)
            #print("temp", current)
            current = findHumidity(current)
            #print("humidity", current)
            current = locateMe(current)
            #print("location", current)
            #print(" END OF ROUND ")
            if solution == 0:
                solution = int(current)
            elif int(current) < solution:
                solution = int(current)
                print(solution)
                i = int(targets[j]) + int(targets[j+1])
            i += 1
        j += 2
    return solution

def checkValue(a, v):
    current = v
    for j in a:
        if int(j[0]) + (int(j[2].strip())) >= v and v >= int(j[0]):
            current = (int(j[1]) - int(j[0])) + v
            return current
    return current

def checkSeed(a, v):
    j = 0
    r = False
    while j <= (len(a)-2):
        if v in range(int(a[j]), (int(a[j]) + int(a[j+1]))):
            return True
        j += 2
    return r

def partTwoRedux():
    i = 0
    solution = 0
    solved = False
    v = 0
    while not solved:
        while i < 1000000000000:
            start = i
            current = checkValue(location, start)
            current = checkValue(humidity, current)
            current = checkValue(temp, current)
            current = checkValue(light, current)
            current = checkValue(water, current)
            current = checkValue(fertilizer, current)
            current = checkValue(soil, current)
            if checkSeed(targets, current):
                solution = start
                solved = True
                i += 1000000000000
            i += 1
    return solution
    


def partOne():
    current = 0
    solution = 0
    j = 0
    for i in targets:
        #print(i)
        current = findSoil(int(i))
        #print("Soil : ", current)
        current = findFertilizer(current)
        #print("fertilizer", current)
        current = findWater(current)
        #print("water", current)
        current = findLight(current)
        #print("light", current)
        current = findTemp(current)
        #print("temp", current)
        current = findHumidity(current)
        #print("humidity", current)
        current = locateMe(current)
        #print("location", current)
        if solution == 0:
            solution = int(current)
        elif int(current) < solution:
            solution = int(current)
    return solution

if __name__ == "__main__":
    file = open('2023-day5-targets.txt', 'r')
    lines = file.readlines()

    for index, line in enumerate(lines):
        a = line.split(":")
        b = a[1].split(" ")
        for i in b:
            if len(i.strip()) > 0:
                targets.append(i.strip())
                


    file = open('2023-day5.txt', 'r')
    lines = file.readlines()

    for index, line in enumerate(lines):
        if line.strip() != "":
            if t == "soil":
                a = line.split(" ")
                soil.append(a)
            elif t == "fertilizer":
                a = line.split(" ")
                fertilizer.append(a)
            elif t == "water":
                a = line.split(" ")
                water.append(a)
            elif t == "light":
                a = line.split(" ")
                light.append(a)
            elif t == "temp":
                a = line.split(" ")
                temp.append(a)
            elif t == "humidity":
                a = line.split(" ")
                humidity.append(a)
            elif t == "location":
                a = line.split(" ")
                location.append(a)
            if line.strip() == "seed-to-soil map:":
                t = "soil" 
            if line.strip() == "soil-to-fertilizer map:":
                t = "fertilizer"
            if line.strip() == "fertilizer-to-water map:":
                t = "water"
            if line.strip() == "water-to-light map:":
                t = "light"
            if line.strip() == "light-to-temperature map:":
                t = "temp"
            if line.strip() == "temperature-to-humidity map:":
                t = "humidity"
            if line.strip() == "humidity-to-location map:":
                t = "location"
        else:
            t = ""
    s = partOne()
    print("Part 1 : ", s)
    d = partTwoRedux()
    print("Part 2 : ", d)