package main

import (
	"log"
	"unsafe"
)

func main() {
	pid := 10360

	processHandle := OpenProcess(pid)
	log.Println("processHandle", processHandle)

	baseAddress := getBaseAddress(processHandle)
	log.Println("baseAddress", baseAddress)

	result := READ(processHandle, baseAddress + 0x2FF0, 4)
	log.Println("READ", result)

	ptr := []byte{0, 1, 2, 3}
	writted, ok := WRITE(processHandle, baseAddress + 0x2FF0, uintptr(unsafe.Pointer(&ptr[0])), 4)
	log.Println("WRITE", writted, ok)

	result = READ(processHandle, baseAddress + 0x2FF0, 4)
	log.Println("READ", result)
}
