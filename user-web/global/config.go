package global

type AppConfig struct {
	System         System     `json:"system" yaml:"system" mapstructure:"system"`
	UserGrpcServer GrpcServer `json:"userGrpcServer" yaml:"userGrpcServer" mapstructure:"userGrpcServer"`
	Jwt            Jwt        `json:"jwt" yaml:"jwt" mapstructure:"jwt"`
	Consul         Consul     `json:"consul" yaml:"consul" mapstructure:"consul"`
}
type Nacos struct {
	Host      string `json:"host" mapstructure:"host"`
	Port      uint64 `json:"port" mapstructure:"port"`
	NameSpace string `json:"nameSpace" mapstructure:"nameSpace"`
	User      string `json:"user" mapstructure:"user"`
	Password  string `json:"password" mapstructure:"password"`
	DataID    string `json:"dataId" mapstructure:"dataId"`
	Group     string `json:"group" mapstructure:"group"`
}
type Consul struct {
	IP   string `json:"ip" mapstructure:"ip"`
	Port int    `json:"port" mapstructure:"port"`
}
type Jwt struct {
	SigningKey string `yaml:"key" mapstructure:"key"`
	ExpiresAt  int64  `yaml:"expiresAt" mapstructure:"expiresAt"`
}

type System struct {
	Host string   `json:"host" yaml:"host" mapstructure:"host"`
	Port int      `json:"port" yaml:"port" mapstructure:"port"`
	Name string   `json:"name" yaml:"name" mapstructure:"name"`
	Tags []string `json:"tags" yaml:"tags" mapstructure:"tags"`
}
type GrpcServer struct {
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	Port int    `json:"port" yaml:"port" mapstructure:"port"`
	Name string `json:"name" yaml:"name" mapstructure:"name"`
}
