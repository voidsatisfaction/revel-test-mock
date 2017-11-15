package configs

type Config struct {
	SeastallionOwnd string
}

func Initialize() *Config {
	c := &Config{}
	c.SeastallionOwnd = "http://157.112.64.162:8080/api/owned_media?category=0&from_date=2017-10-21&indicator=pv&to_date=2017-10-23&type=ranking_category_site"

	return c
}
