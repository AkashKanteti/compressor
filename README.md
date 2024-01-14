# compressor

## Building A Huffman Encoder/Decoder

For example if we have the string aaabbc, it would normally take up 6 bytes, but if we assign each character a variable length code, with the most frequently occurring character has the shortest code we might give them the following codes:

a: 1
b: 01
c: 10

we could reduce the string aaabbc (six bytes) to 111010110 (nine bits).