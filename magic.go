package monkey

import (
	"reflect"
	"syscall"
	"unsafe"
)

func rawMemoryAccess(p uintptr, length int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
}

func pagestart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}

// this function is super unsafe
// aww yeah
func copyToLocation(location uintptr, data []byte) {
	f := rawMemoryAccess(location, len(data))

	page := rawMemoryAccess(pagestart(location), syscall.Getpagesize())
	err := syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	if err != nil {
		panic(err)
	}
	copy(f, data[:])

	err = syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_EXEC)
	if err != nil {
		panic(err)
	}
}

func magic(from, to uintptr) (original []byte) {
	jumpData := jmpToFunctionValue(to)
	f := rawMemoryAccess(from, len(jumpData))
	original = make([]byte, len(f))
	copy(original, f)

	copyToLocation(from, jumpData[:])
	return
}
