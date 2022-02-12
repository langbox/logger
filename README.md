# logger

## 功能

* 实现功能
* 支持多种输出方式stdout/file
* 支持输出为json 或 plaintext
* 支持彩色输出
* 支持log rotate
*

## 使用说明

### 获取当前包的默认实例

```
var log *logrus.Logger = logger.Logger
```

### 简单使用

```
logger.Debug("dddd")
```
### 输出到文件简单配置

```
# 初始化默认配置
cfg := logger.DefaultDefinition()
logger.InitWithConfig(cfg)
logger.Debug("dddd")
```

### 自定义输出到文件配置

```
cfg := logger.Cfg{
		Writers:       "stdout,file",
		Level:         "DEBUG",
		File:          "log/chassis.log",
		FormatText:    false,
		Color:         false,
		RotateDate:    1,
		RotateSize:    10,
		BackupCount:   7,
		Compress: 	   true,
	}

logger.InitWithConfig(&cfg)
logger.Debug("dddd")

```
