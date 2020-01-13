```
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
```

```
output:

2020/01/13 19:17:56 processHandle 224
2020/01/13 19:17:56 baseAddress 140699741650944
2020/01/13 19:17:56 READ [2 0 102 136]
2020/01/13 19:17:56 WRITE 4 true
2020/01/13 19:17:56 READ [0 1 2 3]
```