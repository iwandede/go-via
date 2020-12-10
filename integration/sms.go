package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/iwandede/go-via/config"
	"github.com/iwandede/go-via/models/integration"
)

func MakeRequestSMS(config *config.Config, DataDTO []*integration.SMSRequest) (interface{}, error) {
	var Response []*integration.ResponseData

	Options := fmt.Sprintf("%s/message/send", config.Integration.SMS.URL)

	Payload, err := json.Marshal(&DataDTO)
	if err != nil {
		return nil, fmt.Errorf("1. %v", err.Error())
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", Options, bytes.NewReader(Payload))

	if err != nil {
		return nil, fmt.Errorf("2. %v", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", config.Integration.SMS.Token)
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("3. %v", err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("4. %v", err)
	}

	err = json.Unmarshal([]byte(body), &Response)
	if err != nil {
		return nil, fmt.Errorf("5. %v", err)
	}

	return Response, nil
}
