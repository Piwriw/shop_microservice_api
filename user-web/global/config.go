package global

type AppConfig struct {
	System     System     `json:"system" mapstructure:"system"`
	GrpcServer GrpcServer `json:"grpcServer" mapstructure:"grpcServer"`
	Jwt        Jwt        `json:"jwt" mapstructure:"jwt"`
	Consul     Consul     `json:"consul" mapstructure:"consul"`
	Nacos      Nacos      `json:"nacos" mapstructure:"nacos"`
}
type Nacos struct {
	Host      string `json:"host" mapstructure:"host"`
	Port      uint64 `json:"port" mapstructure:"port"`
	NameSpace string `json:"nameSpace" mapstructure:"nameSpace"`
	User      string `json:"user" mapstructure:"user"`
	Password  string `json:"password" mapstructure:"password"`
	DataID    string `json:"dataID" mapstructure:"dataId"`
	Group     string `json:"group" mapstructure:"group"`
}
type Consul struct {
	IP   string `json:"ip" mapstructure:"ip"`
	Port int    `json:"port" mapstructure:"port"`
}
type Jwt struct {
	SigningKey string `mapstructure:"key"`
	ExpiresAt  int64  `mapstructure:"expiresAt"`
}

type System struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
	Name string `json:"name" mapstructure:"name"`
}
type GrpcServer struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}
