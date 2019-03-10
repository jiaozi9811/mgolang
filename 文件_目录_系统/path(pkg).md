# path

[TOC]

## func Match(pattern, name string) (matched bool, err error)

如果name匹配shell文件名模式匹配字符串，Match函数返回真。该模式匹配字符串语法为：

```m
pattern:
	{ term }
term:
	'*'                                  匹配0或多个非/的字符
	'?'                                  匹配1个非/的字符
	'[' [ '^' ] { character-range } ']'  字符组（必须非空）
	c                                    匹配字符c（c != '*', '?', '\\', '['）
	'\\' c                               匹配字符c
character-range:
	c           匹配字符c（c != '\\', '-', ']'）
	'\\' c      匹配字符c
	lo '-' hi   匹配区间[lo, hi]内的字符
```

## func Base(path string) string

`fmt.Println(path.Base("/a/b"))   // b`

## func Clean(path string) string

```go
paths := []string{
    "a/c",
    "a//c",
    "a/c/.",
    "a/c/b/..",
    "/../a/c",
    "/../a/b/../././/c",
}
for _, p := range paths {
    fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
}
/*
Clean("a/c") = "a/c"
Clean("a//c") = "a/c"
Clean("a/c/.") = "a/c"
Clean("a/c/b/..") = "a/c"
Clean("/../a/c") = "/a/c"
Clean("/../a/b/../././/c") = "/a/c"
*/
```

## func Dir(path string) string

```go
fmt.Println(path.Dir("/a/b/c"))
fmt.Println(path.Dir("a/b/c"))
fmt.Println(path.Dir("/a/"))
fmt.Println(path.Dir("a/"))
fmt.Println(path.Dir("/"))
fmt.Println(path.Dir(""))
/*
/a/b
a/b
/a
a
/
.
*/
```

## func Ext(path string) string

返回path文件扩展名
`fmt.Println(path.Ext("/a/b/c/bar.css"))  // .css`

## func IsAbs(path string) bool 绝对路径

IsAbs返回路径是否是一个绝对路径
`fmt.Println(path.IsAbs("/dev/null")) // true`

## func Join(elem ...string) string

Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠
`fmt.Println(path.Join("a", "b", "c")) //a/b/c`

## func Split(path string) (dir, file string)

Split函数将路径从最后一个斜杠后面位置分隔为两个部分（dir和file）并返回
`fmt.Println(path.Split("static/myfile.css")) // static/ myfile.css`