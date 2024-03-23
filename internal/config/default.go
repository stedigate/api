package config

import (
	"github.com/pushgate/core/pkg/tron"
	"time"

	"github.com/pushgate/core/pkg/encryption"
	"github.com/pushgate/core/pkg/jwt"
	"github.com/pushgate/core/pkg/logger"
	"github.com/pushgate/core/pkg/mailer"
	"github.com/pushgate/core/pkg/postgresql"

	"github.com/pushgate/core/pkg/redis"
)

func Default() *Config {
	return &Config{
		App: &App{
			Port:    4000,
			Env:     "development",
			Version: "1.0.0",
		},
		Cors: &Cors{
			TrustedOrigins: []string{"http://localhost:3000"},
		},
		Limiter: &Limiter{
			Rps:     2,
			Burst:   4,
			Enabled: true,
		},
		Jwt: &jwt.Config{
			PrivatePem: `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIF0V3x7RkGyiVZGXCny8vtnBajmD2TOT2TkhounyUkBR
-----END PRIVATE KEY-----
`,
			PublicPem: `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEA1JsMvBD61BAYv8+JZtvex1K7Y1CgYeNnO9WMhgxNrv8=
-----END PUBLIC KEY-----
`,
			Expiration:        30 * 24 * time.Hour,
			RefreshExpiration: 3 * 30 * 24 * time.Hour,
			CookieTokenName:   "__Secure_token",
		},
		Redis: &redis.Config{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			Db:       0,
		},
		Db: &postgresql.Config{
			Dsn:          "postgres://sail:password@localhost:5432/stedigate?sslmode=disable",
			Host:         "localhost",
			Port:         "5432",
			Username:     "sail",
			Password:     "password",
			Database:     "stedigate",
			SSLMode:      "false",
			MaxOpenConns: 64,
			MaxIdleConns: 64,
			MaxIdleTime:  "15m",
		},

		Mailer: &mailer.Config{
			Host:         "smtp.mailtrap.io",
			Port:         2525,
			Username:     "info",
			Password:     "info",
			ResendAPIKey: "re_KV12BrjV_56WDs6qW17AqAhAeEXh4o9ZB",
			Sender:       "noreply@mail.stedigate.bid",
		},
		Logger: &logger.Config{
			Level:    "debug",
			Path:     "logs/error.log",
			Env:      "development",
			Encoding: "console",
		},
		Encryption: &encryption.Config{
			Key:       []byte("39kQ2y7BgQQOXAzlUY6hnSqmQdRFH3Yy"),
			Algorithm: "AES-256",
		},
		Tron: &tron.Config{
			TrongridApiUrl:        "https://api.shasta.trongrid.io",
			TrongridApiKey:        "fbb27dd1-4ce1-47c3-8048-230266f64b29",
			TrongridJwtKeyId:      "e92357c44cf9471eb7a1a1c3aa7d7527",
			TrongridJwtToken:      "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImU5MjM1N2M0NGNmOTQ3MWViN2ExYTFjM2FhN2Q3NTI3In0.eyJhdWQiOiJ0cm9uZ3JpZC5pbyJ9.oncBAcP-0S1GevfKGyCKKbYfXlYMFmDHxzTFEgDbCHk37ZL--uAYaIac83yNmy6wkU7HMgRBPcDcI0LYsGDk-VQKM4xpil58KjKdbC7TWe7WjF7RqihCrypEdZizXw-OiEUROEYn-t-eT8IP5PkR2y16B_39Srfu54cVL0M48hz1Vxsk8U35fmur24iH2NyrJhhDs5vHV35QHAJQxuCN94akUWEP5Yb2NP36HJVBdsVYZJG6SpFS0i43nACvRaTuHfxFmC4qHL8XGH7tUnhgWKz4sJkk-QG0vi9-9mZEqsSCqkhooK7vXmagdK5MhKtE66FNc43A6tuYyfVpny5gsw",
			Trc20ContractAddress:  "TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs",
			Trc20ContractAbi:      "",
			Trc20ContractNetwork:  "shasta",
			Trc20ContractCurrency: "TRX",
			Trc20ContractDecimals: 6,
			Trc20ContractFeeLimit: 300000,
		},
	}
}
