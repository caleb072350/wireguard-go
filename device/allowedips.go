package device

import (
	"math/bits"
	"net"
	"unsafe"
)

type trieEntry struct {
	cidr  uint
	child [2]*trieEntry
	bits  net.IP
	peer  *Peer

	// index of "branching" bit

	bit_at_byte  uint
	bit_at_shift uint
}

func isLittleEndian() bool {
	one := uint32(1)
	return *(*byte)(unsafe.Pointer(&one)) != 0
}

func swapU32(i uint32) uint32 {
	if !isLittleEndian() {
		return i
	}
	return bits.ReverseBytes32(i)
}

func swapU64(i uint64) uint64 {
	if !isLittleEndian() {
		return i
	}
	return bits.ReverseBytes64(i)
}

func commonBits(ip1 net.IP, ip2 net.IP) uint {
	size := len(ip1)
	if size == net.IPv4len {
		a := (*uint32)(unsafe.Pointer(&ip1[0]))
		b := (*uint32)(unsafe.Pointer(&ip2[0]))
		x := *a ^ *b
		return uint(bits.LeadingZeros32(swapU32(x)))
	} else if size == net.IPv6len {
		a := (*uint64)(unsafe.Pointer(&ip1[0]))
		b := (*uint64)(unsafe.Pointer(&ip2[0]))
		x := *a ^ *b
		if x != 0 {
			return uint(bits.LeadingZeros64(swapU64(x)))
		}
		a = (*uint64)(unsafe.Pointer(&ip1[8]))
		b = (*uint64)(unsafe.Pointer(&ip2[8]))
		x = *a ^ *b
		return 64 + uint(bits.LeadingZeros64(swapU64(x)))
	} else {
		panic("Wrong size bit string")
	}
}
