# DCN-SDK介绍

## SDK下载地址：
**golang**: <https://github.com/huarong8/dcn_sdk>
**详细的文档参考**: <https://dcn-doc.readthedocs.io/en/latest/>

## 简介
DCN-SDK实现了区块服务、交易服务等接口，同时还提供了接口使用示例说明，开发者可以调用该SDK方便快捷的生成DCN主链的快速接入

## 相关接口

1. 根据BlockIdHash获取区块的信息
    1. 接口名: GetBlockInfoByHash
    2. 示例: 
    ```golang
    apiUrl := ""
   	blockService := NewBlockService(apiUrl)
   	blockHash := "0xca42c55ce0c708d92a77954660a0865557aaea5f8491a87891e0491bc482b60d"
   	result, err := blockService.GetBlockInfoByHash(blockHash)
   	if err != nil{
   		t.Log(err)
   	}
   	jsonOut, _ := json.Marshal(result)
   	t.Logf("%s\n", jsonOut)
   	//output
   	//{"dateVerified":1653547894719,"difficulty":1758681,"extraData":"0xd683010a11846765746886676f312e3138856c696e7578","gasLimit":15000000,"gasUsed":0,"hash":"0xca42c55ce0c708d92a77954660a0865557aaea5f8491a87891e0491bc482b60d","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x9b8f1667f34cd69ffeeedc2a19861149b552f82e","mixHash":"0xddfccf8be27c647e8c1b5b59bf57fce08cf6c1443f645ee7cc5d0f50c092fb36","number":88330,"parentHash":"0x61a9dc8c3343406dd0768c60091de158bf46b3dd2f4fd3a9f288f2a2b08aed6d","receiptRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":537,"stateRoot":"0x54ba5ac7a37f92ddefa8cdef601d0e8cb2687fa33659033a9969b728e9111e03","timestampISO":"2022-05-26T06:51:34Z","timestampVerifiedISO":"2022-05-26T06:51:34.719Z","totalDifficulty":165079341925,"transactionCount":0,"transactionsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","uncles":[]}

    ```

1. 根据AccountIdHash获取账户信息
    1. 接口名 GetAccountInfo
    2. 示例：
    ```golang
   func TestAccountService_GetAccountInfo(t *testing.T) {
   	apiUrl := ""
   	service := NewAccount(apiUrl)
   	accountHash := "0xca42c55ce0c708d92a77954660a0865557aaea5f8491a87891e0491bc482b60d"
   	result, err := service.GetAccountInfo(accountHash)
   	if err != nil{
   		t.Log(err)
   	}
   	jsonOut, _ := json.Marshal(result)
   	t.Logf("%s\n", jsonOut)
   } 
   ```

1. 根据ContractIdHash获取合约信息
    1. 接口名 GetContractInfoByAddress
    2. 示例：
    ```golang
   
   func TestContractService_GetContractInfoByAddress(t *testing.T) {
   	apiUrl := ""
   	service := NewContractService(apiUrl)
   	contractHash := "0x6da01f9b09debd9f4462c2c1cb9850cf22a738bfd6478aa99b22a0d29a5b9e31"
   	result, err := service.GetContractInfoByAddress(contractHash)
   	if err != nil{
   		t.Log(err)
   	}
   	jsonOut, _ := json.Marshal(result)
   	t.Logf("%s\n", jsonOut)
   }
   ```

1. 根据TransactionsId获取交易信息
    1. 接口名 GetTransactionInfoByAddress
    2. 示例：
    ```golang
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

   ```   
   
