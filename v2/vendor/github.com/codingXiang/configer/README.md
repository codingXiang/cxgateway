# Configer
封裝 github.com/spf13/viper 的套件

## 如何使用
```
go get -u github.com/codingXiang/configer
```

## 範例
### 建立 Configer
```
Config = NewConfiger()
```

### 建立 ConfigCore
```
// 參數依序為：
/// 1. 設定檔類型 (支援 yaml、yml、json、properties、ini、hcl、toml)
/// 2. 檔案名稱 (例如檔名為 config.yaml 就輸入 config)
/// 3. 後續皆為檔案路徑，可以支援多個路徑尋找檔案
var config = NewConfigerCore("yaml", "config", "./config", ".")
```

### 加入與取得 ConfigCore 到 Configer
```
// 設定 core 的 key
Config.AddCore("config", config)
// 透過 key 取得 core
Config.GetCore("config")
```

### 取得組態內容
```
// 判斷讀取 core 是否出現錯誤
if data, err := Config.GetCore("config").ReadConfig(); err == nil {
    // 取得組態裡面設定為 content 的資料
    fmt.Println(data.Get("content"))
}
```