package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"syscall"
)

func main() {
	// Memory Mapping using syscalls

	// Handle any panics/page fault error
	// As seen here => https://github.com/buildbarn/bb-storage/blob/c346ca331930f1bc5e4f9bde75de96ee3e6c8a9c/pkg/blockdevice/memory_mapped_block_device_unix.go#L49
	old := debug.SetPanicOnFault(true)
	defer func() {
		debug.SetPanicOnFault(old)
		if recover() != nil {
			fmt.Println("page fault error")
		}
	}()

	// Open the File
	f, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}

	// Retrieve the integer unix file descriptor
	fd := int(f.Fd())

	// Retrieve the file size
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	size := int(stat.Size())

	// Make a Syscall & map to memory
	b, err := syscall.Mmap(fd, 0, size, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b[50:]))

	// Delete the mappings for the specified address range,
	// Any references to this will generate invalid memory address
	err = syscall.Munmap(b)
	if err != nil {
		panic(err)
	}

	// This will panic because we've trashed the mapping in the baove step with Munmap
	//fmt.Println(string(b[50:]))
}
