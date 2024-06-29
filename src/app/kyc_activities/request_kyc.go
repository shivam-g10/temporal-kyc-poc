package kyc_activity

// import (
// 	"context"
// 	model "kyc/src/app/models"
// 	"time"

// 	log "github.com/sirupsen/logrus"
// 	"go.temporal.io/sdk/activity"
// 	"go.temporal.io/sdk/workflow"
// )

// const NOTIFICATION_INTERVAL = time.Second * 10

// func RequestKYC(ctx context.Context, user model.User) (string, error) {
// 	for {
// 		log.Infof("Sent REQUEST_KYC message to %s to complete KYC", user.Id)
// 		var kyc_request model.KYCRequest
// 		kyc_signal_channel := workflow.GetSignalChannel(ctx, "new_kyc_request")
// 		activity.RecordHeartbeat(ctx)
// 		select {
// 		case <-ctx.Done():
// 			return "", ctx.Err()
// 		case <-time.After(NOTIFICATION_INTERVAL):
// 		}
// 	}
// }
