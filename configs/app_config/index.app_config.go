package app_config

import "os"

var PORT = ":8000"
var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")

	if portEnv != "" {
		PORT = portEnv
	}

	staticRouteEnv := os.Getenv("APP_STATIC_ROUTE")

	if staticRouteEnv != "" {
		STATIC_ROUTE = staticRouteEnv
	}

	staticDirEnv := os.Getenv("APP_STATIC_DIR")

	if staticDirEnv != "" {
		STATIC_DIR = staticDirEnv
	}
}
