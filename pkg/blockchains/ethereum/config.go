package ethereum

type Config struct {
	ApiUrl          string `koanf:"ethereum.api_url"`
	ApiKey          string `koanf:"ethereum.api_key"`
	EtherscanApiUrl string `koanf:"ethereum.etherscan_api_url"`
	EtherscanApiKey string `koanf:"ethereum.etherscan_api_key"`
	EURCAddress     string `koanf:"ethereum.eurc_address"`
	EURTAddress     string `koanf:"ethereum.eurt_address"`
	USDCAddress     string `koanf:"ethereum.usdc_address"`
	USDTAddress     string `koanf:"ethereum.usdt_address"`
}
