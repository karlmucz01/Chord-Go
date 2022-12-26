package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

const FT_SIZE = 16
const UINT64_MAX = 18446744073709551615

func main() {
	var id uint64 = get_hash(os.Args[1])
	fmt.Println("SHA1 Hash: ", id)
}

func get_hash(message string) uint64 {
	var full_hash [20]byte = sha1.Sum([]byte(message))            // creating the hash byte array from the passed string
	var truncated_hash string = hex.EncodeToString(full_hash[:8]) // getting first 8 bytes of a hash
	id, err := strconv.ParseUint(truncated_hash, 16, 64)
	if err != nil {
		fmt.Println("Error parsing node ID: ", err)
	}
	return id
}
