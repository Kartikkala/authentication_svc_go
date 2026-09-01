package hooks

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

func NewAuthenticationSvcHooks(NATS *nats.Conn) *AuthenticationSvcHooks {
	return &AuthenticationSvcHooks{
		NATS: NATS,
	}
}

// Move to the drive service!
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

	// TODO: Complete this
	// svc.NATS.Publish("auth.newuser")
	
	// _, err = svc.authorizationSvc.WriteRelationship(ctx, "node", node.ID.String(), "owner", "user", strconv.FormatUint(userid, 10))
	// if err != nil {
	// 	return err
	// }
	return nil
}
