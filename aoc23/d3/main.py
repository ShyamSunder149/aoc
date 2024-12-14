import re 
f = open("input.txt", "r").read()

f = f.split("\n")
c = 0
flen = len(f)-1
vlen = len(f[0])
for i in range(flen) :
    values = re.findall("\d+", f[i])
    for j in values :
        ok = False
        validxf = f[i].index(j)
        validxl = validxf + len(j) - 1
        
        # check first col 
        if validxf - 1 >= 0 :
            if f[i][validxf-1] != "." : ok = True
            if i - 1 >= 0 and f[i-1][validxf-1] != "." : ok = True
            if i + 1 < flen and f[i+1][validxf-1] != "." : ok = True 
        
        # check last col 
        if validxl + 1 < vlen :
            if f[i][validxl+1] != "." : ok = True 
            if i - 1 >= 0 and f[i-1][validxl+1] != "." : ok = True 
            if i + 1 < flen and f[i+1][validxl+1] != "." : ok = True 
        
        # check all cols top and bottom 
        for k in j : 
            idx = f[i].index(k)
            if i - 1 >= 0 and f[i-1][idx] != "." : ok = True 
            if i + 1 < flen and f[i+1][idx] != "." : ok = True 
        if ok : c += int(j)

print(c)
