package sequebuf

import (
	"log"
	"reflect"
)

func (bs *binaryStream) unmarshalConverter(field *reflect.Value) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(bs.readString())
	case reflect.Bool:
		field.SetBool(bs.readBool())
	case reflect.Uint8:
		field.SetUint(uint64(bs.readUInt8()))
	case reflect.Uint16:
		field.SetUint(uint64(bs.readUInt16()))
	case reflect.Uint32:
		field.SetUint(uint64(bs.readUInt32()))
	case reflect.Uint64:
		field.SetUint(bs.readUInt64())
	case reflect.Int8:
		field.SetInt(int64(bs.readInt8()))
	case reflect.Int16:
		field.SetInt(int64(bs.readInt16()))
	case reflect.Int32:
		field.SetInt(int64(bs.readInt32()))
	case reflect.Int64:
		field.SetInt(bs.readInt64())
	case reflect.Float32:
		field.SetFloat(float64(bs.readFloat32()))
	case reflect.Float64:
		field.SetFloat(bs.readFloat64())
	case reflect.Slice:
		// 读取数量
		num := bs.readUInt64()
		newV := reflect.MakeSlice(field.Type(), 0, int(num))

		for i := 0; i < int(num); i++ {

			newField := reflect.New(field.Type().Elem()).Elem()
			err := bs.unmarshalConverter(&newField)
			if err != nil {
				return err
			}
			newV = reflect.Append(newV, newField)
		}
		field.Set(newV)
	case reflect.Struct:
		for k := 0; k < field.NumField(); k++ {
			field2 := field.Field(k)
			err := bs.unmarshalConverter(&field2)
			if err != nil {
				return err
			}
		}
	case reflect.Map:
		mapSize := bs.readInt64()
		var i int64 = 0
		keyType := field.Type().Key()
		valType := field.Type().Elem()
		e := reflect.MakeMap(field.Type())
		for i = 0; i < mapSize; i++ {
			k := reflect.New(keyType).Elem()
			err := bs.unmarshalConverter(&k)
			if err != nil {
				return err
			}
			v := reflect.New(valType).Elem()
			err = bs.unmarshalConverter(&v)
			if err != nil {
				return err
			}
			e.SetMapIndex(k, v)
		}
		field.Set(e)
	default:
		log.Fatal("Unsupported Field type : ", field.Kind())
	}
	return nil

}
