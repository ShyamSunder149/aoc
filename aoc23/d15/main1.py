f = open("input.txt", "r").read()

f = f.replace("\n", "")
dit = {}

f = [i for i in f.split(",") if i != ""]
for i in f :
    cv = 0
    lenn = 0
    if "-" in i : 
        lenn = len(i) - 1
    else : 
        lenn = len(i)-2
    for j in range(lenn) :
        cv += ord(i[j])
        cv *= 17
        cv = cv %256
    
    if "=" in i :
        if cv not in dit.keys() : 
            dit[cv] = {i[:-2] : i[-1]} 
        else : 
            dit[cv][i[:-2]] = i[-1]
    elif cv in dit.keys() and "-" in i and i[:-1] in dit[cv].keys():
        del dit[cv][i[:-1]]
ans = 0
for i in dit.keys() :
    for k, j in enumerate(dit[i].keys()) :
        ans += int(dit[i][j]) * (i+1) * (k+1)
print(ans)

