Sequence Buffer 序列缓冲器
===

序列缓冲器，原始的纯数据报文。 Sequence buffer, the original pure data message.

编码后的数据为 []byte, 可以方便进行二次压缩及加密等操作。 The encoded data is [] byte, which is convenient for secondary compression and
encryption.

## 优点 Advantages

* 编码和解码速度快
* 相比带有 key 的报文编码器，可以小幅度缩减报文长度。
* The speed of encoding and decoding is fast
* Compared with the message encoder with key, the length of message can be reduced slightly.

## 缺点 Disadvantages

* 结构体依赖键的顺序
* 结构体无法随意变更键类型，
* 结构体可以在尾部增加键不可以删减，在低版本中不存在的键会自动忽略后续数据。

* The structure depends on the order of the key
* Structs can't change keys at will. They can be added and can't be deleted. Keys that do not exist in lower versions
  automatically ignore subsequent data.
