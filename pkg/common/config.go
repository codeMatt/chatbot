package config


type Config struct {
    Port  string
}

func LoadConfig() (c Config, err error) {

	c.Port = "localhost:3000"	
    return
}