# archive/zip

[zip.go](./zip.go)
[unzip.go](./unzip.go)

## type FileHeader

- func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)
- func (h *FileHeader) FileInfo() os.FileInfo
- func (h *FileHeader) Mode() (mode os.FileMode)
- func (h *FileHeader) SetMode(mode os.FileMode)
- func (h *FileHeader) ModTime() time.Time
- func (h *FileHeader) SetModTime(t time.Time)

```go
type FileHeader struct {
    // Name是文件名，它必须是相对路径，不能以设备或斜杠开始，只接受'/'作为路径分隔符
    Name string
    CreatorVersion     uint16
    ReaderVersion      uint16
    Flags              uint16
    Method             uint16
    ModifiedTime       uint16 // MS-DOS时间
    ModifiedDate       uint16 // MS-DOS日期
    CRC32              uint32
    CompressedSize64   uint64
    UncompressedSize64 uint64
    Extra              []byte
    ExternalAttrs      uint32 // 其含义依赖于CreatorVersion
    Comment            string
}
```

## type File

- func (f *File) DataOffset() (offset int64, err error)
- func (f *File) Open() (rc io.ReadCloser, err error)

## type Reader

- func NewReader(r io.ReaderAt, size int64) (*Reader, error)
- type ReadCloser
- func OpenReader(name string) (*ReadCloser, error)
- func (rc *ReadCloser) Close() error

## type Writer

- func NewWriter(w io.Writer) *Writer
- func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)
- func (w *Writer) Create(name string) (io.Writer, error)
- func (w *Writer) Close() error
