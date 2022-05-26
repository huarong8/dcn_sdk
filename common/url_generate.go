package common

import "fmt"

func AccountInfoUrl(url string, address string) string{
	return fmt.Sprintf("%s/accounts/%s", url, address)
}

func GetBlockInfo(url string, address string) string{
	return fmt.Sprintf("%s/blocks/%s", url, address)
}

func GetContractInfo(url string, address string) string{
	return fmt.Sprintf("%s/contracts/%s", url, address)
}

func GetTransactionInfoUrl(url string, address string) string{
	return fmt.Sprintf("%s/transactions/%s", url, address)
}
