package config

type config struct {
	*Env
}

var conf *config

func init() {
	conf = &config{}
}

func Config() *config {
	return conf
}

func SetEnv(env *Env) {
	conf.Env = env
}
