package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	// "os"
	"strconv"
)

// defining constants
const FT_SIZE = 16
const UINT64_MAX = 18446744073709551615

// definition for a chord node interface
type Chord_Node interface {
	find_successor(id uint64) *node
	closest_preceding_node(id uint64) *node
	create()
}

// definition for a chord node struct
type node struct {
	id           uint64
	successor    *node
	predecessor  *node
	finger_table map[uint64]*node
}

func main() {
	var self node
	self.id = get_hash("node1-self")
	self.successor = &self
	self.predecessor = &self
	fmt.Println(self)
}

func get_hash(message string) uint64 {
	var full_hash [20]byte = sha1.Sum([]byte(message))            // creating the hash byte array from the passed string
	var truncated_hash string = hex.EncodeToString(full_hash[:8]) // getting first 8 bytes of a hash
	fmt.Println(truncated_hash)
	id, err := strconv.ParseUint(truncated_hash, 16, 64)
	if err != nil {
		fmt.Println("Error parsing node ID: ", err)
	}
	return id
}
