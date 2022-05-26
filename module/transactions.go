package module


import (
	"dcn_sdk/common"
	"dcn_sdk/models"
	"dcn_sdk/util"
)

type TransactionService struct {
	url string
}

func NewTransactionService(url string) *TransactionService{
	return &TransactionService{url: url}
}

func (c *TransactionService)GetTransactionInfoByAddress(address string)(*models.TransactionModel, error){
	url := common.GetTransactionInfoUrl(c.url, address)
	model := &models.TransactionModel{}
	err := util.Get(url, model)
	if err != nil{
		return nil, err
	}
	return model, nil
}
