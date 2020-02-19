package router

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type SendHttpRequest struct {
}

func SendHttpRequestProvider() SendHttpRequest {
	return SendHttpRequest{}
}

func (s *SendHttpRequest) Send(mobile string) error {
	url := "http://localhost:2020/api/charge"
	var jsonStr = map[string]string{"Mobile": mobile}
	marshalData, _ := json.Marshal(jsonStr)
	_, err := http.Post(url, "application/json",bytes.NewBuffer(marshalData))
	if err != nil {
		log.Printf("Wallet not charged: %v",err.Error())
		return err
	}
	return nil
}
