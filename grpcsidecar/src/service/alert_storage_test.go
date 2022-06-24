package service_test

import (
	"grsidecar/sample"
	"testing"
)

func TestCreateAlertStore(tst *testing.T) {
	tst.Parallel()

	alertWithoutCID := sample.NewAlert()
	alertWithoutCID.Cid = ""

	testCases := []struct {
		
	}
}
