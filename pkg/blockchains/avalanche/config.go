package avalanche

type Config struct {
	ApiUrl          string `koanf:"avalanche.api_url"`
	ApiKey          string `koanf:"avalanche.api_key"`
	EtherscanApiUrl string `koanf:"avalanche.etherscan_api_url"`
	EtherscanApiKey string `koanf:"avalanche.etherscan_api_key"`
	EURCAddress     string `koanf:"avalanche.eurc_address"`
	EURTAddress     string `koanf:"avalanche.eurt_address"`
	USDCAddress     string `koanf:"avalanche.usdc_address"`
	USDTAddress     string `koanf:"avalanche.usdt_address"`
}
