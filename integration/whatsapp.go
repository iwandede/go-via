package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/iwandede/go-via/config"
	"github.com/iwandede/go-via/models/integration"
)

func MakeRequestWA(config *config.Config, DataDTO *integration.WhatsappDTO) (*integration.ResponseData, error) {
	var Response *integration.ResponseData

	Options := fmt.Sprintf("%s?", config.Integration.Whatsapp.URL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", Options, nil)

	if err != nil {
		return nil, fmt.Errorf("1. %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("username", config.Integration.Whatsapp.Username)
	q.Add("password", config.Integration.Whatsapp.Password)
	q.Add("GSM", DataDTO.GSM)
	q.Add("text", DataDTO.Text)
	q.Add("output", "json")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("2. %v", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("3. %v", err)
	}

	err = json.Unmarshal([]byte(body), &Response)
	if err != nil {
		return nil, fmt.Errorf("3. %v", err)
	}

	return Response, nil
}
