package e

import "strings"

//HandleSQLError : 處理SQL所報的錯誤資訊
func HandleSQLError(err string) error {
	if strings.Contains(err, "1062") {
		return DuplicateError("已有重複資料")
	}
	if strings.Contains(err, "record not found") {
		return NoContentError("找不到資料")
	}
	return UnknownError(err)
}
