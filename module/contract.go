package module

import (
	"dcn_sdk/common"
	"dcn_sdk/models"
	"dcn_sdk/util"
)

type ContractService struct {
	url string
}

func NewContractService(url string) *ContractService{
	return &ContractService{url: url}
}

func (c *ContractService)GetContractInfoByAddress(address string)(*models.ContractModel, error){
	url := common.GetContractInfo(c.url, address)
	model := &models.ContractModel{}
	err := util.Get(url, model)
	if err != nil{
		return nil, err
	}
	return model, nil
}
