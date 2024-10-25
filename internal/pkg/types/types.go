package types

type HTTPSever struct {
	Addr string `yaml:"addr" env:"ADDR" env-default:":8080" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-default:"/storage/storage.db" env-required:"true"`
	HTTPSever   `yaml:"http_server"`
}