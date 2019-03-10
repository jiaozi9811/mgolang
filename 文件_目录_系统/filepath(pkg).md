# path/filepath


func Rel(basepath, targpath string) (string, error) //返回一个相对路径，将basepath和该路径用路径分隔符连起来的新路径在词法上等价于targpath
func SplitList(path string) []string        /将多个路径分割开
func Split(path string) (dir, file string)  //将路径从最后一个路径分隔符后面位置分隔为两个部分（dir和file）并返回
func Join(elem ...string) string        //将任意数量的路径元素放入一个单一路径里，会根据需要添加路径分隔符
func FromSlash(path string) string      //将path中的斜杠（'/'）替换为路径分隔符
func ToSlash(path string) string        //将path中的路径分隔符替换为斜杠
func VolumeName(path string) (v string) //返回最前面的卷名
func Dir(path string) string        //返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录
func Base(path string) string       //返回路径的最后一个元素
func Ext(path string) string        //返回path文件扩展名
func Clean(path string) string
func EvalSymlinks(path string) (string, error)  //返回path指向的符号链接（软链接）所包含的路径
func Match(pattern, name string) (matched bool, err error)
func Glob(pattern string) (matches []string, err error)
type WalkFunc
func Walk(root string, walkFn WalkFunc) error
func HasPrefix(p, prefix string) bool


## func Abs IsAbs

func IsAbs(path string) (b bool)
func Abs(path string) (string, error)

returns an absolute representation of path
//返回path代表的绝对路径，如果path不是绝对路径，会加入当前工作目录以使之成为绝对路径

## Base(name) Dir(name) Ext(扩展名)

func Base(path string) string
func Dir(path string) string
func Ext(path string) string

```go
fmt.Println(path.Base("/usr/local/go")) // go
fmt.Println(path.Dir("/usr/local/go")) // /usr/local
fmt.Printf("No dots: %q\n", filepath.Ext("index")) // ""
fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js")) // ".js"
```

## func Rel(basepath, targpath string) (string, error)

返回一个相对路径(relative path)，当使用分隔符连接到basepath时，该路径在词法上等效于targpath

## func Clean(path string) string

1. 将连续的多个路径分隔符替换为单个路径分隔符
2. 剔除每一个.路径名元素（代表当前目录）
3. 剔除每一个路径内的..路径名元素（代表父目录）和它前面的非..路径名元素
4. 剔除开始一个根路径的..路径名元素，即将路径开始处的"/.."替换为"/"（假设路径分隔符是'/'）

## func EvalSymlinks(path string) (string, error) 软链接

返回path指向的符号链接（软链接）所包含的路径

## FromSlash ToSlash

func FromSlash(path string) string
将path中的斜杠（'/'）替换为路径分隔符并返回替换结果
func ToSlash(path string) string
将path中的路径分隔符替换为斜杠（'/'）并返回替换结果

## func Glob(pattern string) (matches []string, err error)

返回所有匹配模式匹配字符串pattern的文件或者nil（如果没有匹配的文件）。pattern的语法和Match函数相同

```go
    list, err := filepath.Glob("/usr/*/*/[Bb][Aa]*")
    if err != nil {
        fmt.Println(err)
    }
    for _, v := range list {
        fmt.Println(v)
    }
```

## func Join(elem ...string) string

将任意数量的路径元素放入一个单一路径里，会根据需要添加路径分隔符

## func Split(path string) (dir, file string)

将路径从最后一个路径分隔符后面位置分隔为两个部分（dir和file）并返回
func SplitList(path string) []string
将PATH或GOPATH等环境变量里的多个路径分割开(这些路径被OS特定的表分隔符连接起来)
`fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin")) // On Unix: [/a/b/c /usr/bin]`

## func Match(pattern, name string) (matched bool, err error)

```go
pattern:
	{ term }
term:
	'*'         matches any sequence of non-Separator characters
	'?'         matches any single non-Separator character
	'[' [ '^' ] { character-range } ']'
	            character class (must be non-empty)
	c           matches character c (c != '*', '?', '\\', '[')
	'\\' c      matches character c

character-range:
	c           matches character c (c != '\\', '-', ']')
	'\\' c      matches character c
	lo '-' hi   matches character c for lo <= c <= hi
```

## func VolumeName(path string) string

返回最前面的卷名。如Windows系统里提供参数"C:\foo\bar"会返回"C:"；Unix/linux系统的"\\host\share\foo"会返回"\\host\share"

## func Walk(root string, walkFn WalkFunc) error

遍历root指定的目录下的文件树，对每一个该文件树中的目录和文件都会调用walkFn，包括root自身

## type WalkFunc func(path string, info os.FileInfo, err error) error

对每一个文件/目录都会调用WalkFunc函数类型值