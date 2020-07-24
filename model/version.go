package model

type Version struct {
	ServerVersion   string `json:"serverVersion,omitempty"`
	DatabaseVersion string `json:"databaseVersion,omitempty"`
	RedisVersion    string `json:"redisVersion,omitempty"`
}
