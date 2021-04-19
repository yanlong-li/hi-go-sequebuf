package sequebuf

import (
	"encoding/binary"
	"math"
)

func (bs *binaryStream) writeBool(data bool) {
	if data {
		bs.data = append(bs.data, byte(1))
	} else {
		bs.data = append(bs.data, byte(0))
	}
}

func (bs *binaryStream) writeUint8(n uint8) {
	bs.data = append(bs.data, n)
}

func (bs *binaryStream) writeUint16(n uint16) {
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, n)
	bs.data = append(bs.data, data...)
}

func (bs *binaryStream) writeUint32(n uint32) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, n)
	bs.data = append(bs.data, data...)
}

func (bs *binaryStream) writeUint64(n uint64) {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, n)
	bs.data = append(bs.data, data...)
}

func (bs *binaryStream) writeInt8(n int8) {
	bs.writeUint8(uint8(n))
}

func (bs *binaryStream) writeInt16(n int16) {
	bs.writeUint16(uint16(n))
}

func (bs *binaryStream) writeInt32(n int32) {
	bs.writeUint32(uint32(n))
}

func (bs *binaryStream) writeInt64(n int64) {
	bs.writeUint64(uint64(n))
}

func (bs *binaryStream) writeFloat32(n float32) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, math.Float32bits(n))
	bs.data = append(bs.data, data...)
}

func (bs *binaryStream) writeFloat64(n float64) {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, math.Float64bits(n))
	bs.data = append(bs.data, data...)
}

func (bs *binaryStream) writeString(n string) {
	bs.writeUint16(uint16(len(n)))
	bs.data = append(bs.data, n...)
}
