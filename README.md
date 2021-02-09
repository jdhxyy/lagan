# lagan

## 介绍
基于go语言的日志库.

lagan取名来自于宜家的水龙头"拉根"。

## 功能
- 支持日志在终端实时打印
- 支持日志保存在文件
- 支持日志文件自动分割
- 支持终端交互控制日志输出级别等功能

## 示例
````go
package main

import "gitee.com/jdhxyy/lagan"

const Tag = "testlog"

func main () {
    // 日志模块载入.全局载入一次,参数是分割文件大小,默认是10M
    _ = lagan.Load(0)

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
