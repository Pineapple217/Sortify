package config

type Server struct {
	Port   int    `yaml:"port"`
	Bind   string `yaml:"bind"`
	Debug  bool   `yaml:"debug"`
	Secret string `yaml:"secret"`
}

func (s *Server) SetDefault() {
	s.Port = 3000
	s.Bind = "127.0.0.1"
	s.Debug = false
}
