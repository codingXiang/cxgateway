package response

import "github.com/gin-gonic/gin"

type Handler struct {
	*gin.Context
}

func New(ctx *gin.Context) *Handler {
	return &Handler{ctx}
}

func (h *Handler) SetStatusCode(code int) *Handler {
	h.Set(StatusCodeKey, code)
	return h
}

func (h *Handler) SetMessage(msg string) *Handler {
	h.Set(MessageKey, msg)
	return h
}

func (h *Handler) SetData(data interface{}) *Handler {
	h.Set(DataKey, data)
	return h
}

func (h *Handler) AddFails(msg ...string) *Handler {
	fails := h.GetStringSlice(FailKey)
	if fails == nil {
		h.Set(FailKey, msg)
	} else {
		fails = append(fails, msg...)
		h.Set(FailKey, fails)
	}
	return h
}

func (h *Handler) SetPageInfo(info *PageInfo) *Handler {
	h.Set(PageKey, info.Setup(h.Context))
	return h
}

func SetResponseWithStatus(c *gin.Context, statusCode int, message string, data interface{}, fails []string, pageInfo *PageInfo) {
	c.Set(StatusCodeKey, statusCode)
	c.Set(MessageKey, message)
	c.Set(DataKey, data)
	if pageInfo != nil {
		c.Set(PageKey, pageInfo.Setup(c))
	}
	if fails != nil {
		c.Set(FailKey, fails)
	}
}

//HttpStatusHandler 依照 http status code 判斷要回傳什麼
type HttpStatusHandler func(c *gin.Context)

//WrapperStatus 在 register routing 時加入錯誤 handler
func WrapperStatus(handler HttpStatusHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(c)
		c.JSON(c.GetInt(StatusCodeKey), newResponse(c))
		return
	}
}
