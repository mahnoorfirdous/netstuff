package sample

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Samplewords []string

func randomName(number int) []string {
	getURL := "https://random-word-api.herokuapp.com/word?number=" + fmt.Sprintf("%v", number)
	log.Info(getURL)
	recevd, err := http.Get(getURL)
	if err != nil {
		return nil
	}
	bodyrec, err := ioutil.ReadAll(recevd.Body)
	if err != nil {
		return nil
	}
	alertname := make([]string, number)
	err = json.Unmarshal(bodyrec, &alertname)
	if err != nil {
		return nil
	}
	return alertname
}

func (sw *Samplewords) Initstore(number int) error { //we don't really care about duplicates
	log.SetLevel(log.DebugLevel)
	if sw == nil {
		log.Debug("Could not initialize")
		return errors.New("could not init")
	}
	*sw = randomName(number)
	return nil
}

func (sw *Samplewords) Initialize() {
	*sw = []string{"chirography", "aplastic", "inbreathed", "clysters", "desensitize", "bathos", "coevolution", "japanned", "stairheads", "unchancy", "miswrite", "pes", "waggoners", "cranch", "irreligionist",
		"yesterdays", "unwed", "riffs", "sagittaries", "spleeny", "busker", "patients", "wardresses",
		"recriminate", "airsheds", "minitowers", "aviatrices", "noncoverage", "vigesimal", "epidemiology", "upkeeps", "heatstroke", "peon", "prates", "merciful", "crestfallen", "uncalcined", "bide", "gainsaid", "crowners", "leukocytosis",
		"obstinately", "authenticates", "roamed", "copay", "bairnish", "cocurricular", "paganized", "grease", "vizcachas", "evicts", "snugnesses", "homotransplants", "disbowelling",
		"rechanging", "misdescribed", "verbatim", "scuffs", "diriment", "trophic", "paulin", "dopa", "myoscope", "endplate", "thunderers", "fluctuant", "formalness", "freethinkings", "obesity", "frumpily", "voraciously",
		"coeds", "nosedove", "jitterbug", "legmen", "misanalysis", "oreides", "semeiologies", "begrudgingly", "intersessions", "confiscates", "afrit", "sanghs", "freakiness", "mycobacterial", "benefic",
		"analyzable", "misidentified", "roominesses", "relook", "kilohertzes", "decompressing", "mesotrophic", "miasms", "lowercase", "atonal", "calderas", "alcoholically", "crematorium", "isocyclic"}
}

func (sw *Samplewords) getrandomName() string {
	if sw == nil {
		return "not initialized"
	}
	return (*sw)[rand.Intn(len(*sw))]
}
