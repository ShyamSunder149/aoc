
f = open("input.txt", "r").read()

ans = 0
f = f.replace("\n", "")
for i in f.split(",") :
    cv = 0
    for j in i :
        cv += ord(j)
        cv *= 17
        cv = cv %256
    ans += cv 
print(ans)
