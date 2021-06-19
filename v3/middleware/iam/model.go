package iam

import (
	"encoding/json"
	"errors"
	"github.com/codingXiang/cxgateway/v3/util/response"
	"github.com/gin-gonic/gin"
	"time"
)

type Object struct {
	Type    string `json:"type"`
	Subject string `json:"subject"`
	Object  string `json:"object"`
	Action  string `json:"action"`
}

type Response struct {
	User *User `json:"user"`
}

func Resp2User(in []byte) (resp response.Response, err error) {
	err = json.Unmarshal(in, &resp)
	return
}

type User struct {
	ID              int64         `json:"id"`
	Name            string        `json:"name"`
	RogAccount      string        `json:"rogaccount"`
	Email           string        `json:"email"`
	EmailVerifiedAt time.Time     `json:"email_verified_at"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Roles           []*Role       `json:"roles"`
	Permissions     []*Permission `json:"permissions"`
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
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	GuardName string    `json:"guardName"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Pivot     *Pivot    `json:"pivot"`
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
