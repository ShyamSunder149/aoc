f = open("input.txt", "r").read()

game_ip = [i.split(":")[1] for i in f.split("\n") if i != ""]
ans = 0
for idx, i in enumerate(game_ip) :
    r, g, b = 0, 0, 0
    turns = i.split(";")
    for j in turns : 
        col = j.split(" ")
        for k in range(0, len(col), 2):
            if col[k].startswith("red") : r = max(int(col[k-1]), r)
            elif col[k].startswith("blue") : b = max(int(col[k-1]), b)
            elif col[k].startswith("green") : g = max(int(col[k-1]), g)
    ans += r * g * b 
print(ans)



