if __name__ == "__main__":
    file = open('2021-day2.txt', 'r')
    lines = file.readlines()

    x = 0
    y = 0
    aim = 0
    for index, line in enumerate(lines):
        a = line.split(" ")
        if a[0] == "forward":
            b = a[1].replace("/n", "")
            x += int(b)
            y += (aim * int(b))
        if a[0] == "up":
            b = a[1].replace("/n", "")
            aim -= int(b)
            #y -= int(b)
        if a[0] == "down":
            b = a[1].replace("/n", "")
            aim += int(b)
            #y += int(b)
    
    print("Product of depth : ", (x * y))
