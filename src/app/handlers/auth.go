package handlers

import (
	"context"
	"encoding/json"
	"kyc/src/app"
	kyc_workflow "kyc/src/app/kyc_workflows"
	model "kyc/src/app/models"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.temporal.io/sdk/client"
)

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User Sign up and trigger background process
func SignUp(w http.ResponseWriter, r *http.Request) {
	new_user := &NewUser{}
	err := json.NewDecoder(r.Body).Decode(new_user)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.WithError(err).Error("Something went wrong")
		w.WriteHeader(http.StatusInternalServerError)
		resp := make(map[string]string)
		resp["message"] = "Some Error Occurred"
		json.NewEncoder(w).Encode(resp)
	} else {

		user := model.User{
			Email: new_user.Email,
			Id:    uuid.New().String(),
		}

		triggerWorkflow(user)

		log.Infof("Created User %s", user.Id)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func triggerWorkflow(user model.User) string {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        user.Id,
		TaskQueue: app.KYCTaskQueue,
	}
	// create workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, kyc_workflow.KYCWorkflow, user)
	if err != nil {
		log.WithError(err).Error("unable to complete Workflow")
	}
	log.Infof("Started workflow %s with run id %s", we.GetID(), we.GetRunID())
	return we.GetRunID()
}
