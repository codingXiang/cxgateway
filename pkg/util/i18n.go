package util

import (
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/gogo-i18n"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/language"
)

func GetI18nData(c *gin.Context) (gogo_i18n.GoGoi18nInterface) {
	data, _ := c.Get("i18n")
	return data.(*gogo_i18n.GoGoi18n)
}

func LoadI18nDataFromDatabase(db *gorm.DB, storeType string, storePath string) {
	//設定多語系
	i18ns := make([]*gogo_i18n.GoGoi18nMessage, 0)
	i18nMsgs := make([]gogo_i18n.GoGoi18nMessageInterface, 0)
	logger.Log.Info("load i18n data from database...")
	if err := db.Find(&i18ns).Error; err == nil {
		for _, i18n := range i18ns {
			msg := createI18nMessageFromRecord(i18n)
			i18nMsgs = append(i18nMsgs, msg)
		}
		logger.Log.Info("save i18n data to", storePath, "type = ", storeType)
		gogo_i18n.StoreDataToFile(storeType, storePath, i18nMsgs)
	} else {
		logger.Log.Error("load i18n data from database failed, err =", err.Error())
	}
}

func createI18nMessageFromRecord(i18n *gogo_i18n.GoGoi18nMessage) gogo_i18n.GoGoi18nMessageInterface {
	logger.Log.Debug("load i18n record =", i18n)
	switch i18n.Language {
	case "zh-Hant":
		return gogo_i18n.NewGoGoi18nMessage(language.TraditionalChinese, i18n.Key, i18n.Value)
	case "zh-Hans":
		return gogo_i18n.NewGoGoi18nMessage(language.SimplifiedChinese, i18n.Key, i18n.Value)
	case "es":
		return gogo_i18n.NewGoGoi18nMessage(language.English, i18n.Key, i18n.Value)
	default:
		return gogo_i18n.NewGoGoi18nMessage(language.TraditionalChinese, i18n.Key, i18n.Value)
	}
}


//-------------------------------//
// error type
const (
	ERRMSG = "errMsg"
	//參數錯誤
	PARAM    = "parameter"
	FORMAT   = "format"
	INT      = "int"
	required = "required"
)

//成功與否
const (
	REASON  = "reason"
	SUCCESS = "success"
	FAILED  = "failed"
)

//標點符號
const (
	COMMA  = "comma"
	PERIOD = "period"
	SPACE  = " "
)

//CRUD 方法
const (
	CREATE = "create"
	UPDATE = "update"
	MODIFY = "modify"
	GET    = "get"
	DELETE = "delete"
	APPEND = "append"
	REMOVE = "remove"
)

//模組
const (
	MODULE = "module"
)

type (
	I18nMsgInterface interface {
		//基本方法
		SetCore(data gogo_i18n.GoGoi18nInterface)
		SetModule(module string) *I18nMsg
		//error handler
		ParamFormatError(data map[string]interface{}) string
		ParamIntError(data map[string]interface{}) string
		ParamRequiredError(data map[string]interface{}) string
		//取得模組
		GetModule(name string) string
		//http方法
		Get(isSuccess bool, data map[string]interface{}) string
		Create(isSuccess bool, data map[string]interface{}) string
		Update(isSuccess bool, data map[string]interface{}) string
		Modify(isSuccess bool, data map[string]interface{}) string
		Delete(isSuccess bool, data map[string]interface{}) string
		Append(isSuccess bool, module string, data map[string]interface{}) string
		Remove(isSuccess bool, module string, data map[string]interface{}) string
	}

	I18nMsg struct {
		Module string
		Core   gogo_i18n.GoGoi18nInterface
	}

	I18nMsgContent struct {
		Data  string `json:"data"`
		Error string `json:"error"`
	}
)

func NewI18nMsg(module string) I18nMsgInterface {
	return &I18nMsg{
		Module: module,
	}
}

func (i *I18nMsg) SetCore(data gogo_i18n.GoGoi18nInterface) {
	i.Core = data
}

func (i *I18nMsg) GetModule(name string) string {
	return i.Core.GetMessage(MODULE+"."+name, nil)
}

/*
	通用錯誤訊息
 */

func (i *I18nMsg) SetModule(module string) *I18nMsg {
	i.Module = module
	return i
}

func (i *I18nMsg) ParamFormatError(data map[string]interface{}) string {
	msg := ERRMSG + "." + PARAM + "." + FORMAT
	return i.Core.GetMessage(msg, data)
}

func (i *I18nMsg) ParamIntError(data map[string]interface{}) string {
	msg := ERRMSG + "." + PARAM + "." + INT
	return i.Core.GetMessage(msg, data)
}

func (i *I18nMsg) ParamRequiredError(data map[string]interface{}) string {
	msg := ERRMSG + "." + PARAM + "." + required
	return i.Core.GetMessage(msg, data)
}

/*
	http 方法
 */

func (i *I18nMsg) Get(isSuccess bool, data map[string]interface{}) string {
	return i.handlMsg(GET, isSuccess, data)
}

func (i *I18nMsg) Create(isSuccess bool, data map[string]interface{}) string {
	return i.handlMsg(CREATE, isSuccess, data)
}

func (i *I18nMsg) Update(isSuccess bool, data map[string]interface{}) string {
	return i.handlMsg(UPDATE, isSuccess, data)
}

func (i *I18nMsg) Modify(isSuccess bool, data map[string]interface{}) string {
	return i.handlMsg(MODIFY, isSuccess, data)
}

func (i *I18nMsg) Delete(isSuccess bool, data map[string]interface{}) string {
	return i.handlMsg(DELETE, isSuccess, data)
}

func (i *I18nMsg) Append(isSuccess bool, module string, data map[string]interface{}) string {
	if data == nil {
		data = map[string]interface{}{}
	}
	data["object"] = i.GetModule(module)
	return i.handlMsg(APPEND, isSuccess, data)
}

func (i *I18nMsg) Remove(isSuccess bool, module string, data map[string]interface{}) string {
	if data == nil {
		data = map[string]interface{}{}
	}
	data["object"] = i.GetModule(module)
	return i.handlMsg(REMOVE, isSuccess, data)
}

func (i *I18nMsg) handlMsg(method string, isSuccess bool, data map[string]interface{}) string {
	info := map[string]interface{}{}
	if data["object"] != nil {
		info["object"] = data["object"]
	}
	if data["data"] != nil {
		info["data"] = i.GetModule(data["data"].(string))
	} else {
		info["data"] = i.GetModule(i.Module)
	}
	msg := i.Core.GetMessage(method, info)
	if isSuccess {
		msg += SPACE + i.Core.GetMessage(SUCCESS, nil)
	} else {
		msg += SPACE + i.Core.GetMessage(FAILED, nil) + i.Core.GetMessage(COMMA, nil) + i.Core.GetMessage(REASON, data)
	}
	return msg + i.Core.GetMessage(PERIOD, nil)
}
