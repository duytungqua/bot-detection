package config

type Config struct {
	API struct {
		Endpoint string
	}
}
var Cfg Config


func Load() Config{
	Cfg.API.Endpoint = "https://api.example.com"
	return Cfg	

}