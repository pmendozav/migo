package bootstrap

import "strings"

var (
	APP_NAME = `Mio!`
	APP_VERSION = "0.0.1"
	
	DEV_ENV = "dev"
	STAGING_ENV = "stating"
	PROD_ENV = "prod"
	
	REPLACER *strings.Replacer = strings.NewReplacer(".", "_")
	
	APP_CONFIG_PREFIX = `APP`
	APP_CONFIG_NAME = `application`
	DB_CONFIG_PREFIX = `DB`
	DB_CONFIG_NAME = `database`
	
	CONFIG_PATH = `config`
	CONFIG_FILE_TYPE = `yaml`
)