# Known obfuscated flag
obfuscatedFlag = [0x56, 0x2E, 0x1C, 0x7D, 0xB3, 0xA5, 0xDD, 0x92, 0x4A, 0x33]

# Try all possible 1-byte keys for each byte
for key in range(0, 256):
    flag = ''.join(chr(obfuscatedFlag[i] ^ key) for i in range(len(obfuscatedFlag)))
    print(f"Trying key {key}: {flag}")

