// Copyright 2019-2021 The TZIOT Authors. All rights reserved.
// 日志包.可打印实时日志或者流日志
// Authors: jdh99 <jdh821@163.com>
// lagan取名来自于宜家的水龙头"拉根"

package lagan

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// FilterLevel 过滤日志级别类型
type FilterLevel uint8

// 过滤日志级别
const (
	LevelOff FilterLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

var gLevelCh = [...]byte{'O', 'D', 'I', 'W', 'E'}

// 日志颜色
// 30 黑色 31 红色 32 绿色 33 黄色 34 蓝色 35 紫红色 36 青蓝色 37 白色
var levelColor = [...]FilterLevel{30, 37, 36, 35, 31}

const (
	// LogFileSizeDefault 日志文件默认大小.单位:M字节
	LogFileSizeDefault = 10
)

var gInfoLogger, gInfoLoggerStd *log.Logger
var gIsPause = false
var gLogFileSize = 0
var gLogFileMaxSize = 0
var gLogFile *os.File = nil
var gFilterLevel FilterLevel = LevelInfo
var gIsLoad = false
var isColor = false

func init() {
	gInfoLoggerStd = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)
}

// Load 模块载入
// logFileMaxSize是日志文件切割的大小.单位:M字节.如果传入0,表示不使用日志文件
func Load(logFileMaxSize int) error {
	if gIsLoad {
		return errors.New("already load")
	}

	gLogFileMaxSize = logFileMaxSize * 1024 * 1024
	if gLogFileMaxSize <= 0 {
		gIsLoad = true
		return nil
	}
	err := createLogFile()
	if err == nil {
		gIsLoad = true
	}
	return err
}

func createLogFile() error {
	_ = os.Mkdir("log", os.ModePerm)
	file := "log/" + time.Now().Format("20060102-150405") + ".txt"
	logFile, err := os.Create(file)
	if err != nil {
		fmt.Println("create log file fail!", err.Error())
		return err
	}
	if gLogFile != nil {
		_ = gLogFile.Close()
	}
	gLogFile = logFile
	gInfoLogger = log.New(gLogFile, "", log.Ldate|log.Lmicroseconds)
	gLogFileSize = 0
	return nil
}

// SetFilterLevel 设置日志级别
func SetFilterLevel(level FilterLevel) {
	gFilterLevel = level
}

// GetFilterLevel 显示过滤日志等级
func GetFilterLevel() FilterLevel {
	return gFilterLevel
}

// EnableColor 使能日志带颜色输出
func EnableColor(enable bool) {
	isColor = enable
}

// Print 日志打印
func Print(tag string, level FilterLevel, format string, a ...interface{}) {
	if gIsPause || gFilterLevel == LevelOff || level < gFilterLevel {
		return
	}

	prefix := fmt.Sprintf("%c/%s", gLevelCh[level], tag)
	newFormat := prefix + ": " + format
	s := fmt.Sprintf(newFormat, a...)

	if gLogFileMaxSize > 0 {
		gInfoLogger.Println(s)
		gLogFileSize += len(s)
		if gLogFileSize > gLogFileMaxSize {
			_ = createLogFile()
		}
	}
	if isColor {
		gInfoLoggerStd.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, 7, 40, levelColor[level], s, 0x1B)
	} else {
		gInfoLoggerStd.Println(s)
	}
}

// PrintHex 打印16进制字节流
// tag是标记,在字节流打印之前会打印此标记
func PrintHex(tag string, level FilterLevel, bytes []uint8) {
	if gIsPause || gFilterLevel == LevelOff || level < gFilterLevel {
		return
	}

	s := "\n**** : "
	for i := 0; i < 16; i++ {
		s += fmt.Sprintf("%02x ", i)
	}
	s += "\n---- : "
	for i := 0; i < 16; i++ {
		s += fmt.Sprintf("-- ")
	}
	for i := 0; i < len(bytes); i++ {
		if i%16 == 0 {
			s += fmt.Sprintf("\n%04x : ", i)
		}
		s += fmt.Sprintf("%02x ", bytes[i])
	}

	prefix := fmt.Sprintf("%c/%s", gLevelCh[level], tag)
	newFormat := prefix + ": " + "%s"

	s1 := fmt.Sprintf(newFormat, s)

	if gLogFileMaxSize > 0 {
		gInfoLogger.Println(s1)
		gLogFileSize += len(s1)
		if gLogFileSize > gLogFileMaxSize {
			_ = createLogFile()
		}
	}
	if isColor {
		gInfoLoggerStd.Printf("%c[%d;%d;%dm%s%c[0m%s\n", 0x1B, 7, 40, levelColor[level], prefix, 0x1B, s)
	} else {
		gInfoLoggerStd.Println(s1)
	}
}

// Pause 暂停日志打印
func Pause() {
	gIsPause = true
}

// Resume 恢复日志打印
func Resume() {
	gIsPause = false
}

// IsPause 是否暂停打印
func IsPause() bool {
	return gIsPause
}

// Debug 打印debug信息
func Debug(tag string, format string, a ...interface{}) {
	Print(tag, LevelDebug, format, a...)
}

// Info 打印info信息
func Info(tag string, format string, a ...interface{}) {
	Print(tag, LevelInfo, format, a...)
}

// Warn 打印warn信息
func Warn(tag string, format string, a ...interface{}) {
	Print(tag, LevelWarn, format, a...)
}

// Error 打印error信息
func Error(tag string, format string, a ...interface{}) {
	Print(tag, LevelError, format, a...)
}
