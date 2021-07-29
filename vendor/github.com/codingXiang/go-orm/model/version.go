package model

type (
	//VersionInterface 版本參數介面
	VersionInterface interface {
		GetVersion() string
	}
	Version struct {
		Version string `json:"version" gorm:"primary_key;comment:'版本號碼'"` //1.0.0
	}
)

//GetVersion 取得版本
func (v *Version) GetVersion() string {
	return v.Version
}
