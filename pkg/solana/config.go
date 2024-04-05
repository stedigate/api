package solana

type Config struct {
	TrongridApiUrl   string `koanf:"solana.trongrid_api_url"`
	TrongridApiKey   string `koanf:"solana.trongrid_api_key"`
	USDCTokenAddress string `koanf:"solana.usdc_token_address"`
	USDTTokenAddress string `koanf:"solana.usdt_token_address"`
}
