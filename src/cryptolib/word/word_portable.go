//go:build !amd64

package word

import "encoding/binary"

func U16toByte_16(a uint16) (b []byte) {
	b = make([]byte, 2)
	b[0] = byte(a >> 8)
	b[1] = byte(a)
	return
}

func U32toByte_32(a uint32) (b []byte) {
	b = make([]byte, 4)
	b[0] = byte(a >> 24)
	b[1] = byte(a >> 16)
	b[2] = byte(a >> 8)
	b[3] = byte(a)
	return
}

func BytetoU32_32(a []byte) (b uint32) {
	b = uint32(a[0])<<24 | uint32(a[1])<<16 | uint32(a[2])<<8 | uint32(a[3])
	return
}

func U64toByte_64(a uint64) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, a)
	return
}

func BytetoU64_64(a []byte) (b uint64) {
	return binary.BigEndian.Uint64(a[:8])
}

func U64toByte_256(a [4]uint64) (b []byte) {
	b = make([]byte, 32)
	binary.BigEndian.PutUint64(b[0:8], a[3])
	binary.BigEndian.PutUint64(b[8:16], a[2])
	binary.BigEndian.PutUint64(b[16:24], a[1])
	binary.BigEndian.PutUint64(b[24:32], a[0])
	return
}

func BytetoU64_256(x []byte) (y [4]uint64) {
	y[3] = binary.BigEndian.Uint64(x[0:8])
	y[2] = binary.BigEndian.Uint64(x[8:16])
	y[1] = binary.BigEndian.Uint64(x[16:24])
	y[0] = binary.BigEndian.Uint64(x[24:32])
	return
}

// b is explored with big-endian, a is little-endian for u32
func U32toByte_256(a [8]uint32) (b []byte) {
	b = make([]byte, 32)
	for i := 0; i < 8; i++ {
		binary.BigEndian.PutUint32(b[i*4:(i+1)*4], a[i])
	}
	return
}

func BytetoU32_256(a []byte) (b [8]uint32) {
	for i := 0; i < 8; i++ {
		b[i] = binary.BigEndian.Uint32(a[i*4 : (i+1)*4])
	}
	return
}

func U64toU32_256(a [4]uint64) (b [8]uint32) {
	return BytetoU32_256(U64toByte_256(a))
}

func U32toU64_256(a [8]uint32) (b [4]uint64) {
	return BytetoU64_256(U32toByte_256(a))
}

func U32toByte_128(a [4]uint32) (b []byte) {
	b = make([]byte, 16)
	for i := 0; i < 4; i++ {
		binary.BigEndian.PutUint32(b[i*4:(i+1)*4], a[i])
	}
	return
}

func BytetoU32_128(a []byte) (b [4]uint32) {
	for i := 0; i < 4; i++ {
		b[i] = binary.BigEndian.Uint32(a[i*4 : (i+1)*4])
	}
	return
}

func BytetoU32_512(x []byte) (y [16]uint32) {
	for i := 0; i < 16; i++ {
		y[i] = binary.BigEndian.Uint32(x[i*4 : (i+1)*4])
	}
	return
}
