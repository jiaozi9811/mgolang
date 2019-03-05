# Protobuf

[TOC]

## Protobuf

Protobuf即Protocol Buffers,是google开发的一种跨语言和平台的序列化数据结构的方式，是一个灵活，高效的序列化数据的协议

与xml，json相比，protobuf更小，更快，更便捷
protobuf是跨语言的，自带编辑器(protoc)

GitHub地址：<https://github.com/protocolbuffers/protobuf>

安装
https://github.com/protocolbuffers/protobuf/blob/master/src/README.md
protobuf需要一下依赖
autoconf
automake
libtool
make
g++
unzip
安装好后会生成protoc二进制文件

## Protoc

protobuf提供了protoc编辑器，用于通过定义好的.proto文件生成不同语言的代码
protoc --proto_path=IMPORT_PATH --cpp_out=DST_DIR --java_out=DST_DIR --python_out=DST_DIR --go_out=DST_DIR --ruby_out=DST_DIR --javanano_out=DST_DIR --objc_out=DST_DIR --csharp_out=DST_DIR path/to/file.proto
