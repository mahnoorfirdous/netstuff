package main

import (
	"grsidecar/sample"
	"grsidecar/service"

	log "github.com/sirupsen/logrus"
)

//miscellaneous sanity checks

func main() {
	testserver := service.NewAlertCaterServer(&service.AlertStore{})
	sample.Names.Initstore(100)
	dummyreq := sample.NewAlert()

	testserver.AlertsMem.ReadyAlertStore()

	log.SetLevel(log.DebugLevel)
	err := testserver.AlertsMem.StoreAlert(dummyreq.GetAlertsbatch())
	if err != nil {
		log.Debugf("%v", err)
	} else {
		log.Debugf("Alert is Stored : %#v/n", testserver.AlertsMem.Store)
		log.Info("This line was printed without a problem")
		log.Info(log.GetLevel())
		log.Print(sample.Names)
	}

}
