package configs

type Config struct {
	Address  string `json:"address"`
	LogLevel string `json:"log_level"`
	IsDev    bool   `json:"is_dev"`
}
