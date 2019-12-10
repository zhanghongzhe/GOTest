package logTest

import (
	"github.com/golang/glog"
	"io"
	"log"
	"os"
)

func TestgLog()  {
	//初始化命令行参数
	//flag.Parse()

	//glog 这个是google的日志框架
	glog.Info("hello, glog")
	glog.Warning("warning glog")
	glog.Error("error glog")

	glog.Infof("info %d", 1)
	glog.Warningf("warning %d", 2)
	glog.Errorf("error %d", 3)

	glog.V(3).Infoln("info with v 3")
	glog.V(2).Infoln("info with v 2")
	glog.V(1).Infoln("info with v 1")
	glog.V(0).Infoln("info with v 0")

	// 退出时调用，确保日志写入文件中
	glog.Flush()
}

func TestLog()  {
	//自带Log测试
	log.SetFlags(log.Ldate|log.Lshortfile) //定制日志的抬头信息
	log.Println("Println")
	log.Fatal("Fatal")
}

func TestLogger()  {
	//Log封装测试
	Info.Println("Info")
	Error.Println("Error")
}

var (
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

func init()  {
	errFile,err := os.OpenFile(logPath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil{
		log.Fatalln("打开日志失败：", err)
	}

	Info = log.New(os.Stdout, "Info:", log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error", log.Ldate | log.Ltime | log.Lshortfile)
}