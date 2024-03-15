package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var loglevelmap = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"trace": logrus.TraceLevel,
}

const birthdayGolang = "2006-01-02 15:04:05"

// LogTemp 记录临时日志，用于日志模块初始化之前
func LogTemp(format string, a ...interface{}) (n int, err error) {
	fileLineInfo := ":"
	_, file, line, ok := runtime.Caller(1) // 只需找到上一层调用方
	if ok {
		tmp := strings.Split(file, "/")
		fileLineInfo = fmt.Sprintf("%s:%d", tmp[len(tmp)-1], line)
	}

	// "2006-01-02 15:04:05" golang 格式化时间需写死这个。据说是golang诞生的时间
	prefix := fmt.Sprintf("[%s] [%s] ", time.Now().Format(birthdayGolang), fileLineInfo)

	return fmt.Printf(prefix+format+"\n", a...)
}

func getFileLineInfo() string {
	fileLineInfo := "[:] "
	_, file, line, ok := runtime.Caller(3) // 上层调用者为当前实例的成员函数，要再往上一层是该文件的Error/Debug等函数，再往上一层才能找到调用者的位置
	if ok {
		tmp := strings.Split(file, "/")
		fileLineInfo = fmt.Sprintf("[%s:%d] ", tmp[len(tmp)-1], line) // 只打印文件名，不打印全路径
		// fileLineInfo = fmt.Sprintf("[%s:%d] ", file, line)
	}

	return fileLineInfo
}

type commonLoger struct {
	log *logrus.Logger
}

// 日志模块实例，外部调用
var defaultLogger = &commonLoger{}

// 初始化日志实例
func (c *commonLoger) Initialize(level string, logFilePath string, output string) error {
	c.log = logrus.New()
	// c.log.SetReportCaller(true) // 显示行号等信息，打印的是本文件的行数而不是调用者的位置，没什么用
	c.SetLevel(level)
	c.log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: birthdayGolang,
	})

	// c.log.SetFormatter(&logrus.JSONFormatter{})

	switch output {
	case "stdout":
		c.log.SetOutput(os.Stdout)
	case "stderr":
		c.log.SetOutput(os.Stderr)
	default:
		// 检查日志路径是否存在并且是目录，不存在则创建
		dir := filepath.Dir(logFilePath)
		fDir, err := os.Stat(dir)
		if err == nil {
			if !fDir.IsDir() {
				return fmt.Errorf("log path [%s] exists, but is not dir", dir)
			}
		} else {
			if os.IsNotExist(err) {
				err = os.MkdirAll(dir, 0755)
				return err
			}
		}
		logFileCut := LogFileCut(logFilePath)
		writers := []io.Writer{
			logFileCut,
			os.Stdout}

		// 输出到控制台，方便定位到那个文件
		fileAndStdoutWriter := io.MultiWriter(writers...)
		gin.DefaultWriter = fileAndStdoutWriter

		c.log.SetOutput(fileAndStdoutWriter)
	}

	return nil
}

func (c *commonLoger) SetLogger(adapterName string, configs ...string) error {
	return nil
}

func (c *commonLoger) SetLevel(level string) {
	c.log.SetLevel(loglevelmap[strings.ToLower(level)])
}

func (c *commonLoger) Panic(format string, args ...interface{}) {
	c.log.Panicf(getFileLineInfo()+format, args...)
}

func (c *commonLoger) Fatal(format string, args ...interface{}) {
	c.log.Fatalf(getFileLineInfo()+format, args...)
}

func (c *commonLoger) Error(format string, args ...interface{}) {
	c.log.Errorf(getFileLineInfo()+format, args...)
}

func (c *commonLoger) Warn(format string, args ...interface{}) {
	c.log.Warnf(getFileLineInfo()+format, args...)
}

func (c *commonLoger) Info(format string, args ...interface{}) {
	c.log.Infof(getFileLineInfo()+format, args...)
}

func (c *commonLoger) Debug(format string, args ...interface{}) {
	c.log.Debugf(getFileLineInfo()+format, args...)
}

func (c *commonLoger) Trace(format string, args ...interface{}) {
	c.log.Tracef(getFileLineInfo()+format, args...)
}

// **************************************************
// 加入口函数，方便直接通过包名来调用

// Initialize 初始化
func Initialize(level string, logFilePath string, output string) error {
	return defaultLogger.Initialize(level, logFilePath, output)
}

// SetLevel 日志级别
// 0: Panic；1: Fatal；2: Error；3: Warn；4: Info；5: Debug；6: Trace
func SetLevel(level string) {
	defaultLogger.log.SetLevel(loglevelmap[level])
}

// 格式化打印

// Panic Panic
func Panic(format string, args ...interface{}) {
	defaultLogger.Panic(format, args...)
}

// Fatal Fatal
func Fatal(format string, args ...interface{}) {
	defaultLogger.Fatal(format, args...)
}

// Error Error
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// Warn Warn
func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// Info Info
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Debug Debug
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Trace Trace
func Trace(format string, args ...interface{}) {
	defaultLogger.Trace(format, args...)
}

// 配置日志切割
// LogFileCut 日志文件切割
func LogFileCut(filePath string) *rotatelogs.RotateLogs {
	// 配置日志分割
	logFileName := path.Join(filePath, "%Y%m%d.log")
	logier, err := rotatelogs.New(
		// 切割后日志文件名称
		logFileName,
		//rotatelogs.WithLinkName(Current.LogDir),   // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		//rotatelogs.WithRotationCount(3),
		//rotatelogs.WithRotationTime(time.Minute), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  logier,
		logrus.FatalLevel: logier,
		logrus.DebugLevel: logier,
		logrus.WarnLevel:  logier,
		logrus.ErrorLevel: logier,
		logrus.PanicLevel: logier,
	},
		// 设置分割日志样式
		&logrus.TextFormatter{})
	logrus.AddHook(lfHook)
	return logier
}
