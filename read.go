package sequebuf

import (
	"encoding/binary"
	"math"
)

// readBool 读取bool值
func (bs *binaryStream) readBool() bool {
	if bs.readUInt8() > 0 {
		return true
	}
	return false
}

// readUInt8 读取 Uint8 as byte
func (bs *binaryStream) readUInt8() (data uint8) {
	if bs.checkLen(1) {
		data = bs.data[bs.index]
		bs.index++
		return
	}
	return
}

// readUInt16 读取 Uint16
func (bs *binaryStream) readUInt16() (data uint16) {
	if bs.checkLen(2) {
		data = binary.LittleEndian.Uint16(bs.data[bs.index : bs.index+2])
		bs.index += 2
		return
	}
	return
}

// readUInt32 读取 Uint32
func (bs *binaryStream) readUInt32() (data uint32) {
	if bs.checkLen(4) {
		data = binary.LittleEndian.Uint32(bs.data[bs.index : bs.index+4])
		bs.index += 4
		return
	}
	return
}

// readUInt64 读取 Uint64
func (bs *binaryStream) readUInt64() (data uint64) {
	if bs.checkLen(8) {
		data = binary.LittleEndian.Uint64(bs.data[bs.index : bs.index+8])
		bs.index += 8
		return
	}
	return
}

// readInt8 读取 int8 as byte
func (bs *binaryStream) readInt8() (data int8) {
	if bs.checkLen(1) {
		data = int8(bs.data[bs.index])
		bs.index++
		return
	}
	return
}

// readInt16 读取 int16
func (bs *binaryStream) readInt16() (data int16) {
	if bs.checkLen(2) {

		data = int16(binary.LittleEndian.Uint16(bs.data[bs.index : bs.index+2]))
		bs.index += 2
		return
	}
	return
}

// readInt32 读取 int32
func (bs *binaryStream) readInt32() (data int32) {
	if bs.checkLen(4) {
		data = int32(binary.LittleEndian.Uint32(bs.data[bs.index : bs.index+4]))
		bs.index += 4
		return
	}
	return
}

// readInt64 读取 int64
func (bs *binaryStream) readInt64() (data int64) {
	if bs.checkLen(8) {
		data = int64(binary.LittleEndian.Uint64(bs.data[bs.index : bs.index+8]))
		bs.index += 8
		return
	}
	return
}


// readFloat32 读取 Float32
func (bs *binaryStream) readFloat32() (data float32) {
	if bs.checkLen(4) {
		data = math.Float32frombits(binary.LittleEndian.Uint32(bs.data[bs.index : bs.index+4]))
		bs.index += 4
	}
	return
}

// readFloat64 读取 Float64
func (bs *binaryStream) readFloat64() (data float64) {
	if bs.checkLen(8) {
		data = math.Float64frombits(binary.LittleEndian.Uint64(bs.data[bs.index : bs.index+8]))
		bs.index += 8
	}
	return
}

// readString 读取可变长度字符串
func (bs *binaryStream) readString() (data string) {
	return bs.readStringL(bs.readUInt16())
}

// readStringL 读取固定长度字符串
func (bs *binaryStream) readStringL(length uint16) (data string) {
	if bs.checkLen(length) {
		data = string(bs.data[bs.index : bs.index+length])
		bs.index += length
	}
	return
}

// 检查数据长度是否足够
func (bs *binaryStream) checkLen(length uint16) bool {
	if uint16(len(bs.data)) >= bs.index+length {
		return true
	}
	return false
}

// Bool
//	Int
//	Int8
//	Int16
//	Int32
//	Int64
//	Uint
//	Uint8
//	Uint16
//	Uint32
//	Uint64
//	Uintptr
//	Float32
//	Float64
//	Complex64
//	Complex128
//	Array
//	Chan
//	Func
//	Interface
//	Map
//	Ptr
//	Slice
//	String
//	Struct
