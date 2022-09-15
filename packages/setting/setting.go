package setting

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string
	Env       string

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Sslmode  string
}

var DatabaseSetting = &Database{}

type Logger struct {
	Output     string
	OutputFile string
	Level      int
	Format     string
}

var GoogleCloudSetting = &GoogleCloud{}

type GoogleCloud struct {
	BucketName     string
	CredentialFile string
}

var LoggerSetting = &Logger{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local"
	}
	cfg, err = ini.Load(".ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("logger", LoggerSetting)
	mapTo("googleCloud", GoogleCloudSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
