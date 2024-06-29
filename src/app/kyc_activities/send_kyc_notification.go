package kyc_activity

import (
	"context"
	"errors"
	model "kyc/src/app/models"

	log "github.com/sirupsen/logrus"
)

type SendNotificationData struct {
	User      model.User
	NotifType string
	Message   string
}

func SendKYCNotification(ctx context.Context, data SendNotificationData) (string, error) {
	switch data.NotifType {
	case "REQUEST_KYC":
		log.Infof("Sent REQUEST_KYC message to %s to complete KYC: %s", data.User.Id, data.Message)
		return data.Message, nil
	case "REQUEST_KYC_ACTION":
		log.Infof("Sent REQUEST_KYC_ACTION message to %s to complete KYC: %s", data.User.Id, data.Message)
		return data.Message, nil
	default:
		err := errors.New("message not configured")
		log.Error(err)
		return "", err
	}
}
