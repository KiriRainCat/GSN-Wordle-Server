package config

type Configuration struct {
	Server     Server     `mapstructure:"server" json:"server" yaml:"server"`
	Postgresql Postgresql `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	SMTP       SMTP       `mapstructure:"smtp" json:"smtp" yaml:"smtp"`
}

type Server struct {
	Port          string `mapstructure:"port" json:"port,omitempty" yaml:"port"`
	EncryptSalt   string `mapstructure:"encrypt_salt" json:"encrypt_salt" yaml:"encrypt_salt"`
	RequestAuth   string `mapstructure:"request_auth" json:"request_auth" yaml:"request_auth"`
	JwtEncrypt    string `mapstructure:"jwt_encrypt" json:"jwt_encrypt" yaml:"jwt_encrypt"`
	JwtIssuer     string `mapstructure:"jwt_issuer" json:"jwt_issuer" yaml:"jwt_issuer"`
	FileStorePath string `mapstructure:"file_store_path" json:"file_store_path" yaml:"file_store_path"`
}

type Postgresql struct {
	DevHost     string `mapstructure:"dev_host" json:"dev_host,omitempty" yaml:"dev_host"`
	DevDb       string `mapstructure:"dev_db" json:"dev_db,omitempty" yaml:"dev_db"`
	DevUser     string `mapstructure:"dev_user" json:"dev_user,omitempty" yaml:"dev_user"`
	DevPassword string `mapstructure:"dev_password" json:"dev_password,omitempty" yaml:"dev_password"`
	Host        string `mapstructure:"host" json:"host,omitempty" yaml:"host"`
	Port        int    `mapstructure:"port" json:"port,omitempty" yaml:"port"`
	Db          string `mapstructure:"db" json:"db,omitempty" yaml:"db"`
	User        string `mapstructure:"user" json:"user,omitempty" yaml:"user"`
	Password    string `mapstructure:"password" json:"password,omitempty" yaml:"password"`
}

type Redis struct {
	DevHost     string `mapstructure:"dev_host" json:"dev_host,omitempty" yaml:"dev_host"`
	DevPort     int    `mapstructure:"dev_port" json:"dev_port,omitempty" yaml:"dev_port"`
	DevUser     string `mapstructure:"dev_user" json:"dev_user,omitempty" yaml:"dev_user"`
	DevPassword string `mapstructure:"dev_password" json:"dev_password,omitempty" yaml:"dev_password"`
	DevDb       int    `mapstructure:"dev_db" json:"dev_db,omitempty" yaml:"dev_db"`
	Host        string `mapstructure:"host" json:"host,omitempty" yaml:"host"`
	Port        int    `mapstructure:"port" json:"port,omitempty" yaml:"port"`
	Password    string `mapstructure:"password" json:"password,omitempty" yaml:"password"`
	Db          int    `mapstructure:"db" json:"db,omitempty" yaml:"db"`
}

type SMTP struct {
	Host string `mapstructure:"host" json:"host,omitempty" yaml:"host"`
	Port int    `mapstructure:"port" json:"port,omitempty" yaml:"port"`
	Key  string `mapstructure:"key" json:"key,omitempty" yaml:"key"`
	Mail string `mapstructure:"mail" json:"mail,omitempty" yaml:"mail"`
}
