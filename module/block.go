package module

import (
	"dcn_sdk/common"
	"dcn_sdk/models"
	"dcn_sdk/util"
)

type BlockService struct {
	url string
}

func NewBlockService(url string)*BlockService{
	return &BlockService{
		url: url,
	}
}

func (b *BlockService)GetBlockInfoByHash(address string) (*models.BlockModel, error){
	url := common.GetBlockInfo(b.url, address)
	model := &models.BlockModel{}
	err := util.Get(url, model)
	if err != nil{
		return nil, err
	}
	return model, nil
}


