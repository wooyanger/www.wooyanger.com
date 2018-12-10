package models

type Config struct {
	Id			int64		`xorm:"int(20) autoincr pk"`
	Name		string		`xorm:"varchar(255) notnull unique"`
	Value		string		`xorm:"varchar(255) notnull unique"`
}

func (c *Config) GetAllConfig() []Config {
	configs := make([]Config, 0)
	x.Find(&configs)
	return configs
}

func (c *Config) GetSiteName() string {
	config := &Config{Name: "site_name"}
	x.Get(config)
	return config.Value
}

func (c *Config) GetHomeIntroHeader() string {
	config := &Config{Name: "home_intro_header"}
	x.Get(config)
	return config.Value
}

func (c *Config) GetHomeIntroContent() string {
	config := &Config{Name: "home_intro_content"}
	x.Get(config)
	return config.Value
}

func UpdateConfig(c *Config) error {
	_, err := x.Update(c)
	if err != nil {
		return err
	}
	return nil
}