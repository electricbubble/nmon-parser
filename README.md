# nmon-parser

以相对友好的方式展示 nmon 的各项数值。

## 安装
```
go get -u github.com/electricbubble/nmon-parser
```

## 使用
通过文件解析
```
nmon, err := nmonparser.ParseNmonByFilename("/xxx.nmon")
```

通过 `io.Reader` 解析
```
nmon, err := nmonparser.ParseNmonByReader(r)
```


获取 nmon 数据的全部分类
```
fmt.Println(nmon.GetSeriesClass())
```
实际输出结果根据不同系统环境有所差异， 大体上分类格式如下
```
[AAA BBBP CPU_ALL DGBUSY DGREAD DGSIZE DGWRITE DGXFER DISKBSIZE DISKBUSY DISKREAD DISKWRITE DISKXFER JFSFILE MEM NET NETPACKET PROC VM ZZZZ CPU001 CPU002 CPU003 CPU004]
```


获取指定分类的数据
```
sl := nmon.GetSeriesLine("CPU_ALL")
count := sl.Len()
for i := 0; i < count; i++ {
    fmt.Println(sl.Get(i))
}
```
实际输出结果根据不同系统环境有所差异， 大体上 `CPU_ALL` 的输出格式如下
```
CPU Total xxx,User%,Sys%,Wait%,Idle%,Busy,CPUs
T0001,0.9,2.2,0.0,96.9,,4
T0002,1.0,0.7,0.4,97.9,,4
T0003,1.0,0.6,0.6,97.8,,4
T0004,14.2,7.6,0.0,78.1,,4
T0005,13.9,6.7,0.0,79.4,,4
T....,13.3,6.6,2.9,77.2,,4
T....,2.9,19.4,0.0,77.6,,4
T....,1.0,1.2,0.0,97.9,,4
T0069,0.7,0.6,0.0,98.7,,4
T0070,2.3,10.1,0.1,87.5,,4
T0071,0.6,0.4,0.2,98.8,,4
T0072,1.1,0.8,0.0,98.1,,4
```


---
在此基础上，做了个 [Web版 解析工具](https://github.com/ElectricBubble/nmon-analyser-releases)
（🤕其实就是因为不想用 `Excel` 版本的 `nmon Analyser` ，
解析慢，还有文件大小限制）