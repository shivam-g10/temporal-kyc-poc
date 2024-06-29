package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type GovernmentId struct {
	AADHAAR         string
	VOTER_ID        string
	DRIVERS_LICENSE string
}

var GovernmentIds = new_government_ids()

func new_government_ids() *GovernmentId {
	return &GovernmentId{
		AADHAAR:         "AADHAAR",
		VOTER_ID:        "VOTER_ID",
		DRIVERS_LICENSE: "DRIVERS_LICENSE",
	}
}

type NewKYCRequest struct {
	FileURL string `json:"file_url"`
	IdType  string `json:"id_type"`
	Id      string `json:"id"`
}

type KYCRequest struct {
	FileURL string `json:"file_url"`
	IdType  string `json:"id_type"`
	Id      string `json:"id"`
	UserId  string `json:"user_id"`
}

type KYCAction struct {
	Approve bool `json:"approve"`
}

type KYCActionResponse struct {
	UserId  string `json:"user_id"`
	Approve bool   `json:"approve"`
}

// KYC Submitted by User
func KycSubmit(w http.ResponseWriter, r *http.Request) {
	new_kyc := &NewKYCRequest{}
	err := json.NewDecoder(r.Body).Decode(new_kyc)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.WithError(err).Error("Something went wrong")
		w.WriteHeader(http.StatusInternalServerError)
		resp := make(map[string]string)
		resp["message"] = "Some Error Occurred"
		json.NewEncoder(w).Encode(resp)
	} else {
		// validate id types
		var id_type = GovernmentIds.AADHAAR

		if GovernmentIds.AADHAAR != new_kyc.IdType {
			log.WithError(err).Error("Something went wrong")
			w.WriteHeader(http.StatusBadRequest)
			resp := make(map[string]string)
			resp["message"] = "Government Id Type Error"
			json.NewEncoder(w).Encode(resp)
			return
		}

		kyc_request := &KYCRequest{
			FileURL: new_kyc.FileURL,
			IdType:  id_type,
			UserId:  r.PathValue("user_id"),
			Id:      new_kyc.Id,
		}
		log.Infof("Created KYC Request for %s", kyc_request.UserId)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(kyc_request)
	}
}

// Action on KYC by Agent
func ActionKYC(w http.ResponseWriter, r *http.Request) {
	kyc_action := &KYCAction{}
	err := json.NewDecoder(r.Body).Decode(kyc_action)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.WithError(err).Error("Something went wrong")
		w.WriteHeader(http.StatusInternalServerError)
		resp := make(map[string]string)
		resp["message"] = "Some Error Occurred"
		json.NewEncoder(w).Encode(resp)
	} else {
		// trigger action
		log.Infof("KYC Request for %s is %t", r.PathValue("user_id"), kyc_action.Approve)
		w.WriteHeader(http.StatusOK)
		kyc_action_result := &KYCActionResponse{
			UserId:  r.PathValue("user_id"),
			Approve: kyc_action.Approve,
		}
		json.NewEncoder(w).Encode(kyc_action_result)
	}
}
