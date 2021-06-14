package user

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type Response struct {
	User *User `json:"user"`
}

type User struct {
	ID     int64    `json:"id"`
	Roles  []*Role  `json:"roles"`
	Groups []*Group `json:"groups"`
}

func GetUser(c *gin.Context) (*User, error) {
	in, exist := c.Get(UserInfo)
	if !exist {
		return nil, errors.New("user info not exist in context")
	}
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	var user = new(User)
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type Role struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Children []*Role `json:"children"`
}

type Group struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Children []*Group `json:"children"`
}

type Pivot struct {
	ModelID      int64  `json:"model_id,omitempty"`
	RoleID       int64  `json:"role_id,omitempty"`
	ModelType    string `json:"model_type,omitempty"`
	PermissionID int64  `json:"permission_id,omitempty"`
}

type Permission struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	GuardName string    `json:"guardName"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Pivot     *Pivot    `json:"pivot"`
}
