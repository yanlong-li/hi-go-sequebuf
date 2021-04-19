package sequebuf

import "reflect"

//type Interface interface {
//	Marshal(uint8, interface{})
//	Unmarshal(interface{}) []reflect.Value
//	GetData() []byte
//}

type binaryStream struct {
	data  []byte //数据储存体
	index uint16 //当前指针
}

func (bs *binaryStream) GetData() []byte {
	return bs.data
}

func (bs *binaryStream) SetData(Data []byte) {
	bs.data = Data
}

// Marshal 将包结构体反射写入字节流中
func Marshal(model interface{}) ([]byte, error) {
	bs := &binaryStream{}
	// 清空
	elem := reflect.ValueOf(model)
	return bs.data, bs.marshalConverter(elem)
}

// Unmarshal 从字节流中反射出对应的结构体并注入到指定方法中
func Unmarshal(data []byte, model interface{}) error {
	bs := &binaryStream{
		data: data,
	}
	// 从类型中新建变量
	elem := reflect.ValueOf(model).Elem()
	//构造一个存放函数实参 Value 值的数纽
	return bs.unmarshalConverter(&elem)
}
