package config

var Guards = map[string]string{
	"/api": "api-user",
	"/api/panel": "admin-user",
}