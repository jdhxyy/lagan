# lagan

## 介绍
基于go语言的日志库.

lagan取名来自于宜家的水龙头"拉根"。

## 开源
- [github上的项目地址](https://github.com/jdhxyy/lagan)
- [gitee上的项目地址](https://gitee.com/jdhxyy/lagan)

## 功能
- 支持日志在终端实时打印
- 支持日志保存在文件
- 支持日志不保存文件,仅终端打印
- 支持日志文件自动分割
- 支持终端交互控制日志输出级别等功能
- 支持二进制流打印
- 支持带颜色的日志打印

## 示例
````go
package main

import "github.com/jdhxyy/lagan"

const Tag = "testlog"

func main () {
    // 日志模块载入.全局载入一次,参数是分割文件大小,默认值是10M
    _ = lagan.Load(lagan.LogFileSizeDefault)

    // 默认输出级别是info,本行不会打印
    lagan.Debug(Tag, "debug test print")

    lagan.Info(Tag, "info test print")
    lagan.Warn(Tag, "warn test print")
    lagan.Error(Tag, "error test print")
}
````

输出：
````
2021/02/09 19:12:23 I/testlog: info test print
2021/02/09 19:12:23 W/testlog: warn test print
2021/02/09 19:12:23 E/testlog: error test print
````

在本地会新建log文件夹，并新建日志文件。

## 二进制流打印
````go
package main

import "github.com/jdhxyy/lagan"

const Tag = "testlog"

func main () {
    _ = lagan.Load(lagan.LogFileSizeDefault)

    arr := make([]uint8, 100)
    for i := 0; i < 100; i++ {
        arr[i] = uint8(i)
    }

    lagan.Info(Tag, "print hex")
    lagan.PrintHex(Tag, lagan.LevelInfo, arr)
}
````

输出：
````
2021/02/09 19:43:51 I/testlog: print hex
2021/02/09 19:43:51 I/testlog: 
**** : 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 
---- : -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
0000 : 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 
0010 : 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f 
0020 : 20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f 
0030 : 30 31 32 33 34 35 36 37 38 39 3a 3b 3c 3d 3e 3f 
0040 : 40 41 42 43 44 45 46 47 48 49 4a 4b 4c 4d 4e 4f 
0050 : 50 51 52 53 54 55 56 57 58 59 5a 5b 5c 5d 5e 5f 
0060 : 60 61 62 63 
````

## 终端交互
````
*******************************************
            lagan help shell             
current level:I,is pause:false
help:print help
filter_error:filter error level
filter_warn:filter warn level
filter_info:filter info level
filter_debug:filter debug level
filter_off:filter off level
pause:pause log
resume:resume log
*******************************************
````

可以在终端敲对应的命令控制日志功能。

## 颜色控制
可以使用EnableColor函数开控制打开或者关闭日志颜色，默认关闭颜色。
````go
func main() {
    _ = lagan.Load(lagan.LogFileSizeDefault)
    lagan.SetFilterLevel(lagan.LevelDebug)
    lagan.EnableColor(true)
    lagan.Debug("system", "test print:123456789")
    lagan.Info("system", "test print:123456789")
    lagan.Warn("system", "test print:123456789")
    lagan.Error("system", "test print:123456789")

    s := make([]uint8, 100)
    for i := 0; i < 100; i++ {
        s[i] = uint8(i)
    }
    lagan.PrintHex("test", lagan.LevelInfo, s)
}
````

![图片](https://user-images.githubusercontent.com/1323843/111395954-01cb2000-86f9-11eb-9cf7-c689eec1f77a.png)

## 日志文件控制
文件分割大小设置为0,则不会保存到日志文件,仅终端打印：
```python
lagan.load(0)
```
