package tron

type Config struct {
	TrongridApiUrl        string `koanf:"tron.trongrid_api_url"`
	TrongridGrpcUrl       string `koanf:"tron.trongrid_grpc_url"`
	TrongridApiKey        string `koanf:"tron.trongrid_api_key"`
	TronscanApiUrl        string `koanf:"tron.tronscan_api_url"`
	TronscanApiKey        string `koanf:"tron.tronscan_api_key"`
	TrongridJwtKeyId      string `koanf:"tron.trongrid_jwt_key_id"`
	TrongridJwtToken      string `koanf:"tron.trongrid_jwt_token"`
	Trc20ContractAbi      string `koanf:"tron.trc20_contract_abi"`
	Trc20ContractNetwork  string `koanf:"tron.trc20_contract_network"`
	Trc20ContractCurrency string `koanf:"tron.trc20_contract_currency"`
	Trc20ContractDecimals int    `koanf:"tron.trc20_contract_decimals"`
	Trc20ContractFeeLimit int    `koanf:"tron.trc20_contract_fee_limit"`
	USDTAddress           string `koanf:"tron.usdt_address"`
}
