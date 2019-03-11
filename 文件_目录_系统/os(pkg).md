# os

[TOC]

## 读写模式

O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件

## Stdin Stdout Stderr

Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")

var Args []string     //Args保管了命令行参数，第一个是程序名

## func Hostname() (name string, err error)

获取主机名

## func Create(name string) (*File, error) and func NewFile(fd uintptr, name string) *File

创建信息

## func Mkdir(name string, perm FileMode) error and func MkdirAll(path string, perm FileMode) error

创建目录

## 打开文件func OpenFile(name string, flag int, perm FileMode) (*File, error) and func Open(name string) (*File, error)

打开一个文件，一般通过 Open 或 Create，我们看这两个函数的实现

```go
func Open(name string) (*File, error) {
     return OpenFile(name, O_RDONLY, 0) }
func Create(name string) (*File, error) {
    return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666) }
```

## Remove RemoveAll Rename

func Remove(name string) error
func RemoveAll(path string) error
func Rename(oldpath, newpath string) error

## UserCacheDir UserHomeDir TempDir

func UserCacheDir() (string, error)
func UserHomeDir() (string, error)
func TempDir() string

## 存在性IsExist IsNotExist

func IsExist(err error) bool
func IsNotExist(err error) bool

## func IsPathSeparator(c uint8) bool

返回字符c是否是一个路径分隔符

## func IsPermission(err error) bool

返回一个布尔值说明该错误是否表示因权限不足要求被拒绝

## func IsTimeout(err error) bool

## 环境变量操作 Clearenv Setenv Getenv LookupEnv

func Clearenv()
func Environ() []string
func ExpandEnv(s string) string // 替换环境变量
func LookupEnv(key string) (string, bool) // 查看一个环境变量
func Setenv(key, value string) error
func Unsetenv(key string) error

func Getenv(key string) string

## 获取执行文件信息func Executable() (string, error)

## func Getwd() (dir string, err error) //返回一个对应当前工作目录的根路径

## func Chdir(dir string) error // 切换目录

## 替换func Expand(s string, mapping func(string) string) string

替换string中的变量

```go
mapper := func(placeholderName string) string {
    switch placeholderName {
    case "DAY_PART":
        return "morning"
    case "NAME":
        return "Gopher"
    }

    return ""
}
fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper))
//Good morning, Gopher!
```

## func Exit(code int) // 当前程序以给出的状态码code退出

## Chmod Chtimes

func Chmod(name string, mode FileMode) error

func Chtimes(name string, atime time.Time, mtime time.Time) error

## Chown Lchown

func Chown(name string, uid, gid int) error
func Lchown(name string, uid, gid int) error

## Link(链接) Symlink(软链接)

func Link(oldname, newname string) error
func Symlink(oldname, newname string) error

## func Readlink(name string) (string, error) //真实路径

获取name指定的符号链接文件指向的文件的路径

## func SameFile(fi1, fi2 FileInfo) bool 是否同一文件

## func Truncate(name string, size int64) error //修改name指定的文件的大小

## func Pipe() (r *File, w *File, err error)

## func NewSyscallError(syscall string, err error) error

## Getuid Getgid Getpid Getgroups

func Getegid() int
func Geteuid() int
func Getgid() int
func Getgroups() ([]int, error)
func Getpagesize() int
func Getpid() int
func Getppid() int
func Getuid() int


## 读取文件内容：Read

func (f *File) Read(b []byte) (n int, err error)
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
Read 和 ReadAt 的区别：前者从文件当前偏移量处读，且会改变文件当前的偏移量；而后者从 off 指定的位置开始读，且不会改变文件当前偏移量

##  数据写入文件：Write

func (f *File) Write(b []byte) (n int, err error)

## 关闭文件：Close

## 改变文件偏移量：Seek

文件偏移量是指执行下一个Read或Write操作的文件其实位置，会以相对于文件头部起始点的文件当前位置来表示。文件第一个字节的偏移量为0
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
file.Seek(0, os.SEEK_SET)    // 文件开始处
file.Seek(0, SEEK_END)        // 文件结尾处的下一个字节
file.Seek(-1, SEEK_END)        // 文件最后一个字节
file.Seek(-10, SEEK_CUR)     // 当前位置前10个字节
file.Seek(1000, SEEK_END)    // 文件结尾处的下1001个字节

## 文件属性

文件属性具体信息通过 os.FileInfo 接口获取
函数 Stat、Lstat 和 File.Stat 可以得到该接口的实例。这三个函数对应三个系统调用：stat、lstat 和 fstat

stat会返回所命名文件的相关信息
lstat 与 stat 类似，区别在于如果文件是符号链接，那么所返回的信息针对的是符号链接自身（而非符号链接所指向的文件）。
fstat 则会返回由某个打开文件描述符（Go 中则是当前打开文件 File）所指代文件的相关信息。

## 文件属主

func Chown(name string, uid, gid int) error
func Lchown(name string, uid, gid int) error
func (f *File) Chown(uid, gid int) error

## 目录与链接

func Link(oldname, newname string) error
Link 创建一个名为 newname 指向 oldname 的硬链接。如果出错，会返回 *LinkError 类型的错误。

## 符号链接

func Symlink(oldname, newname string) error
Symlink 创建一个名为 newname 指向 oldname 的符号链接。如果出错，会返回 *LinkError 类型的错误

func Readlink(name string) (string, error)
Readlink 获取 name 指定的符号链接指向的文件的路径

### type FileInfo interface

```go
type FileInfo interface {
    Name() string       // 文件的名字（不含扩展名）
    Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
    Mode() FileMode     // 文件的模式位
    ModTime() time.Time // 文件的修改时间
    IsDir() bool        // 等价于Mode().IsDir()
    Sys() interface{}   // 底层数据来源（可以返回nil）
}
```

## Stat Lstat

func Lstat(name string) (FileInfo, error)
func Stat(name string) (FileInfo, error)

## type FileMode

```go
const (
    // 单字符是被String方法用于格式化的属性缩写。
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
    ModeAppend                                     // a: 只能写入，且只能写入到末尾
    ModeExclusive                                  // l: 用于执行
    ModeTemporary                                  // T: 临时文件（非备份文件）
    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
    ModeDevice                                     // D: 设备
    ModeNamedPipe                                  // p: 命名管道（FIFO）
    ModeSocket                                     // S: Unix域socket
    ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
    ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
    ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
    ModeSticky                                     // t: 只有root/创建者能删除/移动文件
    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
)
```

### IsDir IsRegular Perm String

func (m FileMode) IsDir() bool
func (m FileMode) IsRegular() bool // 是否是一个普通文件
func (m FileMode) Perm() FileMode // 返回m的Unix权限位
func (m FileMode) String() string

## type File

func (f *File) Chdir() error
func (f *File) Chmod(mode FileMode) error
func (f *File) Chown(uid, gid int) error
func (file *File) Close() error
func (file *File) Fd() uintptr // 返回与文件f对应的整数类型的Unix文件描述符
func (f *File) Name() string
func (f *File) Read(b []byte) (n int, err error)
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
func (f *File) Readdir(n int) ([]FileInfo, error)
func (f *File) Readdirnames(n int) (names []string, err error)
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
func (f *File) SetDeadline(t time.Time) error
func (f *File) SetReadDeadline(t time.Time) error
func (f *File) SetWriteDeadline(t time.Time) error
func (file *File) Stat() (FileInfo, error)
func (f *File) Sync() error // 递交文件的当前内容进行稳定的存储
func (f *File) SyscallConn() (syscall.RawConn, error)
func (f *File) Truncate(size int64) error
func (f *File) Write(b []byte) (n int, err error)
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
func (f *File) WriteString(s string) (n int, err error)


## type Process

func FindProcess(pid int) (*Process, error)
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
func (p *Process) Kill() error
func (p *Process) Release() error
func (p *Process) Signal(sig Signal) error
func (p *Process) Wait() (*ProcessState, error)

## type ProcessState

func (p *ProcessState) ExitCode() int
func (p *ProcessState) Exited() bool
func (p *ProcessState) Pid() int
func (p *ProcessState) String() string
func (p *ProcessState) Success() bool
func (p *ProcessState) Sys() interface{}
func (p *ProcessState) SysUsage() interface{}
func (p *ProcessState) SystemTime() time.Duration
func (p *ProcessState) UserTime() time.Duration

## type Signal

```go
type Signal interface {
    String() string
    Signal() // 用来区分其他实现了Stringer接口的类型
}
```