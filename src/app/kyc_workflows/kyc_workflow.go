package kyc_workflow

import (
	"fmt"
	"kyc/src/app"
	kyc_activity "kyc/src/app/kyc_activities"
	model "kyc/src/app/models"
	"time"

	log "github.com/sirupsen/logrus"
	"go.temporal.io/sdk/workflow"
)

// Workflow to manage
func KYCWorkflow(ctx workflow.Context, user model.User) (string, error) {
	// activities should complete in 5 seconds
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// set up signal to listen for new KYC signal
	var kyc_request model.KYCRequest
	kyc_req_signal_channel := workflow.GetSignalChannel(ctx, app.NEW_KYC_SIGNAL)

	selector := workflow.NewSelector(ctx)
	// trigger recurring notification
	res := workflow.ExecuteChildWorkflow(ctx, RequestKYCWorkflow, user)

	// Group child workflow and signal futures in a single selector
	// We will let the child workflow run and forward it the NEW_KYC_SIGNAL
	// to tell it when to stop
	selector.AddFuture(res, func(f workflow.Future) {
		f.Get(ctx, nil)
	})
	selector.AddReceive(kyc_req_signal_channel, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &kyc_request)
		log.Info("RECEIVED SIGNAL")
		// ping child to stop
		res.SignalChildWorkflow(ctx, app.NEW_KYC_SIGNAL, kyc_request).Get(ctx, nil)
		log.Info("SENT SIGNAL TO CHILD")
	})
	// Block for futures
	selector.Select(ctx)

	// Wait for action to be taken by agent
	// No need to poll in this case since we don't need to constantly notify the agent
	// Maybe we can add a 24 hour service level agreement for escalation to supervisor
	var kyc_request_action model.KYCAction
	kyc_signal_channel := workflow.GetSignalChannel(ctx, app.KYC_ACTION_SIGNAL)
	kyc_signal_channel.Receive(ctx, &kyc_request_action)

	action := "Rejected"
	if kyc_request_action.Approve {
		action = "Approved"
	}

	// on kyc action send notification to user
	request_kyc_notification := &kyc_activity.SendNotificationData{
		User:      user,
		NotifType: "REQUEST_KYC_ACTION",
		Message:   fmt.Sprintf("Your KYC request has been %s", action),
	}

	// Trigger notification for KYC results
	var result string
	err := workflow.ExecuteActivity(ctx, kyc_activity.SendKYCNotification, request_kyc_notification).Get(ctx, &result)

	return result, err
}
