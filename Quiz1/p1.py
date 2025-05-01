# Problem 1 

# (a) Use frequency analysis to attempt to recover the original plaintext, ASCII 32-126
frequency_order = "EARIOTNSLCUDPMHGBFYWKVXZJQ"

cipher = open("ciphertext.txt", "r").read()

count_freq = [0] * 127
for c in cipher:
    count_freq[ord(c)] += 1

total = sum(count_freq)

index_freq = [(i, count_freq[i]) for i in range(127)]
index_freq.sort(key=lambda x: x[1], reverse=True)

# (b) 
pairs = []
for a in range(0, 95):
    for b in range(0, 95):
        if (a * ord(' ') + b) % 95 + 32 == ord('3') and (a * ord('e') + b) % 95 + 32 == index_freq[1][0] :
            pairs.append((a, b))

print(pairs)

mapping = {}
for i in range(32, 127):
    mapping[(pairs[0][0] * i + pairs[0][1]) % 95 + 32] = i

visit = [False] * 127
for c in cipher:
    visit[ord(c)] = True

for i in range(32, 127):
    if visit[i] == True:
        print(chr(i), chr(mapping[i]))


#print(decryption)

