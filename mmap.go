package main

import (
	"fmt"
	"golang.org/x/exp/mmap"
)

func mmap2() error {
	// Provide the filename & let mmap do the magic
	r, err := mmap.Open("text.txt")
	if err != nil {
		return err
	}

	//Read the entire file in memory at an offset of 0 => te entire thing
	p := make([]byte, r.Len())
	_, err = r.ReadAt(p, 0)
	if err != nil {
		return err
	}

	// print the file
	fmt.Println(string(p))
	return nil
}
