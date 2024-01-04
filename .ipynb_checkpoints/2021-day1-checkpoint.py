if __name__ == "__main__":
    file = open('2021-day1.txt', 'r')
    lines = file.readlines()

    current = 0
    count = 0
    total = []
    tcount = 0
    tcurrent = 0
    for index, line in enumerate(lines):
        total.append(int(line))
        if current == 0:
            current = int(line)
            continue
        if int(line) > current:
            count += 1
        current = int(line)
    i = 0

    while i+2 < len(total):
        if tcurrent == 0:
            tcurrent = (total[i] + total[i+1] + total[i+2])
            continue
        if (total[i] + total[i+1] + total[i+2]) > tcurrent:
            tcount += 1
        tcurrent = (total[i] + total[i+1] + total[i+2])
        i+=1

    print("Count of depth increasing : ", count)
    print("Count of Triple Increase :", tcount)