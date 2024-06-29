package main

import (
	"kyc/src/app"
	kyc_activity "kyc/src/app/kyc_activities"
	kyc_workflow "kyc/src/app/kyc_workflows"

	log "github.com/sirupsen/logrus"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, app.KYCTaskQueue, worker.Options{})

	// Register workers and activities
	w.RegisterWorkflow(kyc_workflow.RequestKYCWorkflow)
	w.RegisterWorkflow(kyc_workflow.KYCWorkflow)
	w.RegisterActivity(kyc_activity.SendKYCNotification)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}

}
