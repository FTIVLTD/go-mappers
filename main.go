package mappers

/*
DBConfig - Postgres config
*/
type DBConfig struct {
	User,
	Password,
	Host string
	Port int
	Database,
	SSLmode string
}

/*
HTTPConfig - HTTP(s) config
*/
type HTTPConfig struct {
	User,
	Password,
	Host string
	Port int
}

/*
NewPostgres - constructor for Postgres data mapper
*/
func NewPostgres(config DBConfig) Postgres {
	return Postgres{
		DBConfig: config,
	}
}

// NewRedis - redis constructor
func NewRedis(conf DBConfig) Redis {
	return Redis{
		DBConfig: conf,
	}
}

/*
NewHTTP - constructor for HTTP(S) data mapper
*/
func NewHTTP(conf HTTPConfig) HTTP {
	return HTTP{
		Host: conf.Host,
		Port: conf.Port,
	}
}

/*
NewMongo - constructor for Mongodb data mapper
*/
func NewMongo(config DBConfig, c string, limit int) Mongo {
	return Mongo{
		DBConfig:   config,
		Collection: c,
		Limit:      limit,
	}
}
