
f = open("input.txt", "r").read()

game_ip = [i.split(":")[1] for i in f.split("\n") if i != ""]
ans = 0
for idx, i in enumerate(game_ip) :
    r, g, b = 0, 0, 0
    turns = i.split(";")
    for j in turns : 
        col = j.split(" ")
        for k in range(0, len(col), 2): 
            if col[k].startswith("red") : r += int(col[k-1])
            elif col[k].startswith("blue") : b += int(col[k-1])
            elif col[k].startswith("green") : g += int(col[k-1])
    if r <= 12 and g <= 13 and b <= 14 : ans += idx + 1 
print(ans)

