package config

// Profiles 属性配置文件
type Profiles interface {
	GetInclude() string
}

// defaultProfiles 属性配置文件
type defaultProfiles struct {
	Include string `json:"include"`
}

// Include 包含的配置文件
func (p defaultProfiles) GetInclude() string {
	return p.Include
}
