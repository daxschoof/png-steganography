# PNG Steganography
### Encrypting a PNG with secret messages

## Setup
The first step to encoding a file is opening the file and making sure of two things:
- The file actually has data
- The first 8 bytes of the file are: [137 80 78 71 13 10 26 10], as these 8 bytes are always the first 8 of any PNG