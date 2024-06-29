package kyc_workflow

import (
	"kyc/src/app"
	kyc_activity "kyc/src/app/kyc_activities"
	model "kyc/src/app/models"
	"time"

	log "github.com/sirupsen/logrus"
	"go.temporal.io/sdk/workflow"
)

// Dummy notification time interval
const NOTIFICATION_INTERVAL = time.Second * 30

// Keep sending notifications to user until we get a KYC submission
func RequestKYCWorkflow(ctx workflow.Context, user model.User) (*model.KYCRequest, error) {
	// activity should complete in 5 seconds
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	request_kyc_notification := &kyc_activity.SendNotificationData{
		User:      user,
		NotifType: "REQUEST_KYC",
		Message:   "Please submit KYC",
	}

	// trigger Notification activity
	err := workflow.ExecuteActivity(ctx, kyc_activity.SendKYCNotification, request_kyc_notification).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	// set up signal
	var kyc_request model.KYCRequest
	kyc_signal_channel := workflow.GetSignalChannel(ctx, app.NEW_KYC_SIGNAL)

	// Create selector to wait for signal as well as timer
	// Which ever one triggers first will decide what happens in the if check later
	selector := workflow.NewSelector(ctx)
	selector.AddReceive(kyc_signal_channel, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &kyc_request)
		log.Info("RECEIVED IN CHILD")
	})
	// Wait for Interval
	// If timer triggers first then signal will be received in the next run
	selector.AddFuture(workflow.NewTimer(ctx, NOTIFICATION_INTERVAL), func(f workflow.Future) {
		err = f.Get(ctx, nil)
	})
	selector.Select(ctx)

	// If we didn't get signal then start workflow again
	if kyc_request.Id == "" { // KYC request not completed in time
		err := workflow.NewContinueAsNewError(ctx, RequestKYCWorkflow, user)
		return nil, err
	} else {
		return &kyc_request, nil
	}
}
