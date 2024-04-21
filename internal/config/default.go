package config

import (
	"github.com/stedigate/core/pkg/blockchains/avalanche"
	"github.com/stedigate/core/pkg/blockchains/ethereum"
	"github.com/stedigate/core/pkg/blockchains/solana"
	"github.com/stedigate/core/pkg/blockchains/tron"
	"time"

	"github.com/stedigate/core/pkg/encryption"
	"github.com/stedigate/core/pkg/jwt"
	"github.com/stedigate/core/pkg/logger"
	"github.com/stedigate/core/pkg/mailer"
	"github.com/stedigate/core/pkg/postgresql"

	"github.com/stedigate/core/pkg/redis"
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
		Avalanche: &avalanche.Config{
			ApiUrl:          "wss://dev.infura.io/ws/v3/", // Mainnet wss://mainnet.infura.io/ws/v3/
			ApiKey:          "23df91c697a244778c63f1a8bc0b5e7a",
			EtherscanApiUrl: "https://api.etherscan.io/api", // Mainnet https://api.etherscan.io/api
			EtherscanApiKey: "1BEJJ5JDIDFVMASGI17QGP7AYGXPCDA85N",
			EURCAddress:     "0x5e44db7996c682e92a960b65ac713a54ad815c6b", // Mainnet 0xc891eb4cbdeff6e073e859e987815ed1505c2acd
			USDCAddress:     "0x5425890298aed601595a70ab815c96711a31bc65", // Mainnet 0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E
			USDTAddress:     "0xdAC17F958D2ee523a2206206994597C13D831ec7", // Mainnet 0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7
		},
		Ethereum: &ethereum.Config{
			ApiUrl:          "wss://dev.infura.io/ws/v3/", // Mainnet wss://mainnet.infura.io/ws/v3/
			ApiKey:          "23df91c697a244778c63f1a8bc0b5e7a",
			EtherscanApiUrl: "https://api.etherscan.io/api", // Mainnet https://api.etherscan.io/api
			EtherscanApiKey: "1BEJJ5JDIDFVMASGI17QGP7AYGXPCDA85N",
			EURCAddress:     "0x08210F9170F89Ab7658F0B5E3fF39b0E03C594D4", // Mainnet 0x1aBaEA1f7C830bD89Acc67eC4af516284b1bC33c
			EURTAddress:     "0xdAC17F958D2ee523a2206206994597C13D831ec7", // Mainnet 0xC581b735A1688071A1746c968e0798D642EDE491
			USDCAddress:     "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238", // Mainnet 0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48
			USDTAddress:     "0xdAC17F958D2ee523a2206206994597C13D831ec7", // Mainnet 0xdac17f958d2ee523a2206206994597c13d831ec7
		},
		Solana: &solana.Config{
			RpcUrl:         "https://api.devnet.solana.com", // Mainnet https://api.mainnet-beta.solana.com
			WssUrl:         "wss://api.devnet.solana.com",   // Mainnet wss://api.mainnet-beta.solana.com
			TrongridApiKey: "fbb27dd1-4ce1-47c3-8048-230266f64b29",
			EURCAddress:    "HzwqbKZw8HxMN6bF2yFZNrht3c2iXXzpKcFu7uBEDKtr", // Mainnet HzwqbKZw8HxMN6bF2yFZNrht3c2iXXzpKcFu7uBEDKtr
			USDCAddress:    "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU", // Mainnet EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v
			USDTAddress:    "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB", // Mainnet Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB
		},
		Tron: &tron.Config{
			TrongridGrpcUrl:       "grpc.shasta.trongrid.io:50051",  // mainnet grpc.trongrid.io:50051
			TrongridApiUrl:        "https://api.shasta.trongrid.io", // mainnet https://api.trongrid.io
			TrongridApiKey:        "fbb27dd1-4ce1-47c3-8048-230266f64b29",
			TronscanApiUrl:        "https://api.shasta.trongrid.io", // mainnet https://api.trongrid.io
			TronscanApiKey:        "2325af42-0c39-4b1c-acf5-4ed685e4a8e3",
			TrongridJwtKeyId:      "e92357c44cf9471eb7a1a1c3aa7d7527",
			TrongridJwtToken:      "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImU5MjM1N2M0NGNmOTQ3MWViN2ExYTFjM2FhN2Q3NTI3In0.eyJhdWQiOiJ0cm9uZ3JpZC5pbyJ9.oncBAcP-0S1GevfKGyCKKbYfXlYMFmDHxzTFEgDbCHk37ZL--uAYaIac83yNmy6wkU7HMgRBPcDcI0LYsGDk-VQKM4xpil58KjKdbC7TWe7WjF7RqihCrypEdZizXw-OiEUROEYn-t-eT8IP5PkR2y16B_39Srfu54cVL0M48hz1Vxsk8U35fmur24iH2NyrJhhDs5vHV35QHAJQxuCN94akUWEP5Yb2NP36HJVBdsVYZJG6SpFS0i43nACvRaTuHfxFmC4qHL8XGH7tUnhgWKz4sJkk-QG0vi9-9mZEqsSCqkhooK7vXmagdK5MhKtE66FNc43A6tuYyfVpny5gsw",
			USDTAddress:           "TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs", // mainnet TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t
			Trc20ContractAbi:      "",
			Trc20ContractNetwork:  "shasta", // mainnet mainnet
			Trc20ContractCurrency: "TRX",
			Trc20ContractDecimals: 6,
			Trc20ContractFeeLimit: 300000,
		},
	}
}
