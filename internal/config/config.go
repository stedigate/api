package config

import (
	"github.com/stedigate/core/pkg/encryption"
	"github.com/stedigate/core/pkg/jwt"
	"github.com/stedigate/core/pkg/logger"
	"github.com/stedigate/core/pkg/mailer"
	"github.com/stedigate/core/pkg/postgresql"
	"github.com/stedigate/core/pkg/redis"
	"github.com/stedigate/core/pkg/tron"
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
