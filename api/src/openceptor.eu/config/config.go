package config

type Config struct {
	Database struct {
		Host			string `yaml:"host" env:"DB_HOST" env-default:"127.0.0.1"`
		Port			int    `yaml:"port" env:"DB_PORT" env-default:"5432"`
		Username		string `yaml:"username" env:"DB_USERNAME"`
		Password		string `yaml:"password" env:"DB_PASSWORD"`
		Database		string `yaml:"database" env:"DB_DATABASE" env-default:"openceptor"`
		UseSsl   		bool   `yaml:"ssl" env:"DB_SSL" env-default:"false"`
		MaxOpenConns	int    `yaml:"max_open_conns" env:"DB_MAX_OPEN_CONNS" env-default:"0"`
		MaxIdleConns	int    `yaml:"max_idle_conns" env:"DB_MAX_IDLE_CONNS" env-default:"0"`
	} `yaml:"database"`

	Queue struct {
		Host     string `yaml:"host" env:"QUEUE_HOST" env-default:"127.0.0.1"`
		Port     int    `yaml:"port" env:"QUEUE_PORT" env-default:"5672"`
		Username string `yaml:"username" env:"QUEUE_USERNAME" env-default:"guest"`
		Password string `yaml:"password" env:"QUEUE_PASSWORD" env-default:"guest"`
		Vhost    string `yaml:"vhost" env:"QUEUE_VHOST" env-default:"/openceptor"`
		UseSsl   bool   `yaml:"ssl" env:"QUEUE_SSL" env-default:"false"`
	} `yaml:"queue"`

	Server struct {
		Host          string `yaml:"host" env:"SERVER_HOST" env-default:"127.0.0.1"`
		Port          int    `yaml:"port" env:"SERVER_PORT" env-default:"6660"`
		UploadMaxSize int    `yaml:"upload_max_size" env:"SERVER_UPLOAD_MAX_SIZE" env-default:"1024"`
		LogLevel      string `yaml:"log_level" env:"LOG_LEVEL" env-default:"error"`
	} `yaml:"server"`
}
