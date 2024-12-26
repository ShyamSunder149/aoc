f = open("input.txt", "r").read()

seeds, *f = f.split('\n\n')
seedranges = list(map(int, seeds.split(':')[1].split()))
seeds = []

for i in range(0, len(seedranges), 2) :
    seeds.append((seedranges[i], seedranges[i] + seedranges[i + 1] + 1))

for row in f :
    ranges = []
    for line in row.splitlines()[1:] : ranges.append(list(map(int, line.split())))
    new = []
    while len(seeds) > 0 : 
        s, e = seeds.pop()
        for a,b,c in ranges : 
            os = max(s, b)
            oe = min(e, b + c) 
            if os < oe : 
                new.append((os-b+a, oe-b+a))
                if os > s : seeds.append((s, os))
                if oe < e : seeds.append((oe, e))
                break
        else : new.append((s, e))
    seeds = new

print(min(seeds))


