package gopwn

import (
	"encoding/hex"
	"os"
	"strings"
)


func FileExists(FilePath string) bool{
	data, err := os.Stat(FilePath)

	if  os.IsNotExist(err) {
		return false
	}
	// something exists, make sure its not a directory.
	return !data.IsDir()
}

//Fix this, it's really fucked.
func Pad(char string, n int) []byte{
	// convert to hex.
	padded := strings.Repeat(char, n)
	dst := make([]byte, len(padded))

	hex.Encode(dst, []byte(padded))
	return dst
}

func p64(){}
func p32(){}

// merge several slices
func merge(args ...interface{}){}

