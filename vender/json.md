序列化
func Marshall(v interface{}) ([]byte,error)

反序列化
fun Unmarshal(data []byte,v interface{}) error

标签
type mystruct struct {
    SomeField string `json:"some_field"`
}

## 指定字段为empty时的行为
使用omitempty,则field的值为zero,序列化后的json将不包含此字段

## 跳过字段
"-"表示跳过指定的field，即序列化的时候不输出。可以保护需保护的字段不被序列化
