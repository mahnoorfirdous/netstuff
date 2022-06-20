package service

import (
	"errors"
	"fmt"
	"grsidecar/pbgen"
	"sync"

	"github.com/jinzhu/copier"
)

type AlertStore struct {
	take  sync.RWMutex
	store map[string]*pbgen.AlertDetail
}

var ErrAlertDuplicate = errors.New("alert is stored and will be serviced soon")

func (as *AlertStore) ReadyAlertStore() {
	*as = AlertStore{store: make(map[string]*pbgen.AlertDetail)}
}

func (as *AlertStore) StoreAlert(alert *pbgen.AlertList) error {
	for _, indiv := range alert.Alerts {

		if err := as.StoreLocally(indiv); err != nil {
			if errors.Is(err, ErrAlertDuplicate) {
				continue //ignore the duplicate
			} else {
				return err
			}
		}
	}
	return nil
}

func (as *AlertStore) StoreLocally(alertr *pbgen.AlertDetail) error {
	as.take.Lock()
	defer as.take.Unlock()

	if as.store[alertr.Name] != nil {
		return ErrAlertDuplicate
	} else {
		var err error
		as.store[alertr.Name], err = deepCopy(alertr)
		if err != nil {
			return err
		}
	}
	return nil
}

func deepCopy(alertr *pbgen.AlertDetail) (*pbgen.AlertDetail, error) {
	newentry := &pbgen.AlertDetail{}

	err := copier.Copy(newentry, alertr)
	if err != nil {
		return nil, fmt.Errorf("problems copying alert details %v", err)
	}
	return newentry, nil
}
