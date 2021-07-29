package model

type (
	//RedisInterface Redis設定參數介面
	RedisInterface interface {
		GetURL() string
		GetPort() int
		GetPassword() string
		GetDB() int
	}
	//Redis : Redis設定參數
	Redis struct {
		URL      string `yaml:"url"`      //Redis Server 位置
		Port     int    `yaml:"port"`     //Redis Server 的 Port
		Password string `yaml:"password"` //Redis Server 的密碼
		DB       int    `yaml:"db"`       //Redis Server 指定 DB
	}
)

//GetURL 取得連線位置
func (r *Redis) GetURL() string {
	return r.URL
}

//GetPort 取得 Port
func (r *Redis) GetPort() int {
	return r.Port
}

//GetPassword 取得密碼
func (r *Redis) GetPassword() string {
	return r.Password
}

//GetDB 取得連線 db
func (r *Redis) GetDB() int {
	return r.DB
}
