package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/iwandede/go-via/integration"
	"github.com/iwandede/go-via/lib"
	"github.com/iwandede/go-via/models"
	thirdparty "github.com/iwandede/go-via/models/integration"
)

func (Notification Controllers) SendWhatsApp(w http.ResponseWriter, r *http.Request) {
	var DataRequest *thirdparty.WhatsappDTO

	if err := json.NewDecoder(r.Body).Decode(&DataRequest); err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err))
		return
	}

	SendService, err := integration.MakeRequestWA(Notification.Config, DataRequest)
	if err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err.Error()))
		return
	}

	Response := &models.ResponseThirdParty{
		Status:      SendService.Results[0].Status,
		Messageid:   SendService.Results[0].Messageid,
		Destination: SendService.Results[0].Destination,
	}

	json.NewEncoder(w).Encode(lib.ResponseSuccess(Response))
	return
}

func (Notification Controllers) SendSMS(w http.ResponseWriter, r *http.Request) {
	var DataRequest []*thirdparty.SMSRequest
	//var Response []*models.ResponseThirdParty
	if err := json.NewDecoder(r.Body).Decode(&DataRequest); err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err))
		return
	}

	SendService, err := integration.MakeRequestSMS(Notification.Config, DataRequest)
	if err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err.Error()))
		return
	}
	/*for _, row := range SendService {
		Response = append(Response, &models.ResponseThirdParty{
			Status:      row.Status,
			Messageid:   row.Message,
			Destination: row.PhoneNumber,
		})
	}*/

	// data partner
	json.NewEncoder(w).Encode(lib.ResponseSuccess(SendService))
	return
}
