package config

import (
	"github.com/pushgate/core/pkg/encryption"
	"github.com/pushgate/core/pkg/jwt"
	"github.com/pushgate/core/pkg/logger"
	"github.com/pushgate/core/pkg/mailer"
	"github.com/pushgate/core/pkg/postgresql"
	"github.com/pushgate/core/pkg/redis"
	"github.com/pushgate/core/pkg/tron"
)

type Config struct {
	App        *App               `koanf:"app"`
	Cors       *Cors              `koanf:"cors"`
	Limiter    *Limiter           `koanf:"limiter"`
	Logger     *logger.Config     `koanf:"logger"`
	Redis      *redis.Config      `koanf:"redis"`
	Db         *postgresql.Config `koanf:"postgresql"`
	Mailer     *mailer.Config     `koanf:"mailer_templates"`
	Jwt        *jwt.Config        `koanf:"jwt"`
	Encryption *encryption.Config `koanf:"encryption"`
	Tron       *tron.Config       `koanf:"tron"`
}
