package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricRecordforDeletion struct {
}

const alphas string = "abcdefghijklmnopqrstuvwxyz"

func RandomWords(length int) string {
	buff := strings.Builder{}
	alphlen := len(alphas)
	for i := 0; i < length; i++ {
		c := alphas[rand.Intn(alphlen)]
		buff.WriteByte(c)
	}
	return buff.String()
}

var ddos_attack_labels = []string{"container_label_io_kubernetes_pod_name", "attackType", "ruleName",
	"direction", "ruleProtocol", "ruleAction", "ipVer", "source",
	"destination", "destPort", "mitigated"}

func main() {

	thakkireg := prometheus.NewRegistry()

	ddos_attack_detected := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ddos_attack_detected",
			Help: "Active DDoS attack.",
		}, ddos_attack_labels)

	thakkireg.MustRegister(ddos_attack_detected)

	http.Handle("/metrics", promhttp.HandlerFor(thakkireg, promhttp.HandlerOpts{Registry: thakkireg}))

	go func() {
		var DdosAttackDetected float64 = rand.Float64()
		for {
			DdosLabels := prometheus.Labels{
				"container_label_io_kubernetes_pod_name": RandomWords(10),
				"attackType":                             []string{"cat", "dog"}[rand.Intn(2)],
				"ruleName":                               "ddos",
				"direction":                              "ingress",
				"ruleProtocol":                           []string{"cat", "dog"}[rand.Intn(2)],
				"ipVer":                                  []string{"IPV6", "IPV4"}[rand.Intn(2)],
				"source":                                 []string{"unknown", "your basement"}[rand.Intn(2)],
				"destination":                            RandomWords(10),
				"destPort":                               strconv.Itoa(rand.Intn(65536)),
				"ruleAction":                             []string{"slap", "hug"}[rand.Intn(2)],
				"mitigated":                              "0",
			}

			ddos_attack_detected.With(DdosLabels).Set(float64(DdosAttackDetected))
			time.Sleep(8 * time.Second)
		}
	}()

	go func() {
		/* Cleanup */
		for {
			time.Sleep(20 * time.Second)
			log.Println("Deleting metric vector")
			ddos_attack_detected.MetricVec.Reset()
		}
	}()

	http.ListenAndServe(":8030", nil)

}
