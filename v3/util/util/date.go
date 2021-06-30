package util

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Operator string

const (
	Less        Operator = "<"
	LessThan    Operator = "<="
	Equal       Operator = "="
	Greater     Operator = ">"
	GreaterThan Operator = ">="
)

type DateTimeCondition struct {
	Operator Operator `json:"operator"`
	Number   string   `json:"number"`
}

type QueryDateTimeCondition struct {
	Column string
	Year   *DateTimeCondition
	Month  *DateTimeCondition
	Date   *DateTimeCondition
}

func NewCondition(str string) *DateTimeCondition {
	temp := strings.Split(str, "_")
	if len(temp) > 1 {
		return &DateTimeCondition{
			Operator: Operator(temp[0]),
			Number:   temp[1],
		}
	}
	return &DateTimeCondition{
		Operator: Equal,
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
