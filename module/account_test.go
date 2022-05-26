package module

import (
	"encoding/json"
	"testing"
)

func TestAccountService_GetAccountInfo(t *testing.T) {
	apiUrl := ""
	service := NewAccount(apiUrl)
	accountHash := "0x6bac49893b1583315c7a489620fe80714dd76932"
	result, err := service.GetAccountInfo(accountHash)
	if err != nil{
		t.Log(err)
	}
	jsonOut, _ := json.Marshal(result)
	t.Logf("%s\n", jsonOut)
}
