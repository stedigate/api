package solana

type Config struct {
	RpcUrl         string `koanf:"solana.rpc_api_url"`
	WssUrl         string `koanf:"solana.wss_api_url"`
	TrongridApiKey string `koanf:"solana.trongrid_api_key"`
	EURCAddress    string `koanf:"solana.eurc_address"`
	EURTAddress    string `koanf:"solana.eurt_address"`
	USDCAddress    string `koanf:"solana.usdc_address"`
	USDTAddress    string `koanf:"solana.usdt_address"`
}
