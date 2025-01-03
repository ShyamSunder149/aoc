import re 
f = open("input.txt", "r").read()

grid = f.splitlines()
total = 0

for r, row in enumerate(grid) :
    for c, ch in enumerate(row) :
        if ch == "*" :
            ps = set()
            for cr in [r-1, r, r+1] :
                for cc in [c-1, c, c+1] :
                    if cr < 0 or cr >= len(grid) or cc < 0 or cc >= len(grid[cr]) or not grid[cr][cc].isdigit() : continue
                    while cc > 0 and grid[cr][cc-1].isdigit() : cc -= 1 
                    ps.add((cr, cc))
            
            if len(ps) != 2 : continue 
            ns = []
            for cr, cc in ps : 
                s = ""
                while cc < len(grid[r]) and grid[cr][cc].isdigit() :
                    s += grid[cr][cc]
                    cc += 1 
                ns.append(int(s)) 
            total += ns[0] * ns[1]
print(total)

