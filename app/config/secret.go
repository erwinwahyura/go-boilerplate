package config

import (
	"github.com/spf13/viper"
)

var DB_NAME string
var DB_DRIVER int
var DB_MASTER string
var DB_SLAVE string

// log
var DB_LOG_NAME string
var DB_LOG_DRIVER string
var DB_LOG_MASTER string
var DB_LOG_SLAVE int

// env
var ENV string

// log
var LOG_LEVEL string

// host
var HOST_ADDRESS string
var HOST_IDLE_TIMEOUT string
var HOST_PORT string
var HOST_READ_TIMEOUT int
var HOST_WRITE_TIMEOUT string
var LEVEL string

var MYVALUE_BASE_URL string

var MYVALUE_CLIENT_ID string
var MYVALUE_CLIENT_SECRET int
var MYVALUE_EXTERNAL_URL string
var MYVALUE_REDIRECT_URI string
var SERVICE_NAME string

var SLACK_BOT_NAME string
var SLACK_CHANNEL string
var SLACK_COLOR string
var SLACK_ICON string
var SLACK_URL string
var ISSUER int

var SECRETKEY string
var APP_VERSION string

var REDIS_HOST string
var REDIS_PORT string

var IMAGE_BASE_URL string

var MEILI_HOST string
var MEILI_PORT string
var MEILI_API_KEY string

// Reload reload secret from system's ENV
func Reload() {
	// postgres
	DB_NAME = viper.GetString("DB_NAME")
	DB_DRIVER = viper.GetInt("DB_DRIVER")
	DB_MASTER = viper.GetString("DB_MASTER")
	DB_SLAVE = viper.GetString("DB_SLAVE")

	// log mongodb
	DB_LOG_NAME = viper.GetString("DB_LOG_NAME")
	DB_LOG_DRIVER = viper.GetString("DB_LOG_DRIVER")
	DB_LOG_MASTER = viper.GetString("DB_LOG_MASTER")
	DB_LOG_SLAVE = viper.GetInt("DB_LOG_SLAVE")

	// others
	ENV = viper.GetString("ENV")
	LOG_LEVEL = viper.GetString("LOG_LEVEL")
	HOST_ADDRESS = viper.GetString("HOST_ADDRESS")
	HOST_IDLE_TIMEOUT = viper.GetString("HOST_IDLE_TIMEOUT")
	HOST_PORT = viper.GetString("HOST_PORT")
	HOST_READ_TIMEOUT = viper.GetInt("HOST_READ_TIMEOUT")
	HOST_WRITE_TIMEOUT = viper.GetString("HOST_WRITE_TIMEOUT")
	LEVEL = viper.GetString("LEVEL")
	SECRETKEY = viper.GetString("SECRETKEY")
	SERVICE_NAME = viper.GetString("SERVICE_NAME")
	ISSUER = viper.GetInt("ISSUER")

	// myvalue
	MYVALUE_BASE_URL = viper.GetString("MYVALUE_BASE_URL")
	MYVALUE_CLIENT_ID = viper.GetString("MYVALUE_CLIENT_ID")
	MYVALUE_CLIENT_SECRET = viper.GetInt("MYVALUE_CLIENT_SECRET")
	MYVALUE_EXTERNAL_URL = viper.GetString("MYVALUE_EXTERNAL_URL")
	MYVALUE_REDIRECT_URI = viper.GetString("MYVALUE_REDIRECT_URI")

	// slacks
	SLACK_BOT_NAME = viper.GetString("SLACK_BOT_NAME")
	SLACK_CHANNEL = viper.GetString("SLACK_CHANNEL")
	SLACK_COLOR = viper.GetString("SLACK_COLOR")
	SLACK_ICON = viper.GetString("SLACK_ICON")
	SLACK_URL = viper.GetString("SLACK_URL")

	// redis
	REDIS_HOST = viper.GetString("REDIS_HOST")
	REDIS_PORT = viper.GetString("REDIS_PORT")

	IMAGE_BASE_URL = viper.GetString("IMAGE_BASE_URL")

	// meilisearch
	MEILI_HOST = viper.GetString("MEILI_HOST")
	MEILI_PORT = viper.GetString("MEILI_PORT")
	MEILI_API_KEY = viper.GetString("MEILI_API_KEY")

}

func ViperBind() {
	// database
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_DRIVER")
	viper.BindEnv("DB_MASTER")
	viper.BindEnv("DB_SLAVE")
	viper.BindEnv("DB_LOG_NAME")
	viper.BindEnv("DB_LOG_DRIVER")
	viper.BindEnv("DB_LOG_MASTER")
	viper.BindEnv("DB_LOG_SLAVE")

	// myvalue
	viper.BindEnv("MYVALUE_BASE_URL")
	viper.BindEnv("MYVALUE_CLIENT_ID")
	viper.BindEnv("MYVALUE_CLIENT_SECRET")
	viper.BindEnv("MYVALUE_EXTERNAL_URL")
	viper.BindEnv("MYVALUE_REDIRECT_URI")

	// others
	viper.BindEnv("SERVICE_NAME")
	viper.BindEnv("LEVEL")
	viper.BindEnv("ENV")
	viper.BindEnv("LOG_LEVEL")
	viper.BindEnv("HOST_ADDRESS")
	viper.BindEnv("HOST_IDLE_TIMEOUT")
	viper.BindEnv("HOST_PORT")
	viper.BindEnv("HOST_READ_TIMEOUT")
	viper.BindEnv("HOST_WRITE_TIMEOUT")
	viper.BindEnv("SECRETKEY")
	viper.BindEnv("ISSUER")

	// slack
	viper.BindEnv("SLACK_BOT_NAME")
	viper.BindEnv("SLACK_CHANNEL")
	viper.BindEnv("SLACK_COLOR")
	viper.BindEnv("SLACK_ICON")
	viper.BindEnv("SLACK_URL")

	// redis
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_PORT")

	// image
	viper.BindEnv("IMAGE_BASE_URL")

	// meilisearch
	viper.BindEnv("MEILI_HOST")
	viper.BindEnv("MEILI_PORT")
	viper.BindEnv("MEILI_API_KEY")

}
