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
	Store map[string]*pbgen.AlertDetail
}

var ErrAlertDuplicate = errors.New("alert is stored and will be serviced soon")

func (as *AlertStore) ReadyAlertStore() {
	*as = AlertStore{Store: make(map[string]*pbgen.AlertDetail)}
}

func (as *AlertStore) StoreAlert(alert *pbgen.AlertList) error {

	if alert == nil {
		return errors.New("alert is empty")
	}
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

	if alertr == nil {
		return errors.New("alert is empty")
	}

	if as.Store[alertr.Name] != nil {
		return ErrAlertDuplicate
	} else {
		var err error
		as.Store[alertr.Name], err = deepCopy(alertr)
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
