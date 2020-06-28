package config

type Env struct {
	Mysql struct {
		Host     string `default:"127.0.0.1"`
		User     string
		Password string
		Port     int    `default:"3308"`
		Db       string `default:"at_project"`
	}
}
