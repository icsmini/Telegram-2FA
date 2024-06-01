package apiserver

type Config struct {
	bindAddr string 'toml: "bind_addr"'

}

func NewConfig() *Config {
	return &Config {
		bindAddr: ":8080",
	}
}