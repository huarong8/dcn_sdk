package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(urlStr string, result interface{}) error{
	client := http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil{
		return err
	}
	req.Header.Add("Context-Type", "application/json")
	rsp, err := client.Do(req)
	if err != nil{
		return err
	}
	defer rsp.Body.Close()
	body, _ := ioutil.ReadAll(rsp.Body)

	err = json.Unmarshal(body, &result)
	if err != nil{
		return err
	}
	return nil
}

func Post(urlStr string, params interface{}, result interface{}) error{
	client := http.Client{}
	byteParams, err := json.Marshal(params)
	if err != nil{
		return err
	}
	req, err := http.NewRequest("POST", urlStr, bytes.NewReader(byteParams))
	rsp, err := client.Do(req)
	if err != nil{
		return err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil{
		return err
	}
	err = json.Unmarshal(body, result)
	if err != nil{
		return err
	}
	return nil
}
