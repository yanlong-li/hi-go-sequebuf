package sequebuf

import (
	"errors"
	"reflect"
)

func (bs *binaryStream) marshalConverter(field reflect.Value) error {
	switch field.Kind() {
	case reflect.String:
		bs.writeString(field.String())
	case reflect.Int:
		bs.writeInt64(field.Int())
	case reflect.Bool:
		bs.writeBool(field.Bool())
	case reflect.Uint8:
		bs.writeUint8(uint8(field.Uint()))
	case reflect.Uint16:
		bs.writeUint16(uint16(field.Uint()))
	case reflect.Uint32:
		bs.writeUint32(uint32(field.Uint()))
	case reflect.Uint64:
		bs.writeUint64(field.Uint())
	case reflect.Int8:
		bs.writeInt8(int8(field.Int()))
	case reflect.Int16:
		bs.writeInt16(int16(field.Int()))
	case reflect.Int32:
		bs.writeInt32(int32(field.Int()))
	case reflect.Int64:
		bs.writeInt64(field.Int())
	case reflect.Float32:
		bs.writeFloat32(float32(field.Float()))
	case reflect.Float64:
		bs.writeFloat64(field.Float())
	case reflect.Slice:
		bs.writeInt64(int64(field.Len()))
		for i := 0; i < field.Len(); i++ {
			elm := field.Index(i)
			err := bs.marshalConverter(elm)
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		for k := 0; k < field.NumField(); k++ {
			field := field.Field(k)
			err := bs.marshalConverter(field)
			if err != nil {
				return err
			}
		}
	case reflect.Map:
		keys := field.MapKeys()
		bs.writeInt64(int64(field.Len()))
		for _, k := range keys {
			value := field.MapIndex(k)
			err := bs.marshalConverter(k)
			if err != nil {
				return err
			}
			err = bs.marshalConverter(value)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("Unsupported Field type : " + field.Kind().String())
	}

	return nil
}
