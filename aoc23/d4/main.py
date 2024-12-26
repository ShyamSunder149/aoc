f = open("input.txt", "r").read()

f = f.splitlines()
total = 0

for row in f :
    points = 0
    numbers = row.split(":")[1]
    win = numbers.split("|")[0].split(" ")
    win = [x for x in win if x]
    mynos = numbers.split("|")[1].split(" ")
    mynos = [x for x in mynos if x]
    for no in win :
        if no in mynos : 
            if points == 0 : points = 1
            else : points *= 2
    total += points 

print(total)

