f = open("input.txt", "r").read()

seeds, *f = f.split('\n\n')
seeds = list(map(int, seeds.split(':')[1].split()))

for row in f :
    ranges = []
    for line in row.splitlines()[1:] : ranges.append(list(map(int, line.split())))
    new = []
    for x in seeds : 
        for a, b, c in ranges : 
            if b <= x < b + c : 
                new.append(x - b + a)
                break
        else : 
            new.append(x)
    seeds = new

print(min(seeds))

