package i18n

import (
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/codingXiang/gogo-i18n"
)

type (
	I18nMessageHandlerInterface interface {
		SetCore(data gogo_i18n.GoGoi18nInterface)
		SetModule(module string)
		GetSuccess(result interface{}) (int, *e.Response)
		CreateSuccess(result interface{}) (int, *e.Response)
		UpdateSuccess(result interface{}) (int, *e.Response)
		ModifySuccess(result interface{}) (int, *e.Response)
		DeleteSuccess(result interface{}) (int, *e.Response)
		AppendSuccess(module string, result interface{}) (int, *e.Response)
		RemoveSuccess(module string, result interface{}) (int, *e.Response)
		ParameterFormatError() *e.APIException
		ParameterIntError(data string, err error) *e.APIException
		GetError(err error) *e.APIException
		GetModuleError(module string, err error) *e.APIException
		CreateError(err error) *e.APIException
		UpdateError(err error) *e.APIException
		ModifyError(err error) *e.APIException
		DeleteError(err error) *e.APIException
		AppendError(module string, err error) *e.APIException
		RemoveError(module string, err error) *e.APIException
	}
	i18nMessageHandler struct {
		i18n util.I18nMsgInterface
	}
)

//NewI18nMessageHandler 建立 i18nMessageHandler 實例
func NewI18nMessageHandler(moduleName string) I18nMessageHandlerInterface {
	return &i18nMessageHandler{
		i18n: util.NewI18nMsg(moduleName),
	}
}

func (handler *i18nMessageHandler) SetCore(data gogo_i18n.GoGoi18nInterface) {
	handler.i18n.SetCore(data)
}

func (handler *i18nMessageHandler) SetModule(module string) {
	handler.i18n.SetModule(module)
}

/*
成功訊息
*/

//GetSuccess 取得成功
func (handler *i18nMessageHandler) GetSuccess(result interface{}) (int, *e.Response) {
	return e.StatusSuccess(handler.i18n.Get(true, nil), result)
}

//CreateSuccess 建立成功
func (handler *i18nMessageHandler) CreateSuccess(result interface{}) (int, *e.Response) {
	return e.StatusCreated(handler.i18n.Create(true, nil), result)
}

//UpdateSuccess 更新成功
func (handler *i18nMessageHandler) UpdateSuccess(result interface{}) (int, *e.Response) {
	return e.StatusCreated(handler.i18n.Update(true, nil), result)
}

//ModifySuccess 修正成功
func (handler *i18nMessageHandler) ModifySuccess(result interface{}) (int, *e.Response) {
	return e.StatusCreated(handler.i18n.Modify(true, nil), result)
}

//DeleteSuccess 刪除成功
func (handler *i18nMessageHandler) DeleteSuccess(result interface{}) (int, *e.Response) {
	return e.StatusNoContent("")
}

//AppendSuccess 加入成功
func (handler *i18nMessageHandler) AppendSuccess(module string, result interface{}) (int, *e.Response) {
	return e.StatusCreated(handler.i18n.Append(true, module, nil), result)
}

//RemoveSuccess 移除成功
func (handler *i18nMessageHandler) RemoveSuccess(module string, result interface{}) (int, *e.Response) {
	return e.StatusSuccess(handler.i18n.Remove(true, module, nil), result)
}

/*
錯誤訊息
*/
func (handler *i18nMessageHandler) ParameterIntError(data string, err error) *e.APIException {
	return e.ParameterError(
		handler.i18n.ParamIntError(map[string]interface{}{
			"data":  data,
			"error": err.Error(),
		}),
	)
}

func (handler *i18nMessageHandler) ParameterFormatError() *e.APIException {
	return e.ParameterError(
		handler.i18n.ParamRequiredError(nil),
	)
}

func (handler *i18nMessageHandler) GetError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Get(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) GetModuleError(module string, err error) *e.APIException {
	return e.UnknownError(handler.i18n.Get(false, map[string]interface{}{
		"data":  module,
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) CreateError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Create(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) UpdateError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Update(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) ModifyError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Modify(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) DeleteError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Delete(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) AppendError(module string, err error) *e.APIException {
	return e.UnknownError(handler.i18n.Append(false, module, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) RemoveError(module string, err error) *e.APIException {
	return e.UnknownError(handler.i18n.Remove(false, module, map[string]interface{}{
		"error": err.Error(),
	}))
}
