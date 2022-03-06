package logger

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	File           = ""
	Level          = "debug"
	RotatePolicy   = "size" // size/daily/time
	RotateTime     = 0      // Hour
	RotateSize     = 0      // megabytes
	MaxBackup      = 0      // number
	MaxAge         = 0      // days
	FormatText     = false
	FormatColor    = false
	FormatCompress = false
)

//Cfg is the struct for log information
type Cfg struct {
	Level        string // 日志级别
	File         string // 文件路径
	RotatePolicy string // 分割策略 size/daily/time
	RotateTime   int    // 分割日期(单位 小时)
	RotateSize   int    // 分割大小(单位 M)
	MaxBackup    int    // 备份文件数目
	MaxAge       int    // 文件有效期(单位 天)
	FormatText   bool   // 日志格式
	FormatColor  bool   // 强制彩色
	FormatReport bool   // 是否显示 calling method
	// FormatCompress bool   // 是否压缩
}

// Logger is the global variable
// var Logger *logrus.Logger
var Logger = logrus.New()

//InitWithConfig 初始化
func InitWithConfig(def *Cfg) (*logrus.Logger, error) {
	err := initWithConfig(Logger, def)
	return Logger, err
}

func initWithConfig(logger *logrus.Logger, def *Cfg) error {
	logger.SetReportCaller(def.FormatReport)

	//  只输出不低于当前级别是日志数据
	levelText := Level
	if def.Level != "" {
		levelText = def.Level
	}

	level, err := logrus.ParseLevel(levelText)
	if err != nil {
		return err
	}
	logger.SetLevel(level)

	// 输出日志格式
	var formatter logrus.Formatter
	formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}

	if def.FormatText {
		formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     def.FormatColor,
		}
	}

	logger.SetFormatter(formatter)

	if def.File == "" {
		return nil
	}

	var file io.Writer
	fileName := def.File
	maxAge := MaxAge
	if def.MaxAge > 0 {
		maxAge = def.MaxAge
	}
	maxBackup := MaxBackup
	if def.MaxBackup > 0 {
		maxBackup = def.MaxBackup
	}

	// 分割策略
	if def.RotatePolicy == "" || def.RotatePolicy == RotatePolicy { // 基于大小分割
		rotateSize := RotateSize
		if def.RotateSize > 0 {
			rotateSize = def.RotateSize
		}
		file = &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    rotateSize,     // megabytes
			MaxBackups: maxBackup,      // count
			MaxAge:     maxAge,         //days
			Compress:   FormatCompress, // disabled by default
		}
	} else { // 基于时间分割
		rotateTime := RotateTime
		fileNameAll := fileName + ".%Y%m%d%H%M"
		if def.RotatePolicy == "daily" {
			rotateTime = 24
			fileNameAll = fileName + ".%Y%m%d"
		}

		if maxBackup > 0 {
			maxAge = 0
		}

		file, err = rotatelogs.New(
			fileNameAll,
			rotatelogs.WithLinkName(fileName),
			rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(rotateTime)*time.Hour),
			rotatelogs.WithRotationCount(uint(maxBackup)),
		)
		if err != nil {
			return err
		}
	}

	// 设置 output
	fileAndStdoutWriter := io.MultiWriter(file, os.Stdout)
	logger.SetOutput(fileAndStdoutWriter)
	return nil
}

// Trace Trace
func Trace(args ...interface{}) {
	Logger.Trace(args...)
}

// Tracef Tracef
func Tracef(format string, args ...interface{}) {
	Logger.Tracef(format, args...)
}

// Debug Debug
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// Debugf Debugf
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Info Info
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Infof Infof
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Warn Warn
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Warnf Warnf
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Error Error
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Errorf Errorf
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Fatal Log with os.exit(1)
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Fatalf Log with os.exit(1)
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

// Panic Log with panic
func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

// Panicf Log with panic
func Panicf(format string, args ...interface{}) {
	Logger.Panicf(format, args...)
}
