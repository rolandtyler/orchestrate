package keymanager

type PayloadRequest struct {
	Data      string `json:"data" validate:"required" example:"my data to sign"`
	Namespace string `json:"namespace,omitempty" example:"tenant_id"`
}

type VerifyPayloadRequest struct {
	Data      string `json:"data" validate:"required" example:"my data to sign"`
	Signature string `json:"signature" validate:"required,isHex" example:"0x6019a3c8..."`
	Address   string `json:"address" validate:"required,isHex" example:"0x905B88EFf8Bda1543d4d6f4aA05afef143D27E18"`
}
