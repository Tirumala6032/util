package config

type Config struct {
	Server           	Server           	`yaml:"server" json:"server"`
	LogLevel         	string           	`yaml:"logLevel" json:"logLevel"`
	Database         	Database         	`yaml:"database" json:"database"`
	AuthService      	AuthService      	`yaml:"authService" json:"authService"`
	WebSocketService 	WebSocketService 	`yaml:"webSocketService" json:"webSocketService"`
	UserService      	UserService      	`yaml:"userService" json:"userService"`
	LeaderBoardService     	LeaderBoardService      `yaml:"leaderBoardService" json:"leaderBoardService"`
}

type Server struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
}

type Database struct {
	DBHost     string `yaml:"dbHost" json:"dbHost"`
	DBPort     int    `yaml:"dbPort" json:"dbPort"`
	DBName     string `yaml:"dbName" json:"dbName"`
	DBDriver   string `yaml:"dbDriver" json:"dbDriver"`
	DBUsername string `yaml:"dbUsername" json:"dbUsername"`
	DBPassword string `yaml:"dbPassword" json:"dbPassword"`
}

type AuthService struct {
	ServerHost string `yaml:"serverHost" json:"serverHost"`
	ServerPort int    `yaml:"serverPort" json:"serverPort"`
}

type WebSocketService struct {
	ServerHost string `yaml:"serverHost" json:"serverHost"`
	ServerPort int    `yaml:"serverPort" json:"serverPort"`
}

type UserService struct {
	ServerHost string `yaml:"serverHost" json:"serverHost"`
	ServerPort int    `yaml:"serverPort" json:"serverPort"`
}

type LeaderBoardService struct {
	ServerHost string `yaml:"serverHost" json:"serverHost"`
	ServerPort int    `yaml:"serverPort" json:"serverPort"`
}
