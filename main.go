package mappers

/*
NewPostgres - constructor for Postgres data mapper
*/
func NewPostgres(config DBConfig, source string) Postgres {
	return Postgres{
		DBConfig: config,
		Source:   source,
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
