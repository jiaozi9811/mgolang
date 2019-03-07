bytes 
该包定义了一些操作byte slice的便利操作。因为字符串可以表示为 []byte，因此，bytes 包定义的函数、方法等和 strings 包很类似

func Contains(b, subslice []byte) bool// 子slice subslice 在 b 中，返回 true

func Count(s, sep []byte) int// slice sep 在 s 中出现的次数（无重叠）
