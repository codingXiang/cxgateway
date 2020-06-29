# logger

## 如何使用
### 導入
```
 go get -u --insecure github.com/codingXiang/go-logger
```
### 使用
創建實例

```
logger.Log = logger.NewLogger(logger.Logger{
	Level: "debug",
	Format: "json",
})
//Debug
logger.Log.Debug("test")
```
