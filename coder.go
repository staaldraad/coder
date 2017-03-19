package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"net/url"
	"os"
)

func urlDecode(input string) {
	dec, err := url.QueryUnescape(input)
	if err != nil {
		fmt.Println("- URL Decode failed")
	} else {
		fmt.Printf("%s", dec)
	}
}

func urlEncode(input string) {
	enc := url.QueryEscape(input)
	fmt.Printf("%s", enc)
}

func htmlEncode(input string) {
	enc := html.EscapeString(input)
	fmt.Printf("%s", enc)
}

func htmlDecode(input string) {
	enc := html.UnescapeString(input)
	fmt.Printf("%s", enc)
}

func hexDecode(input string) {
	dec, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("- Hex Decode failed")
	} else {
		fmt.Printf("%s", dec)
	}
}

func hexEncode(input string, format int) {
	enc := hex.EncodeToString([]byte(input))
	if format > 0 {
		for i := 0; i < len(enc); i += 2 {
			switch format {
			case 1:
				fmt.Printf("0x%s,", enc[i:i+2])
			case 2:
				fmt.Printf("\\x%s", enc[i:i+2])
			}
		}
		fmt.Println()
	} else {
		fmt.Printf("%s", enc)
	}
}

func hexDumper(input string) {
	fmt.Println(hex.Dump([]byte(input)))
}

func base64Decode(input string,url bool) {
    var data []byte
    var err error
    if url == true {
	    data, err = base64.URLEncoding.DecodeString(input)
    } else {
	    data, err = base64.StdEncoding.DecodeString(input)
    }
	if err != nil {
		fmt.Println("- Error base64 Decoding", err)
	} else {
		fmt.Printf("%s", data)
	}
}

func base64Encode(input string,url bool) {
    var str string
    if url == true {
	    str = base64.URLEncoding.EncodeToString([]byte(input))
    } else {
	    str = base64.StdEncoding.EncodeToString([]byte(input))
    }
	fmt.Printf("%s", str)
}

func main() {
	urlFunc := flag.Bool("u", false, "Do URL encoding/decoding")
	htmlFunc := flag.Bool("ht", false, "Do HTML encoding/decoding")
	base64Func := flag.Bool("b", false, "Do Base64 encoding/decoding")
	hexFunc := flag.Bool("x", false, "Do Hex encoding/decoding")
	hexFormat := flag.Int("xf", 0, "Format Hex encoding as 0 - 00 (Default) 1 - 0x00 2 - \\x00")
	hexDump := flag.Bool("xd", false, "Do hex dump (like hexdump -c)")
	allFunc := flag.Bool("a", false, "Try all methods encoding/decoding")
	encodeFunc := flag.Bool("e", false, "Encode")
	decodeFunc := flag.Bool("d", true, "Decode")
	inputString := flag.String("i", "", "Input if not from stdin")

	flag.Parse()

	if *encodeFunc == true {
		*decodeFunc = false
	}

	if *inputString == "" && (*base64Func || *hexFunc || *hexDump || *urlFunc || *htmlFunc || *allFunc) {
		input, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			fmt.Println("- Error reading from stdin!")
			os.Exit(1)
		} else {
			*inputString = string(input)
		}
	}

	if *allFunc == true {
		*urlFunc = true
		*hexFunc = true
		*htmlFunc = true
		*base64Func = true
	}

	if *urlFunc == true && *base64Func == false {
		if *decodeFunc == true {
			urlDecode(*inputString)
		} else {
			urlEncode(*inputString)
		}
	}

	if *htmlFunc == true {
		if *decodeFunc == true {
			htmlDecode(*inputString)
		} else {
			htmlEncode(*inputString)
		}
	}

	if *base64Func == true {
		if *decodeFunc == true {
			base64Decode(*inputString,*urlFunc)
		} else {
			base64Encode(*inputString,*urlFunc)
		}

	}
	if *hexFunc == true {
		if *decodeFunc == true {
			hexDecode(*inputString)
		} else {
			hexEncode(*inputString, *hexFormat)
		}
	}
	if *hexDump == true {
		hexDumper(*inputString)
	}

	os.Exit(0)
}
