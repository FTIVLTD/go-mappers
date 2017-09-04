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
NewPostgres - constructor for Postgres data mapper
*/
func NewPostgres(config DBConfig, source string) Postgres {
	return Postgres{
		DBConfig: config,
		Source:   source,
	}
}

// NewRedis - redis constructor
func NewRedis(h string, p int, u, pass string) Redis {
	return Redis{
		DBConfig: DBConfig{
			Host:     h,
			Port:     p,
			Password: pass,
			User:     u,
		},
	}
}

/*
NewHTTP - constructor for HTTP(S) data mapper
*/
func NewHTTP(host string, port int, path string) HTTP {
	return HTTP{
		Host: host,
		Port: port,
		Path: path,
	}
}
