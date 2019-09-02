package config

const (
	AppMode = "release"
	AppPort = ":9999"
	AppName = "gin-samples"

	AppReadTimeout  = 120
	AppWriteTimeout = 120

	AppAccessLogName = "log/" + AppName + "-access.log"
	AppErrorLogName  = "log/" + AppName + "-error.log"
)
