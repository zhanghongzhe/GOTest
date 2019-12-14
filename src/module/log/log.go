package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
var log = logrus.New()

func init() {
	// 设置日志格式为json格式
	//log.SetFormatter(&log.JSONFormatter{})
	log.Formatter = &logrus.JSONFormatter{}

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	//log.SetOutput(os.Stdout)
	log.Out = os.Stdout

	// 设置日志级别为warn以上
	//log.SetLevel(log.InfoLevel)
	log.Level = logrus.InfoLevel
}

func Debug(args ...interface{})  {
	log.Debug(args)
}
func Info(args ...interface{})  {
	log.Info(args)
}
func Warn(args ...interface{})  {
	log.Warn(args)
}
func Error(args ...interface{})  {
	log.Error(args)
}
func Fatal(args ...interface{})  {
	log.Fatal(args)
}
func Panic(args ...interface{})  {
	log.Panic(args)
}