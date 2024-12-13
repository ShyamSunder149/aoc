f = open("input.txt", "r").read()

# implementation 

f = f.split("\n")
v = 0
print(f)
for i in f :
    number_indexes = [j for j in i if j.isdigit()]
    if number_indexes != [] : v += int(number_indexes[0] + number_indexes[-1])
print(v)
