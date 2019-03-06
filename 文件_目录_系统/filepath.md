path/filepath

func IsAbs(path string) bool            //返回路径是否是一个绝对路径
func Abs(path string) (string, error)       //返回path代表的绝对路径，如果path不是绝对路径，会加入当前工作目录以使之成为绝对路径
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












