package jwt

import "context"

// UserClaims represent raw claims extracted from an authentication method
type UserClaims struct {
	TenantID string
}

type CustomClaims struct {
	TenantID string `json:"tenant_id"`
}

func (claims *CustomClaims) Validate(_ context.Context) error {
	// TODO: Apply validation on custom claims if needed, currently no validation is needed
	return nil
}
