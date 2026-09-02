package hooks

import "github.com/nats-io/nats.go"

type AuthenticationSvcHooks struct {
	NATS *nats.Conn
}

type DriveNewUserRegMsg struct {
	UserID uint64 `json:"user_id"`
}