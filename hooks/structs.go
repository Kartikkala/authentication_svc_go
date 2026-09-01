package hooks

import "github.com/nats-io/nats.go"

type AuthenticationSvcHooks struct {
	NATS *nats.Conn
}