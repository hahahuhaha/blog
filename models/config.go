package models

type Config struct {
	Id    int    `column:"Id"`
	Name  string `column:"Name"`
	Value string `column:"Value"`
}

func (m *Config) TableName() string {
	return TableName("config")
}
