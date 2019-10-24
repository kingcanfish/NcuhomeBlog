package conf

type ProjectConfig struct {
	DatabaseURI string
}

const debug  = true

var config *ProjectConfig

func InitConfig(debug bool) *ProjectConfig {
	if debug {
		return & ProjectConfig{DatabaseURI:"root:123456@(127.0.0.1)/ncuhome_blog?charset=utf8"}

	}
	return &ProjectConfig{DatabaseURI:"root:123456@(127.0.0.1)/ncuhome_blog?charset=utf8"}

}

func init() {
	config = InitConfig(debug)
}

func GetConfig() *ProjectConfig  {
	return config

}