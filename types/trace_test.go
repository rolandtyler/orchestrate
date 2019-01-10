package types

import (
	"math/big"
	"testing"
)

func TestTrace(t *testing.T) {
	tr := NewTrace()

	tr.Chain().ID = big.NewInt(1024)
	if tr.Chain().ID.Text(16) != "400" {
		t.Errorf("Trace: Expected Chain ID %q but got %q", "1024", tr.Chain().ID.Text(400))
	}

	tr.Receiver().ID = "afg"
	if tr.Receiver().ID != "afg" {
		t.Errorf("Trace: Expected Recveiver ID %q but got %q", "afg", tr.Receiver().ID)
	}

	tr.Sender().ID = "fjt"
	if tr.Sender().ID != "fjt" {
		t.Errorf("Trace: Expected Sender ID %q but got %q", "fjt", tr.Sender().ID)
	}

	tr.Call().MethodID = "xyz"
	if tr.Call().MethodID != "xyz" {
		t.Errorf("Trace: Expected Method ID %q but got %q", "xyz", tr.Call().MethodID)
	}

	tr.Tx().SetNonce(10)
	if tr.Tx().Nonce() != 10 {
		t.Errorf("Trace: Expected Nonce %v but got %v", 10, tr.Tx().Nonce())
	}

	tr.Receipt().Status = 1
	if tr.Receipt().Status != 1 {
		t.Errorf("Trace: Expected Status %v but got %v", 1, tr.Receipt().Status)
	}

	tr.Reset()

	if tr.Chain().ID.Text(16) != "0" {
		t.Errorf("Trace: Expected Chain ID %q but got %q", "0", tr.Chain().ID.Text(16))
	}
	if tr.Receiver().ID != "" {
		t.Errorf("Trace: Expected Recveiver ID %q but got %q", "", tr.Receiver().ID)
	}
	if tr.Sender().ID != "" {
		t.Errorf("Trace: Expected Sender ID %q but got %q", "", tr.Sender().ID)
	}
	if tr.Call().MethodID != "" {
		t.Errorf("Trace: Expected Method ID %q but got %q", "", tr.Call().MethodID)
	}
	if tr.Tx().Nonce() != 0 {
		t.Errorf("Trace: Expected Nonce %v but got %v", 0, tr.Tx().Nonce())
	}

	if tr.Receipt().Status != 0 {
		t.Errorf("Trace: Expected Status %v but got %v", 0, tr.Receipt().Status)
	}
}
