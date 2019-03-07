# bufio

[TOC]

bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象

[bufio.go](./code/bufio.go)


bufio.flush() 会将缓存区内容写入文件，当所有写入完成后，因为缓存区会存储内容，所以需要手动flush()到文件

bufio.Available() 为buf可用容量，等于len(buf) - n