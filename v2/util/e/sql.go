package e

import "strings"

//HandleSQLError : 處理SQL所報的錯誤資訊
func HandleSQLError(err string) error {
	if strings.Contains(err, "1062") {
		return StatusConflict("已有重複資料")
	}
	return UnknownError(err)
}
