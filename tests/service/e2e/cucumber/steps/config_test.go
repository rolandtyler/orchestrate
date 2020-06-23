// +build unit

package steps

import (
	"testing"

	"github.com/spf13/pflag"
)

func TestInitFlags(t *testing.T) {
	flgs := pflag.NewFlagSet("stepTable", pflag.ContinueOnError)
	InitFlags(flgs)
}
