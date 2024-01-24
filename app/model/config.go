package model

type (
	// Config is application configuration
	Config struct {
		Env        string   `mapstructure:"ENV"`
		Level      string   `mapstructure:"LOG_LEVEL"`
		Host       Host     `mapstructure:",squash"`
		Database   Database `mapstructure:",squash"`
		Slack      Slack    `mapstructure:",squash"`
		SecretKey  string   `mapstructure:"SECRETKEY"`
		Issuer     string   `mapstructure:"ISSUER"`
		MyValue    MyValue  `mapstructure:",squash"`
		AppVersion string   `mapstructure:"APP_VERSION"`

		PromoService PromoService `mapstructure:",squash"`
		Redis        Redis        `mapstructure:",squash"`
		Meilisearch  Meilisearch  `mapstructure:",squash"`
		Image        Image        `mapstructure:",squash"`
	}

	// Host server config
	Host struct {
		Address      string `mapstructure:"HOST_ADDRESS"`
		Port         string `mapstructure:"HOST_PORT"`
		WriteTimeout int    `mapstructure:"HOST_WRITE_TIMEOUT" default:"15"`
		ReadTimeout  int    `mapstructure:"HOST_READ_TIMEOUT" default:"15"`
		IdleTimeout  int    `mapstructure:"HOST_IDLE_TIMEOUT" default:"60"`
	}

	// Database all database
	Database struct {
		LogDB DatasourceLog `mapstructure:",squash"`
		DB    Datasource    `mapstructure:",squash"`
	}

	// Datasource datasource detail
	Datasource struct {
		Master string `mapstructure:"DB_MASTER"`
		Slave  string `mapstructure:"DB_SLAVE"`
		Driver string `mapstructure:"DB_DRIVER"`
		DBName string `mapstructure:"DB_DBNAME"`
	}
	// Datasource datasource log detail
	DatasourceLog struct {
		Master string `mapstructure:"DB_LOG_MASTER"`
		Slave  string `mapstructure:"DB_LOG_SLAVE"`
		Driver string `mapstructure:"DB_LOG_DRIVER"`
		DBName string `mapstructure:"DB_LOG_DBNAME"`
	}

	// Slack slack notif
	Slack struct {
		URL     string `mapstructure:"SLACK_URL"`
		Channel string `mapstructure:"SLACK_CHANNEL"`
		BotName string `mapstructure:"SLACK_BOT_NAME"`
		Icon    string `mapstructure:"SLACK_ICON"`
	}

	// MyValue
	MyValue struct {
		ClientID     string `mapstructure:"MYVALUE_CLIENT_ID"`
		ClientSecret string `mapstructure:"MYVALUE_CLIENT_SECRET"`
		RedirectURI  string `mapstructure:"MYVALUE_REDIRECT_URI"`
		BaseURL      string `mapstructure:"MYVALUE_BASE_URL"`
		ExternalURL  string `mapstructure:"MYVALUE_EXTERNAL_URL"`
	}

	// Promo Service
	PromoService struct {
		BaseURL string `mapstructure:"PROMOSERVICE_BASE_URL"`
		Client  string `mapstructure:"PROMOSERVICE_CLIENT"`
	}

	// Redis
	Redis struct {
		Host string `mapstructure:"REDIS_HOST"`
		Port string `mapstructure:"REDIS_PORT"`
	}
	// Meilisearch
	Meilisearch struct {
		Host   string `mapstructure:"MEILI_HOST"`
		Port   string `mapstructure:"MEILI_PORT"`
		ApiKey string `mapstructure:"MEILI_API_KEY"`
	}

	Image struct {
		BaseURL string `mapstructure:"IMAGE_BASE_URL"`
	}
)
