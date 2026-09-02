package hooks

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func NewAuthenticationSvcHooks(NATS *nats.Conn) *AuthenticationSvcHooks {
	return &AuthenticationSvcHooks{
		NATS: NATS,
	}
}

// Notify drive service that
// a new user came in with the
// user ID
func (svc *AuthenticationSvcHooks) NotifyDriveService(
	ctx context.Context,
	userid uint64,
	email string,
	username string,
	password string,
) error {
	log.Println("new user registered! notifying drive service...")
	driveMsg := &DriveNewUserRegMsg{
		UserID: userid,
	}
	msg, err := json.Marshal(driveMsg)
	if err != nil {
		log.Println("error marshaling drive req...")
		return err
	}
	err = svc.NATS.Publish("authentication.user.registered", msg)
	if err != nil {
		log.Println("error publishing req to drive...")
		return err
	}
	return nil
}
