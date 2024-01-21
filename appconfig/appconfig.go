package appconfig

import "time"

type AppConfig struct {
	Server struct {
		Url                string
		Port               string
		GracePeriodSeconds uint32
		SwaggerEnable      bool
	}
	JWT struct {
		Secret       string
		ClaimExpired uint32
	}
	Database struct {
		DSN string
	}
}

func (cfg *AppConfig) GracePeriodSeconds() time.Duration {
	return time.Duration(cfg.Server.GracePeriodSeconds) * time.Second
}
