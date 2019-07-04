package config

// PostgresConfig Postgre 配置 接口
type PostgresConfig interface {
	GetEnabled() bool
	GetHost() string
	GetPort() int
	GetDBname() string
	GetUser() string
	GetPassword() string
}

// defaultPostgresConfig postgres 配置
type defaultPostgresConfig struct {
	Enable   bool   `json:"enabled"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBname   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Enabled 激活
func (m defaultPostgresConfig) GetEnabled() bool {
	return m.Enable
}

// GetHost host 主机名
func (m defaultPostgresConfig) GetHost() string {
	return m.Host
}

// GetPort 连接端口
func (m defaultPostgresConfig) GetPort() int {
	return m.Port
}

// GetDBname 数据库名称
func (m defaultPostgresConfig) GetDBname() string {
	return m.DBname
}

// GetUser 用户名
func (m defaultPostgresConfig) GetUser() string {
	return m.User
}

// GetDBname 密码
func (m defaultPostgresConfig) GetPassword() string {
	return m.Password
}
