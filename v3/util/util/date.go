package util

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Operator string

const (
	Less        Operator = "lt"
	LessThan    Operator = "lta"
	Equal       Operator = "eq"
	Greater     Operator = "gt"
	GreaterThan Operator = "gta"
)

func (o Operator) Get() string {
	switch o {
	case Less:
		return "<"
	case LessThan:
		return "<="
	case Equal:
		return "="
	case Greater:
		return ">"
	case GreaterThan:
		return ">-"
	default:
		return "="
	}
}

type DateTimeCondition struct {
	Operator string `json:"operator"`
	Number   string `json:"number"`
}

type QueryDateTimeCondition struct {
	Column string             `json:"column"`
	Year   *DateTimeCondition `json:"year"`
	Month  *DateTimeCondition `json:"month"`
	Date   *DateTimeCondition `json:"date"`
}

func NewCondition(str string) *DateTimeCondition {
	temp := strings.Split(str, "__")
	if len(temp) > 1 {
		return &DateTimeCondition{
			Operator: Operator(temp[0]).Get(),
			Number:   temp[1],
		}
	}
	return &DateTimeCondition{
		Operator: Equal.Get(),
		Number:   "0",
	}
}

func FilterDateTimeRange(condition *QueryDateTimeCondition) func(in *gorm.DB) *gorm.DB {
	return func(in *gorm.DB) *gorm.DB {
		if c := condition.Year; c != nil {
			str := fmt.Sprintf("YEAR(CURDATE()) - YEAR(%s) %s %s", condition.Column, c.Operator, c.Number)
			in = in.Where(str)
		}
		if c := condition.Month; c != nil {
			str := fmt.Sprintf("MONTH(CURDATE()) - MONTH(%s) %s %s", condition.Column, c.Operator, c.Number)
			in = in.Where(str)
		}
		if c := condition.Date; c != nil {
			str := fmt.Sprintf("DAY(CURDATE()) - DAY(%s) %s %s", condition.Column, c.Operator, c.Number)
			in = in.Where(str)
		}
		return in
	}
}
