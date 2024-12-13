f = open("input.txt", "r").read()

# implementation 

f = f.split("\n")
dit = ["one","two","three","four","five","six","seven","eight","nine"]

v = 0
fv, lv = [], []
for i in f : 
    digits = []
    for j in range(len(i)) :
        for l, k in enumerate(dit) :
            if i[j:].startswith(k) :
                print(i[j:])
                digits.append(str(l+1))
        if i[j].isdigit() : digits.append(i[j])
    if digits != [] : v += int(digits[0] + digits[-1])
print(v)
