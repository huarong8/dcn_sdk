package module

import (
	"dcn_sdk/common"
	"dcn_sdk/models"
	"dcn_sdk/util"
)

type AccountService struct{
	url string
}

func NewAccount(url string) *AccountService{
	return &AccountService{url: url}
}

func(a *AccountService) GetAccountInfo(addressHash string) (*models.AccountModel, error){
	url := common.AccountInfoUrl(a.url, addressHash)
	model := models.AccountModel{}
	err := util.Get(url,&model)
	if err != nil{
		return nil, err
	}
	return &model, nil
}
