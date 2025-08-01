package config

type BootStrap struct {
	service Service `json:"service"`
	data    Data    `json:"data"`
}

func NewConfig() *BootStrap {
	return &BootStrap{}
}

type Service struct {
	Env     string `json:"env"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Data struct {
	database Database `json:"database"`
}

type Database struct {
	Driver string `json:"driver"`
	Source string `json:"source"`
	Debug  bool   `json:"debug"`
}
