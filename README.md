Coder
----

A small go application to encode and decode various formats. Allows both STDIN and -i for specifying input.

Usage
----
Usage of ./coder:
  -a    Try all methods encoding/decoding
  -b    Do Base64 encoding/decoding
  -d    Decode (default) (default true)
  -e    Encode
  -i string
        Input if not from stdin
  -u    Do URL encoding/decoding
  -x    Do Hex encoding/decoding
  -xd
        Do hex dump (like hexdump -c)
  -xf int
        Format Hex encoding as 0 - 00 (Default)
        1 - 0x00
        2 - \x00

Examples
----
* URL encode from the STDIN
```
echo -n "' or '1'='1" | coder -u -e
%27+or+%271%27%3D%271
```
* URL decode from input string
```
coder -u -i '%27+or+%271%27%3D%271'
' or '1'='1
```

* HEX encode
```
echo -n "string" | coder -x -e
737472696e67
```

* HEX encode with python formatting

```
echo -n "string" | coder -x -e -xf 2
\x73\x74\x72\x69\x6e\x67
```

* HEX dump

```
 cat /tmp/file.bin | coder -xd

 00000210  0f 4c 0f 0f 0e d3 0e 9e  0e 6a 0e 36 0e 02 0d cd  |.L.......j.6....|
 00000220  0d 99 0d 5c 0d 20 0c e3  0c a7 0c 6b 00 00 00 00  |...\. .....k....|
 00000230  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
 00000240  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
 00000250  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
```

