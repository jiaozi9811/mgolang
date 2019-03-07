unicode
go对unicode的支持包含三个包:

unicode
unicode/utf8
unicode/utf16
unicode包包含基本的字符判断函数。utf8包主要负责rune和byte之间的转换。utf16包负责rune和uint16数组之间的转换

go语言的所有代码都是UTF8的，所以如果我们在程序中的字符串都是utf8编码的，但是我们的单个字符（单引号扩起来的）却是unicode的

func IsControl(r rune) bool  // 是否控制字符
func IsDigit(r rune) bool  // 是否阿拉伯数字字符，即1-9
func IsGraphic(r rune) bool // 是否图形字符
func IsLetter(r rune) bool // 是否字母
func IsLower(r rune) bool // 是否小写字符
func IsMark(r rune) bool // 是否符号字符
func IsNumber(r rune) bool // 是否数字字符，比如罗马数字Ⅷ也是数字字符
func IsOneOf(ranges []*RangeTable, r rune) bool // 是否是RangeTable中的一个
func IsPrint(r rune) bool // 是否可打印字符
func IsPunct(r rune) bool // 是否标点符号
func IsSpace(r rune) bool // 是否空格
func IsSymbol(r rune) bool // 是否符号字符
func IsTitle(r rune) bool // 是否title case
func IsUpper(r rune) bool // 是否大写字符
-------------------------------------------------------------------------------------
unicode/utf8
字节和字符的转换
判断是否符合utf8编码的函数：

func Valid(p []byte) bool
func ValidRune(r rune) bool
func ValidString(s string) bool
判断rune的长度的函数：

func RuneLen(r rune) int
判断字节串或者字符串的rune数

func RuneCount(p []byte) int
func RuneCountInString(s string) (n int)
编码和解码rune到byte

func DecodeRune(p []byte) (r rune, size int)
func EncodeRune(p []byte, r rune) int
-------------------------------------------------------------------------------------
unicode/utf16
将int16和rune进行转换

func Decode(s []uint16) []rune
func Encode(s []rune) []uint16