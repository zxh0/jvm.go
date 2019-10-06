package jimage

import (
	"github.com/zxh0/jvm.go/vmutils"
)

const (
	HashMultiplier = 0x01000193
	PositiveMask   = 0x7FFFFFFF
)

func getStringBytes(strings []byte, offset int) []byte {
	nBytes := len(strings)
	for i := offset; i < nBytes; i++ {
		ch := strings[i]
		if ch == 0 {
			return strings[offset:i]
		}
	}
	panic("No terminating zero byte for modified UTF-8 byte sequence")
}

func hashCode(s string, seed int32) int32 {
	return unmaskedHashCode(s, seed) & PositiveMask
}

// public static int unmaskedHashCode(String s, int seed)
// https://github.com/unofficial-openjdk/openjdk/blob/jdk/jdk/src/java.base/share/classes/jdk/internal/jimage/ImageStringsReader.java
func unmaskedHashCode(s string, seed int32) int32 {
	chars := vmutils.UTF8ToUTF16(s)
	slen := len(chars)
	var buffer []byte

	for i := 0; i < slen; i++ {
		ch := chars[i]
		uch := ch // & 0xFFFF

		if uch & ^uint16(0x7F) != 0 {
			if buffer == nil {
				buffer = make([]byte, 8)
			}
			mask := ^uint16(0x3F)
			n := 0

			for {
				buffer[n] = (byte)(0x80 | (uch & 0x3F))
				n++
				uch >>= 6
				mask >>= 1

				if uch&mask == 0 {
					break
				}
			}

			buffer[n] = (byte)((mask << 1) | uch)

			for {
				seed = (seed * HashMultiplier) ^ int32(buffer[n]&0xFF)
				n--
				if n < 0 {
					break
				}
			}
		} else if uch == 0 {
			seed = (seed * HashMultiplier) ^ (0xC0)
			seed = (seed * HashMultiplier) ^ (0x80)
		} else {
			seed = (seed * HashMultiplier) ^ int32(uch)
		}
	}
	return seed
}
