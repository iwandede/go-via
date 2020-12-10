package integration

type WhatsappDTO struct {
	GSM      string `json:"gsm"`
	Text     string `json:"text"`
}

type ResponseData struct {
	Results []struct{
		Status      string `json:"status"`
		Messageid   string `json:"messageid"`
		Destination string `json:"destination"`
	} `json:"results"`
}