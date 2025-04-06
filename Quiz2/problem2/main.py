# SHA-1

import hashlib
def encrypt_sha1(text):
    sha_signature = hashlib.sha1(text.encode()).hexdigest()
    return sha_signature

import requests
password_list_url = "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt"

response = requests.get(password_list_url)

if response.status_code != 200:
    print("Failed to read txt file.")
    exit(0)

password_list = response.text.split("\n")

# Easy password 

EasyHash = "884950a05fe822dddee8030304783e21cdc2b246"
Easy_Counter = 0

for password in password_list:
    Easy_Counter += 1
    if encrypt_sha1(password) == EasyHash:
        print("Hash: ", EasyHash)
        print("Password: ", password)
        print(f"Took {Easy_Counter} attempts to crack message\n")
        break

# Medium password 

MediumHash = "9b467cbabe4b44ce7f34332acc1aa7305d4ac2ba"
Medium_Counter = 0
for password in password_list:
    Medium_Counter += 1
    if encrypt_sha1(password) == MediumHash:
        print("Hash: ", MediumHash)
        print("Password: ", password)
        print(f"Took {Medium_Counter} attempts to crack message\n")
        break

# Leak password 

Leak_Counter = 0 

SaltHash = "dfc3e4f0b9b5fb047e9be9fb89016f290d2abb06"

for salt in password_list:
    Leak_Counter += 1
    if encrypt_sha1(salt) == SaltHash:
        Salt = salt
        break


LeakHash = "9d6b628c1f81b4795c0266c0f12123c1e09a7ad3"

for password in password_list:
    Leak_Counter += 1
    if encrypt_sha1(Salt + password) == LeakHash:
        print("Hash: ", LeakHash)
        print("Password: ", password)
        print(f"Took {Leak_Counter} attempts to crack message\n")
        break