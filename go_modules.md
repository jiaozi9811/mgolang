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

## 迁移项目至 Go Modules
- 第一步: 升级到 Go 1.13。
- 第二步: 让 GOPATH 从你的脑海中完全消失，早一步踏入未来。
-- 修改 GOBIN 路径（可选）：go env -w GOBIN=$HOME/bin。
-- 打开 Go modules：go env -w GO111MODULE=on。
-- 设置 GOPROXY：go env -w GOPROXY="https://goproxy.cn,direct" # 在中国是必须的，因为它的默认值被墙了。
- 第三步(可选): 按照你喜欢的目录结构重新组织你的所有项目。
- 第四步: 在你项目的根目录下执行 go mod init <OPTIONAL_MODULE_PATH> 以生成 go.mod 文件。
- 第五步: 想办法说服你身边所有的人都去走一下前四步。

## go mod命令
- download	download modules to local cache(下载依赖包)
- edit	edit go.mod from tools or scripts（编辑go.mod
- graph	print module requirement graph (打印模块依赖图)
- init	initialize new module in current directory（在当前目录初始化mod）
- tidy	add missing and remove unused modules(拉取缺少的模块，移除不用的模块)
- vendor	make vendored copy of dependencies(将依赖复制到vendor下)
- verify	verify dependencies have expected content (验证依赖是否正确）
- why		explain why packages or modules are needed(解释为什么需要依赖)

go.mod提供了module、require、replace、和exclude四个命令
- module 语句指定包的名字（路径）
- require 语句指定的依赖项模块
- replace 语句可以替换依赖项模块
- exclude 语句可以忽略依赖项模块

**
- go list -m -u all 来检查可以升级的package
- go get -u need-upgrade-package 升级后会将新的依赖版本更新到go.mod * 
- go get -u 升级所有依赖
**

## go get 升级

运行 go get -u 将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
运行 go get -u=patch 将会升级到最新的修订版本
运行 go get package@version 将会升级到指定的版本号version
运行go get如果有版本的更改，那么go.mod文件也会更改





