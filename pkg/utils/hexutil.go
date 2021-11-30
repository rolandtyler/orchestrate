package utils

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func HexBigInt(v *big.Int) *hexutil.Big {
	return ToPtr(hexutil.Big(*v)).(*hexutil.Big)
}

type IStringObj interface {
	String() string
}

func HexToString(obj IStringObj) string {
	if obj == nil {
		return ""
	}
	
	return obj.String()
}

func StringToHexBytes(v string) hexutil.Bytes {
	if v == "" {
		return nil
	}
	
	if vb, err := hexutil.Decode(v); err == nil {
		return vb
	}
	
	return nil
}
func StringToHexInt(v string) *hexutil.Big {
	if v == "" {
		return nil
	}

	if bv, ok := new(big.Int).SetString(v, 10); ok {
		return (*hexutil.Big)(bv)
	}

	return nil
}

func StringToUint64(v string) *uint64 {
	if v == "" {
		return nil
	}

	
	if vi, err := strconv.ParseUint(v, 10, 64); err == nil {
		return &vi
	}

	return nil
}

func IsHexString(s string) bool {
	_, err := hexutil.Decode(s)
	return err == nil
}
