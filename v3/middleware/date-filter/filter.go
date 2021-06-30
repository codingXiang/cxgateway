package date_filter

import (
	"github.com/codingXiang/cxgateway/v3/middleware"
	"github.com/codingXiang/cxgateway/v3/util/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// example: datetime_filter_column=birthday&month_filter==0&date_filter==0

const (
	FilterKey    = "datetime_column_filter"
	YearRangeKey = "year_filter"
	MonthKey     = "month_filter"
	DateKey      = "date_filter"
	ColumnKey    = "datetime_filter_column"
)

type Handler struct{}

func New() middleware.Object {
	return new(Handler)
}

func (*Handler) SetConfig(config *viper.Viper) {

}

func (c *Handler) GetConfig() *viper.Viper {
	return viper.New()
}

func (c *Handler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			res *util.QueryDateTimeCondition
		)
		if in, ok := c.GetQuery(ColumnKey); ok {
			res = new(util.QueryDateTimeCondition)
			res.Column = in
		}
		if in, ok := c.GetQuery(YearRangeKey); ok && res != nil {
			res.Year = util.NewCondition(in)
		}
		if in, ok := c.GetQuery(MonthKey); ok && res != nil {
			res.Month = util.NewCondition(in)
		}
		if in, ok := c.GetQuery(DateKey); ok && res != nil {
			res.Date = util.NewCondition(in)
		}

		if res != nil {
			c.Set(FilterKey, res)
		}
		c.Next()
	}
}
