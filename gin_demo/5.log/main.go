package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)
/**
	日志处理
 */

const (
	PORT       = ":8080"
	APP_NAME   = "ginDemo"
	APP_SECRET = "6YJSuc50uJ18zj45"
	API_EXPIRY = "120"

	Log_FILE_PATH = "C:/Users/Administrator/Desktop/ginlogs"
	LOG_FILE_NAME = "system.log"
)

func main() {

	//这个是开启日志颜色高亮显示功能。相对的，也就有禁止日志颜色高亮显示功能，那就是gin.DisableConsoleColor()这个函数
	//日志高亮是啥意思呢？就是在请求具体的接口信息的时候，会在控制台高亮显示请求状态和方法的属性值。
	//gin.ForceConsoleColor()

	//当需要把日志记录到文本文件中，就没有必要开启日志颜色高亮的功能了
	gin.DisableConsoleColor()

	r := gin.Default()

	r.Use(LogrusLoggerToFileRotateMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(PORT)
}
//这个中间件生成的日志文件不会进行滚动记录
func LogrusLoggerToFileMiddleware() gin.HandlerFunc {

	logFilePath := Log_FILE_PATH
	logFileName := LOG_FILE_NAME

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//这里最好做一个判断，文件是否存在，不存在就创建
	if err != nil {
		fmt.Println("err", err)
		return nil
	}

	//实例化logrus
	logrusInstance := logrus.New()
	//设置输出
	logrusInstance.Out = src
	//设置日志级别
	logrusInstance.SetLevel(logrus.DebugLevel)
	//设置日志格式:普通文本方式
	/*
	logrusInstance.SetFormatter(&logrus.TextFormatter{
		//设置方便阅读的时间格式
		TimestampFormat: "2006-01-02 15:04:05",
	})
	*/
	//设置日志格式：以JSON方式记录
	logrusInstance.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		//获取执行时间
		latencyTime := endTime.Sub(startTime)
		//获取请求方式
		reqMethod := c.Request.Method
		//获取请求路由
		reqUri := c.Request.RequestURI
		//获取状态码
		statusCode := c.Writer.Status()
		//获取请求IP
		clientIP := c.ClientIP()

		/*
		//日志格式，这种配合普通文本输出
		logrusInstance.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
		*/
		//这种配合JSON格式输出
		logrusInstance.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}

//这个中间件生成的日志会进行每天滚动记录
func LogrusLoggerToFileRotateMiddleware() gin.HandlerFunc {

	logFilePath := Log_FILE_PATH
	logFileName := LOG_FILE_NAME

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//这里最好做一个判断，文件是否存在，不存在就创建
	if err != nil {
		fmt.Println("err", err)
		return nil
	}

	//实例化logrus
	logrusInstance := logrus.New()
	//设置输出
	logrusInstance.Out = src
	//设置日志级别
	logrusInstance.SetLevel(logrus.DebugLevel)

	//设置rotatelogs
	logWriter, err := rotatelogs.New(
		//分割后的文件名称
		fileName+".%Y%m%d.log",
		//生成软链接，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		//设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//新增Hook
	logrusInstance.AddHook(lfHook)

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		//获取执行时间
		latencyTime := endTime.Sub(startTime)
		//获取请求方式
		reqMethod := c.Request.Method
		//获取请求路由
		reqUri := c.Request.RequestURI
		//获取状态码
		statusCode := c.Writer.Status()
		//获取请求IP
		clientIP := c.ClientIP()

		logrusInstance.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}

//将日志记录到MongoDB
func LogrusLoggerToMongoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//将日志记录到ES
func LogrusLoggerToESMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//将日志记录到MQ
func LogrusLoggerToMQMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

/**
	将日志记录到文本文件中
 */
func test1() {
	//当需要把日志记录到文本文件中，就没有必要开启日志颜色高亮的功能了
	gin.DisableConsoleColor()

	//定义将日志输出到指定的文件中
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	//定义日志输出格式
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \" %s %s %s %d \" %s\" %s\" %s \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}
