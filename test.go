package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//firststring := "0x000x000x340x410x340x350x360x370x380x39"
	firststring := "\\x00\\x00\\x34\\x41\\x34\\x35\\x36\\x37\\x38\\x39"
	dd, _ := StrToByte(firststring)
	//fmt.Println(firststring)
	//var buf string
	reader := bufio.NewReader(os.Stdin)
	buf, _ := reader.ReadString('\n')
	fmt.Printf("%s", hex.Dump(dd))
	fmt.Println(hex.EncodeToString([]byte("\x00\x41\x00\x31")))
	//fmt.Scan(&buf)
	//fmt.Printf("%s", hex.Dump([]byte("\x00\x41\x00\x31")))
	fmt.Printf("%s", hex.Dump([]byte(buf)))
	fmt.Println(hex.EncodeToString([]byte(buf)))
	//fmt.Println(parseBinToHex(""))
}
func StrToByte(firststring string) ([]byte, error) {
	var s string
	for {
		fmt.Println(firststring)
		if strings.LastIndex(firststring, "\\x") != -1 {
			firststring = strings.Trim(firststring, "\\x")
			s = s + firststring[:2]
			fmt.Println(firststring[:2])
			firststring = firststring[2:]
			fmt.Println(firststring)
		} else {
			break
		}
	}
	//fmt.Println(s)
	////string to []byte
	return hex.DecodeString(s)
}
func parseBinToHex(s string) string {
	ui, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%x", ui)
}
