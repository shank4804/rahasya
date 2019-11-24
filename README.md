# rahasya
A simple Go binary used to encrypt / decrypt file using password.

## Usage:
### Help:
```
./rahasya -h
Usage of ./rahasya:
  -f string
        filepath to do the type operation
  -p string
        password to encrypt / decrypt
  -t string
        operation to encrypt / decrypt (enc | dec)
```

### Encryption:
```
./rahasya -f plaintext.txt -p hello -t enc
```

### Decryption:
```
./rahasya -f encryptedtext -p hello -t dec
```

## Requirements:

```
go version go1.12.1 windows/amd64
```
