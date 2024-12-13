
f = open("input.txt", "r").read()

f = f.split("\n")
time = f[0].split(" ")[1:]
distance = f[1].split(" ")[1:]
time = [int(i) for i in time if i != '']
distance = [int(i) for i in distance if i != '']

c = 1
for i in range(len(time)) :
    val = 0
    for j in range(time[i]+1) :
        if (time[i] - j) * j > distance[i] : val += 1
    c *= val
print(c)
