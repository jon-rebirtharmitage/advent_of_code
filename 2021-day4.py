boards = {}
bingos = []
finalCall = 0
totalWinners = {}

def markBingos(m):
    z = 0
    a = 0
    while z < len(boards):
        a = 0
        while a < 5:
            t = []
            b = 0
            while b < 5:
                if boards[z][(b,a)] == m:
                    print("This is the check", m, " and the value : ", boards[z][(b,a)])
                    boards[z][(b,a)] = "w"
                b += 1          
            a += 1
        z += 1

def checkWins(beta):
    z = 0
    a = 0
    while z < len(boards):
        a = 0
        while a < 5:
            t = []
            b = 0
            while b < 5:
                if boards[z][(b,a)] == "w":
                    t.append(boards[z][(b,a)])
                b += 1
            if len(t) == 5:
                print("There is a winner : ", z)
                totalWinners[z] = "Winner"
                if len(totalWinners) == len(boards):
                    calcWinner(z, beta)
                    return True
            a += 1
        z += 1
    z = 0
    a = 0
    while z < len(boards):
        a = 0
        while a < 5:
            t = []
            b = 0
            while b < 5:
                if boards[z][(a,b)] == "w":
                    t.append(boards[z][(a,b)])
                b += 1
            if len(t) == 5:
                print("There is a winner : ", z)
                totalWinners[z] = "Winner"
                if len(totalWinners) == len(boards):
                    calcWinner(z, beta)
                    return True
            a += 1
        z += 1
    return False

def calcWinner(z, beta):
    q = 0
    a = 0
    t = []
    su = 0
    a = 0
    while a < 5:
        b = 0
        while b < 5:
            if boards[z][(a,b)] != "w":
                su += boards[z][(a,b)]
                t.append(boards[z][(a,b)])
            b += 1
        a += 1

    finalScore = su * beta
    print("Final Score : ", finalScore, su, beta)


if __name__ == "__main__":
    file = open('2021-day4.txt', 'r')
    lines = file.read().splitlines()


    i = 0
    j = 0
    k = 0
    tempBoard = []
    for index, line in enumerate(lines):
        j = 0
        cboard = {}
        if i == 0:
            a = line.split(",")
            for s in a:
                bingos.append(int(s))
            i += 1
        else:
            m = line.split(" ")
            for n in m:
                if n != "":
                    tempBoard.append(n)
        i += 1

    i = 0
    q = 0
    while i < len(tempBoard):
        cboard = {}
        j = 0
        k = 0
        while j < 5:
            k = 0
            while k < 5:
                cboard[k,j]=int(tempBoard[i])
                i += 1
                k += 1
            j += 1
        boards[q] = cboard
        q += 1

    r = 0
    alpha = False
    while not alpha:
        while r < len(bingos):
            markBingos(bingos[r])
            alpha = checkWins(bingos[r])
            if alpha:
                finalCall = bingos[r]
                break
            r += 1



    print("========= Boards ============")
    z = 0
    a = 0
    while z < len(boards):
        a = 0
        while a < 5:
            t = []
            b = 0
            while b < 5:
                t.append(boards[z][(b,a)])
                b += 1
            print(t)           
            a += 1
        print("========================")
        z += 1
    print("============End of Boards============")




