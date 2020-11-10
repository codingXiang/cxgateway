package i18n

import (
	"github.com/codingXiang/cxgateway/v2/util/e"
	"github.com/codingXiang/cxgateway/v2/util/util"
	"github.com/codingXiang/gogo-i18n"
)

const (
	Error = "error"
	Data  = "data"
)

type I18nMessageHandler struct {
	i18n util.I18nMsgInterface
}

//NewI18nMessageHandler 建立 I18nMessageHandler 實例
func NewI18nMessageHandler(moduleName string) *I18nMessageHandler {
	return &I18nMessageHandler{
		i18n: util.NewI18nMsg(moduleName),
	}
}

func (handler *I18nMessageHandler) SetCore(data gogo_i18n.GoGoi18nInterface) {
	handler.i18n.SetCore(data)
}

func (handler *I18nMessageHandler) SetModule(module string) {
	handler.i18n.SetModule(module)
}

/*
成功訊息
*/

//GetSuccess 取得成功
func (handler *I18nMessageHandler) GetSuccess(result interface{}) (int, *e.Response) {
	return e.StatusOK(handler.i18n.Get(true, nil), result)
}

//CreateSuccess 建立成功
func (handler *I18nMessageHandler) CreateSuccess(result interface{}) (int, *e.Response) {
	return e.StatusCreated(handler.i18n.Create(true, nil), result)
}

//UpdateSuccess 更新成功
func (handler *I18nMessageHandler) UpdateSuccess(result interface{}) (int, *e.Response) {
	return e.StatusCreated(handler.i18n.Update(true, nil), result)
}

//ModifySuccess 修正成功
func (handler *I18nMessageHandler) ModifySuccess(result interface{}) (int, *e.Response) {
	return e.StatusCreated(handler.i18n.Modify(true, nil), result)
}

//DeleteSuccess 刪除成功
func (handler *I18nMessageHandler) DeleteSuccess(result interface{}) (int, *e.Response) {
	return e.StatusNoContent("")
}

//AppendSuccess 加入成功
func (handler *I18nMessageHandler) AppendSuccess(module string, result interface{}) (int, *e.Response) {
	return e.StatusCreated(handler.i18n.Append(true, module, nil), result)
}

//RemoveSuccess 移除成功
func (handler *I18nMessageHandler) RemoveSuccess(module string, result interface{}) (int, *e.Response) {
	return e.StatusOK(handler.i18n.Remove(true, module, nil), result)
}

/*
錯誤訊息
*/
func (handler *I18nMessageHandler) ParameterIntError(data string, err error) *e.APIException {
	return e.StatusBadRequest(
		handler.i18n.ParamIntError(map[string]interface{}{
			Data:  data,
			Error: err.Error(),
		}),
	)
}

func (handler *I18nMessageHandler) ParameterFormatError() *e.APIException {
	return e.StatusBadRequest(
		handler.i18n.ParamFormatError(nil),
	)
}

func (handler *I18nMessageHandler) GetError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Get(false, map[string]interface{}{
		Error: err.Error(),
	}))
}

func (handler *I18nMessageHandler) GetModuleError(module string, err error) *e.APIException {
	return e.UnknownError(handler.i18n.Get(false, map[string]interface{}{
		Data:  module,
		Error: err.Error(),
	}))
}

func (handler *I18nMessageHandler) CreateError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Create(false, map[string]interface{}{
		Error: err.Error(),
	}))
}

func (handler *I18nMessageHandler) UpdateError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Update(false, map[string]interface{}{
		Error: err.Error(),
	}))
}

func (handler *I18nMessageHandler) ModifyError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Modify(false, map[string]interface{}{
		Error: err.Error(),
	}))
}

func (handler *I18nMessageHandler) DeleteError(err error) *e.APIException {
	return e.UnknownError(handler.i18n.Delete(false, map[string]interface{}{
		Error: err.Error(),
	}))
}

func (handler *I18nMessageHandler) AppendError(module string, err error) *e.APIException {
	return e.UnknownError(handler.i18n.Append(false, module, map[string]interface{}{
		Error: err.Error(),
	}))
}

func (handler *I18nMessageHandler) RemoveError(module string, err error) *e.APIException {
	return e.UnknownError(handler.i18n.Remove(false, module, map[string]interface{}{
		Error: err.Error(),
	}))
}
