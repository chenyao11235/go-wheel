package logger

import (
    "fmt"
    "github.com/natefinch/lumberjack"
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "log"
    "time"
)

// 单例
var Log *zap.SugaredLogger

func InitLog() {
    now := time.Now()
    hook := &lumberjack.Logger{
        Filename:   fmt.Sprintf("%s/%04d%02d%02d%02d%02d%02d", viper.GetString("logger_dir"), now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()), //filePath
        MaxSize:    viper.GetInt("log_rotate_size"),                                                                                                                       // megabytes
        MaxBackups: viper.GetInt("log_backup_count"),
        MaxAge:     viper.GetInt("log_rotate_days"), //days
        Compress:   false,                           // disabled by default
    }
    defer hook.Close()
    /*zap 的 Config 非常的繁琐也非常强大，可以控制打印 log 的所有细节，因此对于我们开发者是友好的，有利于二次封装。
      但是对于初学者则是噩梦。因此 zap 提供了一整套的易用配置，大部分的姿势都可以通过一句代码生成需要的配置。
    */
    enConfig := zap.NewProductionEncoderConfig() //生成配置

    // 时间格式z
    enConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 指定日期格式

    level := zap.InfoLevel
    w := zapcore.AddSync(hook)
    core := zapcore.NewCore(
        zapcore.NewConsoleEncoder(enConfig), //编码器配置
        w,                                   //打印到控制台和文件
        level,                               //日志等级
    )

    logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
    _log := log.New(hook, "", log.LstdFlags)
    Log = logger.Sugar() // 初始化全局的Logger
    _log.Println("Start...")
}
