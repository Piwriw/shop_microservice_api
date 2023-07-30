package global

type AppConfig struct {
	System     System     `json:"system" yaml:"system" mapstructure:"system"`
	GrpcServer GrpcServer `json:"grpcServer" yaml:"grpcServer" mapstructure:"grpcServer"`
	Jwt        Jwt        `json:"jwt" yaml:"jwt" mapstructure:"jwt"`
	Consul     Consul     `json:"consul" yaml:"consul" mapstructure:"consul"`
	Nacos      Nacos      `json:"nacos" yaml:"nacos" mapstructure:"nacos"`
}
type Nacos struct {
	Host      string `json:"host" yaml:"host" mapstructure:"host"`
	Port      uint64 `json:"port" yaml:"port" mapstructure:"port"`
	NameSpace string `json:"nameSpace" yaml:"nameSpace" mapstructure:"nameSpace"`
	User      string `json:"user" yaml:"user" mapstructure:"user"`
	Password  string `json:"password" yaml:"password" mapstructure:"password"`
	DataID    string `json:"dataId" yaml:"dataId" mapstructure:"dataId"`
	Group     string `json:"group" yaml:"group" mapstructure:"group"`
}
type Consul struct {
	IP   string `json:"ip" yaml:"ip" mapstructure:"ip"`
	Port int    `json:"port" yaml:"port" mapstructure:"port"`
}
type Jwt struct {
	SigningKey string `yaml:"signingKey" mapstructure:"key"`
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
	Name string `json:"name" yaml:"name" mapstructure:"port"`
}
