# go modules

[TOC]

- 一个开关环境变量：GO111MODULE
- 5个辅助环境变量：GOPROXY、GONOPROXY、GOSUMDB、GONOSUMDB和GOPRIVATE
- 两个辅助概念：go module proxy和go checksum database
- 两个个主要文件： go.mod和go.sum
- 一个主要管理子命令：go mod

## GO111MODULE

- auto：项目包含go.mod文件时启用go modules
- on：
- off： 禁用go modules

## GOPROXY

主要用于设置go模块的代理
- 值是以英文逗号分割的go module proxy列表。go env -w GOPROXY=https://goproxy.cn,direct
其实值列表中的 “direct” 为特殊指示符，用于指示 Go 回源到模块版本的源地址去抓取(比如 GitHub 等)，当值列表中上一个 Go module proxy 返回 404 或 410 错误时，Go 自动尝试列表中的下一个，遇见 “direct” 时回源，遇见 EOF 时终止并抛出类似 “invalid version: unknown revision...” 的错误

## GOSUMDB

值是一个go checksum database，用于go在拉取模块版本时保证拉取到的模块版本数据未经篡改









