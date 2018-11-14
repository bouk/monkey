//+build !windows

package monkey

import (
	"syscall"
)

func mprotectCrossPage(addr uintptr, len int, prot int) {
	pageSize := syscall.Getpagesize()
	for p := pageStart(addr); p <= addr + uintptr(len); p += uintptr(pageSize) {
		page := rawMemoryAccess(p, syscall.Getpagesize())
		err := syscall.Mprotect(page, prot)
		if err != nil {
			panic(err)
		}
	}
}

// this function is super unsafe
// aww yeah
// It copies a slice to a raw memory location, disabling all memory protection before doing so.
func copyToLocation(location uintptr, data []byte) {
	f := rawMemoryAccess(location, len(data))

	mprotectCrossPage(location, len(data), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	copy(f, data[:])
	mprotectCrossPage(location, len(data), syscall.PROT_READ|syscall.PROT_EXEC)
}
