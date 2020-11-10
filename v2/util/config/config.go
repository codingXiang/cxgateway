package config

import "strings"

const (
	Application = "application"
)

func GetConfigPath(key string, path ...string) string {
	return key + "." + strings.Join(path, ".")
}
