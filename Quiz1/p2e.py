def gcd(a, b):
    if a < b:
        a, b = b, a
    if b == 0:
        return a
    return gcd(b, a % b)

for n in range(31, 100):
    for a in range(1, n):
        if gcd(a, n) != 1:
            continue
        for b in range(0, n):
            if(a * 45 + b) % n != 23:
                continue
            if(a * 2 + b) % n != 39:
                continue
            d1, d2 = -1, -1
            for i in range(0, 10):
                y = 40 + i
                if (a * 12 + b) % n == y:
                    d1 = i
                x = 10 * i + 3
                if (a * x + b) % n == 72:
                    d2 = i
            if d1 != -1 and d2 != -1:  
                print(a, b, n, d1, d2)
                continue         