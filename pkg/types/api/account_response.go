package api

import (
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type AccountResponse struct {
	Alias               string            `json:"alias" example:"personal-account"`
	Address             ethcommon.Address `json:"address" example:"0x1abae27a0cbfb02945720425d3b80c7e09728534" swaggertype:"string"`
	PublicKey           string            `json:"publicKey" example:"048e66b3e549818ea2cb354fb70749f6c8de8fa484f7530fc447d5fe80a1c424e4f5ae648d648c980ae7095d1efad87161d83886ca4b6c498ac22a93da5099014a"`
	CompressedPublicKey string            `json:"compressedPublicKey"`
	TenantID            string            `json:"tenantID" example:"tenantFoo"`
	OwnerID             string            `json:"ownerID,omitempty" example:"foo"`
	Attributes          map[string]string `json:"attributes,omitempty"`
	CreatedAt           time.Time         `json:"createdAt" example:"2020-07-09T12:35:42.115395Z"`
	UpdatedAt           time.Time         `json:"updatedAt,omitempty" example:"2020-07-09T12:35:42.115395Z"`
}
