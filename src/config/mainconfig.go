package config

type Config struct {
	Server   ServerConf
	Database map[string]*struct {
		Master string
		Slave  []string
	}
}

type ServerConf struct {
	Name string
}
