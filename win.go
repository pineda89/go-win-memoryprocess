package main

import (
	"syscall"
	"unsafe"
)

var (
	kernel32 = syscall.MustLoadDLL("kernel32.dll")
	psapi                        = syscall.NewLazyDLL("Psapi.dll")

	procOpenProcess = kernel32.MustFindProc("OpenProcess")
	procReadProcessMemory = kernel32.MustFindProc("ReadProcessMemory")
	procWriteProcessMemory         = kernel32.MustFindProc("WriteProcessMemory")
	procEnumProcessModules       = psapi.NewProc("EnumProcessModulesEx")
)

const PROCESS_ALL_ACCESS = 0x1F0FFF

func OpenProcess(pid int) uintptr {
	handle, _, _ := procOpenProcess.Call(uintptr(PROCESS_ALL_ACCESS), uintptr(1), uintptr(pid))
	return handle
}

func READ(hProcess uintptr, address, size uintptr) []byte {
	var data = make([]byte, size)
	var length uint32

	procReadProcessMemory.Call(hProcess, address,
		uintptr(unsafe.Pointer(&data[0])),
		size, uintptr(unsafe.Pointer(&length)))

	return data
}

func WRITE(hProcess uintptr, lpBaseAddress, lpBuffer, nSize uintptr) (int, bool) {
	var nBytesWritten int
	ret, _, _ := procWriteProcessMemory.Call(
		uintptr(hProcess),
		lpBaseAddress,
		lpBuffer,
		nSize,
		uintptr(unsafe.Pointer(&nBytesWritten)),
	)

	return nBytesWritten, ret != 0
}


func getBaseAddress(handle uintptr) uintptr {
	modules := [1024]uint64{}
	var needed uintptr
	procEnumProcessModules.Call(
		handle,
		uintptr(unsafe.Pointer(&modules)),
		uintptr(1024),
		uintptr(unsafe.Pointer(&needed)),
		uintptr(0x03),
	)
	for i := uintptr(0); i < needed/unsafe.Sizeof(modules[0]); i++ {
		if i == 0 {
			return uintptr(modules[i])
		}
	}
	return 0
}
