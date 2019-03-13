# prometheus

<https://prometheus.io>
<https://github.com/prometheus>

$ go get github.com/prometheus/prometheus/cmd/...
$ prometheus --config.file=your_config.yml

```shell
$ mkdir -p $GOPATH/src/github.com/prometheus
$ cd $GOPATH/src/github.com/prometheus
$ git clone https://github.com/prometheus/prometheus.git
$ cd prometheus
$ make build
$ ./prometheus --config.file=your_config.yml
```

## 四种数据类型

- Counter
- Gauge 常规数值，例如 温度变化、内存使用变化
- Histogram 直方图
- Summary 常用于跟踪事件发生的规模