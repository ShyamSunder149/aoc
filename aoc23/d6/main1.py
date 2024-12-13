f = open("input.txt", "r").read()

f = f.split("\n")
time = f[0].split(" ")[1:]
distance = f[1].split(" ")[1:]
time = int(''.join(time))
distance = int(''.join(distance))

val = 0
print(time ,distance)
for j in range(time) :
    if (time - j) * j > distance : val += 1
print(val)


