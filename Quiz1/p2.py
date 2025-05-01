def gcd(a, b):
    if a < b:
        a, b = b, a
    if b == 0: 
        return a
    return gcd(b, a % b)
# (b)
print("Problem b")
for i in range(1, 30):
    if gcd(i, 30) == 1:
        for j in range(1, 30):
            if (i * j) % 30 == 1:
                print(i, j)
                
# (c)
print("Problem c") 
for n in range(31, 100):
    for a in range(1, n):
        if gcd(a, n) != 1: 
            continue
        for b in range(0, n):
            if(a * 81 + b) % n != 48:
                continue
            if(a * 14 + b) % n != 91:
                continue
            if(a * 3 + b) % n != 72:
                continue
            print(a, b, n)

# (d)
print("Problem d")
a, b, n = 37, 58, 97
inv_a = 0
for c in range(1, n):
    if(a * c) % n == 1:
        print(c)
        inv_a = c

d = inv_a * b % n
d = (n - d) % n
print(d)


