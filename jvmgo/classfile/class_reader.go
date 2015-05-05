package classfile

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil/bigendian"
)

type ClassReader struct {
	data []byte
}

func newClassReader(data []byte) *ClassReader {
	return &ClassReader{data}
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := bigendian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readUint32() uint32 {
	val := bigendian.Int32(self.data)
	self.data = self.data[4:]
	return uint32(val)
}
func (self *ClassReader) readInt32() int32 {
	val := bigendian.Int32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readInt64() int64 {
	val := bigendian.Int64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readFloat32() float32 {
	val := bigendian.Float32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readFloat64() float64 {
	val := bigendian.Float64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}

// todo
func (self *ClassReader) readString() string {
	length := uint32(self.readUint16())
	bytes := self.readBytes(length)
	return string(bytes)
}

// todo
// see java.io.DataInputStream.readUTF(DataInput)
func (self *ClassReader) readMUTF8() string {

	utflen := uint32(self.readUint16())
	bytearr := self.readBytes(utflen)
	chararr := make([]uint16, utflen)

	var c, char2, char3 int32
	var count uint32 = 0
	var chararr_count int32 = 0

	for count < utflen {
		c = int32(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = uint16(c)
		chararr_count++
	}
	println(char2, char3)
	// while (count < utflen) {
	//     c = (int) bytearr[count] & 0xff;
	//     switch (c >> 4) {
	//         case 0: case 1: case 2: case 3: case 4: case 5: case 6: case 7:
	//             /* 0xxxxxxx*/
	//             count++;
	//             chararr[chararr_count++]=(char)c;
	//             break;
	//         case 12: case 13:
	//             /* 110x xxxx   10xx xxxx*/
	//             count += 2;
	//             if (count > utflen)
	//                 throw new UTFDataFormatException(
	//                     "malformed input: partial character at end");
	//             char2 = (int) bytearr[count-1];
	//             if ((char2 & 0xC0) != 0x80)
	//                 throw new UTFDataFormatException(
	//                     "malformed input around byte " + count);
	//             chararr[chararr_count++]=(char)(((c & 0x1F) << 6) |
	//                                             (char2 & 0x3F));
	//             break;
	//         case 14:
	//             /* 1110 xxxx  10xx xxxx  10xx xxxx */
	//             count += 3;
	//             if (count > utflen)
	//                 throw new UTFDataFormatException(
	//                     "malformed input: partial character at end");
	//             char2 = (int) bytearr[count-2];
	//             char3 = (int) bytearr[count-1];
	//             if (((char2 & 0xC0) != 0x80) || ((char3 & 0xC0) != 0x80))
	//                 throw new UTFDataFormatException(
	//                     "malformed input around byte " + (count-1));
	//             chararr[chararr_count++]=(char)(((c     & 0x0F) << 12) |
	//                                             ((char2 & 0x3F) << 6)  |
	//                                             ((char3 & 0x3F) << 0));
	//             break;
	//         default:
	//             /* 10xx xxxx,  1111 xxxx */
	//             throw new UTFDataFormatException(
	//                 "malformed input around byte " + count);
	//     }
	// }
	// // The number of chars produced may be less than utflen
	// return new String(chararr, 0, chararr_count);
	return ""
}
