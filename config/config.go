package config

import (
	"os"

	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
)

type Server struct {
	Network string `yaml:"network"`
	Addr    string `yaml:"addr"`
	Timeout int    `yaml:"timeout"`
}

type Wepro struct {
	AppId      string `yaml:"appid"`
	Secret     string `yaml:"secret"`
	JwtIssUser string `yaml:"jwtissuser"`
	JwtSign    string `yaml:"jwtsign"`
}

type WxPay struct {
	Mchid        string `yaml:"mchid"`
	ApiKey       string `yaml:"api_key"`
	Apiv3Key     string `yaml:"apiv3_key"`
	CertP12      string `yaml:"cert_p12"`
	CertPem      string `yaml:"cert_pem"`
	KeyPem       string `yaml:"key_pem"`
	SerialNumber string `yaml:"serial_number"`
	NotifyUrl    string `yaml:"notify_url"`
}

type Mysql struct {
	User                 string `yaml:"user"`
	Passwd               string `yaml:"passwd"`
	Net                  string `yaml:"net"`
	Addr                 string `yaml:"addr"`
	DBName               string `yaml:"dbname"`
	AllowNativePasswords bool   `yaml:"allow_native_passwords"`
	Collation            string `yaml:"collation"`
}

type Elastic struct {
	Index    string `yaml:"index"`
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Oss struct {
	StsAccessKeyId     string `yaml:"sts_access_key_id"`
	StsAccessKeySecret string `yaml:"sts_access_key_secret"`
	Endpoint           string `yaml:"endpoint"`
	Bucket             string `yaml:"bucket"`
	Domain             string `yaml:"domain"`
}

type Mns struct {
	Topic           string `yaml:"topic"`
	Queue           string `yaml:"queue"`
	QueueEndpoint   string `yaml:"queue_endpoint"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
}

type Plugin struct {
	Url string `yaml:"url"`
}

type Config struct {
	Http    Server  `yaml:"http"`
	MySQL   Mysql   `yaml:"mysql"`
	Wepro   Wepro   `yaml:"wepro"`
	WxPay   WxPay   `yaml:"wxpay"`
	Elastic Elastic `yaml:"elastic"`
	Oss     Oss     `yaml:"oss"`
	Mns     Mns     `yaml:"mns"`
	Plugin  Plugin  `yaml:"plugin"`
}

var (
	GlobalCfg = new(Config)
)

func Init() {
	client, err := api.NewClient(&api.Config{
		Address: os.Getenv("CONSUL_ADDR"),
		Token:   "p2BE1Atpwsx3rxZdC6k+eXA==",
	})
	if err != nil {
		panic(err)
	}

	kv, _, err := client.KV().Get("shellverse/config.yaml", nil)
	if err != nil {
		panic("consul获取配置失败:" + err.Error())
	}

	if err := yaml.Unmarshal(kv.Value, GlobalCfg); err != nil {
		panic(err)
	}
	GlobalCfg.MySQL.AllowNativePasswords = true
	GlobalCfg.MySQL.Collation = "utf8mb4_general_ci"
}
