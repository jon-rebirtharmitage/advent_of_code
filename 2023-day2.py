from dataclasses import dataclass

@dataclass
class Game():
    id: int
    grabs: list

colors = ["red", "green", "blue"]

if __name__ == "__main__":
    file = open('2023-day2.txt', 'r')
    lines = file.readlines()

    games = []
    r = 0
    g = 0
    b = 0
    for index, line in enumerate(lines):
        #Get the game id and set as the ID int
        x = line.split(":")
        x = x[0].split(" ")
        x = int(x[1])
        #Get the games and create the tuples for the grabs list
        y = line.split(":")
        y = y[1].split(";")
        w = []
        for i in y:
            a = i.split(",")
            r = 0
            g = 0
            b = 0
            for j in a:
                s = j.split(" ")
                #print(s[1], ":", s[2])
                if(s[2].strip() == colors[0]):
                    r = int(s[1])
                elif(s[2].strip()  == colors[1]):
                    g = int(s[1])
                elif(s[2].strip()  == colors[2]):
                    b = int(s[1])
            h = (r,g,b)
            w.append(h)
        f = Game(id=x, grabs=w)
        games.append(f)
    result = 0
    hasBeenAdded = []
    isImpossible = []
    for ga in games:
        z = ga.grabs
        for gb in z:
            if(gb[0] > 12 or gb[1] > 13 or gb[2] > 14):
                if ga.id not in isImpossible:
                    isImpossible.append(ga.id)
            else:
                if ga.id not in hasBeenAdded:
                    hasBeenAdded.append(ga.id)
    for a in hasBeenAdded:
        if a not in isImpossible:
            result = result + a
    print("Part 1: ", result)
    powers = 0
    for gc in games:
        z = gc.grabs
        minr = 0
        ming = 0
        minb = 0
        for gd in z:
            if(gd[0]>=minr):
                minr = gd[0]
            if(gd[1]>=ming):
                ming = gd[1]
            if(gd[2]>=minb):
                minb = gd[2]
        powers = powers + (minr * ming * minb)
    print("Part 2: ", powers)

