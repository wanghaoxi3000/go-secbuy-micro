package config

// PostgreConfig Postgre 配置 接口
type PostgreConfig interface {
	GetEnabled() bool
	GetHost() string
	GetPort() int
	GetDBname() string
	GetUser() string
	GetPassword() string
}

// postgre 配置
type defaultPostgreConfig struct {
	Enable   bool   `json:"enabled"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBname   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Enabled 激活
func (m defaultPostgreConfig) GetEnabled() bool {
	return m.Enable
}

// GetHost host 主机名
func (m defaultPostgreConfig) GetHost() string {
	return m.Host
}

// GetPort 连接端口
func (m defaultPostgreConfig) GetPort() int {
	return m.Port
}

// GetDBname 数据库名称
func (m defaultPostgreConfig) GetDBname() string {
	return m.DBname
}

// GetUser 用户名
func (m defaultPostgreConfig) GetUser() string {
	return m.User
}

// GetDBname 密码
func (m defaultPostgreConfig) GetPassword() string {
	return m.Password
}
