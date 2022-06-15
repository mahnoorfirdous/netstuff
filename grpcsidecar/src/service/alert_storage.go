package service

import (
	"errors"
	"fmt"
	"grsidecar/pbgen"
	"sync"

	"github.com/hidevopsio/hiboot/pkg/utils/copier"
)

type AlertStore struct {
	take  sync.RWMutex
	store map[string]*pbgen.AlertDetail
}

var AlertDuplicateFound = errors.New("Alert is stored and will be serviced soon")

func (as *AlertStore) StoreLocally(alertr *pbgen.AlertDetail) error {
	as.take.Lock()
	defer as.take.Unlock()

	if as.store[alertr.Name] != nil {
		return AlertDuplicateFound
	} else {

	}

	return nil
}

func deepCopy(alertr *pbgen.AlertDetail) (*pbgen.AlertDetail, error) {
	newentry := &pbgen.AlertDetail{}

	err := copier.Copy(newentry, alertr)
	if err != nil {
		return nil, fmt.Errorf("Problems copying alert details %v", err)
	}
	return newentry, nil
}
