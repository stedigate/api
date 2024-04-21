package config

import (
	"github.com/stedigate/core/pkg/blockchains/avalanche"
	"github.com/stedigate/core/pkg/blockchains/ethereum"
	"github.com/stedigate/core/pkg/blockchains/solana"
	"github.com/stedigate/core/pkg/blockchains/tron"
	"github.com/stedigate/core/pkg/encryption"
	"github.com/stedigate/core/pkg/jwt"
	"github.com/stedigate/core/pkg/logger"
	"github.com/stedigate/core/pkg/mailer"
	"github.com/stedigate/core/pkg/postgresql"
	"github.com/stedigate/core/pkg/redis"
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
	Solana     *solana.Config     `koanf:"solana"`
	Ethereum   *ethereum.Config   `koanf:"ethereum"`
	Avalanche  *avalanche.Config  `koanf:"avalanche"`
}
