package module

import (
	"encoding/json"
	"testing"
)

func TestTransactionService_GetTransactionInfoByAddress(t *testing.T) {
	apiUrl := ""
	service := NewTransactionService(apiUrl)
	contractHash := "0x6da01f9b09debd9f4462c2c1cb9850cf22a738bfd6478aa99b22a0d29a5b9e31"
	result, err := service.GetTransactionInfoByAddress(contractHash)
	if err != nil{
		t.Log(err)
	}
	jsonOut, _ := json.Marshal(result)
	t.Logf("%s\n", jsonOut)
}
