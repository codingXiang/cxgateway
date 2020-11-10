package model

type UserMeta struct {
	ID     int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT;comment:'id'"`
	UserID string `json:"userId,omitempty" gorm:"not_null;unique_index:idx1;comment:'cloud id'"`
	Key    string `json:"key" gorm:"not_null;unique_index:idx1;Column:key;comment:'鍵'"`
	Value  string `json:"value" gorm:"Column:value;coment:'值'"`
}
