package e

import (
	"github.com/codingXiang/cxgateway/pkg/util"
)

type (
	I18nMessageHandlerInterface interface {
		GetSuccess(result interface{}) (int, *Response)
		CreateSuccess(result interface{}) (int, *Response)
		UpdateSuccess(result interface{}) (int, *Response)
		ModifySuccess(result interface{}) (int, *Response)
		DeleteSuccess(result interface{}) (int, *Response)
		AppendSuccess(module string, result interface{}) (int, *Response)
		RemoveSuccess(module string, result interface{}) (int, *Response)
		ParameterError(data string, err error) *APIException
		GetError(err error) *APIException
		GetModuleError(module string, err error) *APIException
		CreateError(err error) *APIException
		UpdateError(err error) *APIException
		ModifyError(err error) *APIException
		DeleteError(err error) *APIException
		AppendError(module string, err error) *APIException
		RemoveError(module string, err error) *APIException
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

/*
成功訊息
 */

//GetSuccess 取得成功
func (handler *i18nMessageHandler) GetSuccess(result interface{}) (int, *Response) {
	return StatusSuccess(handler.i18n.Get(true, nil), result)
}

//CreateSuccess 建立成功
func (handler *i18nMessageHandler) CreateSuccess(result interface{}) (int, *Response) {
	return StatusSuccess(handler.i18n.Create(true, nil), result)
}

//UpdateSuccess 更新成功
func (handler *i18nMessageHandler) UpdateSuccess(result interface{}) (int, *Response) {
	return StatusSuccess(handler.i18n.Update(true, nil), result)
}

//ModifySuccess 修正成功
func (handler *i18nMessageHandler) ModifySuccess(result interface{}) (int, *Response) {
	return StatusSuccess(handler.i18n.Modify(true, nil), result)
}

//DeleteSuccess 刪除成功
func (handler *i18nMessageHandler) DeleteSuccess(result interface{}) (int, *Response) {
	return StatusSuccess(handler.i18n.Delete(true, nil), result)
}

//AppendSuccess 加入成功
func (handler *i18nMessageHandler) AppendSuccess(module string, result interface{}) (int, *Response) {
	return StatusSuccess(handler.i18n.Append(true, module, nil), result)
}

//RemoveSuccess 移除成功
func (handler *i18nMessageHandler) RemoveSuccess(module string, result interface{}) (int, *Response) {
	return StatusSuccess(handler.i18n.Remove(true, module, nil), result)
}

/*
錯誤訊息
 */
func (handler *i18nMessageHandler) ParameterError(data string, err error) *APIException {
	return ParameterError(
		handler.i18n.ParamIntError(map[string]interface{}{
			"data":  data,
			"error": err.Error(),
		}),
	)
}

func (handler *i18nMessageHandler) GetError(err error) *APIException {
	return UnknownError(handler.i18n.Get(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) GetModuleError(module string, err error) *APIException {
	return UnknownError(handler.i18n.Get(false, map[string]interface{}{
		"data":  module,
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) CreateError(err error) *APIException {
	return UnknownError(handler.i18n.Create(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) UpdateError(err error) *APIException {
	return UnknownError(handler.i18n.Update(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) ModifyError(err error) *APIException {
	return UnknownError(handler.i18n.Modify(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) DeleteError(err error) *APIException {
	return UnknownError(handler.i18n.Delete(false, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) AppendError(module string, err error) *APIException {
	return UnknownError(handler.i18n.Append(false, module, map[string]interface{}{
		"error": err.Error(),
	}))
}

func (handler *i18nMessageHandler) RemoveError(module string, err error) *APIException {
	return UnknownError(handler.i18n.Remove(false, module, map[string]interface{}{
		"error": err.Error(),
	}))
}
