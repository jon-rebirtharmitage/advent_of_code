if __name__ == "__main__":
    file = open('2023-day1.txt', 'r')
    lines = file.readlines()
    l = ["0","one","two","three","four","five","six","seven","eight","nine","1","2","3","4","5","6","7","8","9"]

    x = 0
    y = 1000
    z = 0
    a = 1000
    q = ""
    s = ""
    t = 0
    e = 0

    for index, line in enumerate(lines):
        for element in l:
            if(line.find(element) >= 0):
                x = line.find(element)
                if x < y:
                    y = x
                    if(len(element) > 1):
                        q = l.index(element)
                    else:
                        q = element
            if(line[::-1].find(element[::-1]) >= 0):
                z = line[::-1].find(element[::-1])
                if z < a:
                    a = z
                    if(len(element) > 1):
                        s = l.index(element)
                    else:
                        s = element
        w = str(q) + str(s)
        e = e + int(w)
        x = 0
        y = 1000
        z = 0
        a = 1000
    print(e)
