f = open("input.txt", "r").read()

f = f.splitlines()
total = 0
card = {}
for i in range(len(f)) : card[i+1] = 0

for r, row in enumerate(f) :
    next = 0
    numbers = row.split(":")[1]
    win = numbers.split("|")[0].split(" ")
    win = [x for x in win if x]
    mynos = numbers.split("|")[1].split(" ")
    mynos = [x for x in mynos if x]
    for no in win :
        if no in mynos : next += 1
    card[r + 1] += 1
    for no in range(card[r+1]) :
        for i in range(r + 2, r + next + 2) :
            card[i] += 1

for i in range(len(f)) : total += card[i + 1]
print(total)


