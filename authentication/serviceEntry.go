package authentication

import (
	"context"
	"log"
)

func NewServiceWithHooks(authenticationSvc *Service) *SvcWithHooks {
	return &SvcWithHooks{
		authSvc:                authenticationSvc,
		registrationHooksAfter: []RegisterHook{},
	}
}

func (hl *SvcWithHooks) AppendAfterRegistrationHook(hook RegisterHook) {
	hl.registrationHooksAfter = append(hl.registrationHooksAfter, hook)
}

func (hl *SvcWithHooks) RegisterUserEntry(
	ctx context.Context,
	email string,
	username string,
	password string,
) error {
	uid, err := hl.authSvc.RegisterService(ctx, email, username, password)
	if err != nil {
		return err
	}

	for _, hook := range hl.registrationHooksAfter {
		err := hook(ctx, uid, email, username, password)
		if err != nil {
			log.Printf("error running hook : %s", err)
		}
	}
	return nil
}

func (hl *SvcWithHooks) LoginUserEntry(
	ctx context.Context,
	email string,
	password string,
) (*User, error) {
	return hl.authSvc.LoginService(ctx, email, password)
}

func (hl *SvcWithHooks) GenerateTokenEntry(
	user *User,
) (string, error) {
	return hl.authSvc.GenerateToken(user)
}

func (hl *SvcWithHooks) DecodeTokenEntry(
	token *string,
) (*CustomClaims, error) {
	return hl.authSvc.DecodeToken(token)
}
