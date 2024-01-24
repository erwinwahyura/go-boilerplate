package constant

import "os"

var (
	// Header
	ChannelID = "X-Channel-Id"
	ApiKey    = "api-key"
	RequestID = "request-id"

	// Environment
	ENV_PROD    = "prod"
	ENV_STAGING = "staging"
	ENV_DEV     = "dev"

	// Live
	IS_PRODUCTION = os.Getenv("ENVIRONMENT") == ENV_PROD
)
