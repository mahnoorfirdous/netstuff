package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	thakkireg := prometheus.NewRegistry()

	temporary := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "room_temperature_fahrenheit",
		Help: "This will not actually be room temp but dummy value :P",
	})

	newtry := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "bool_test_val",
		Help: "This will let  me check what happens to bool vals",
	}, []string{"subindicator_1", "subindicator_2"})

	thakkireg.MustRegister(temporary)
	thakkireg.MustRegister(newtry)

	temporary.Set(141.423)
	newtry.WithLabelValues("value1", "value2").Set(0.0)
	//newtry.WithLabelValues("value1", "value2").SetToCurrentTime()

	http.Handle("/metrics", promhttp.HandlerFor(thakkireg, promhttp.HandlerOpts{Registry: thakkireg}))

	go func() {
		var i int
		for true {
			time.Sleep(6 * time.Second) //let me open the page
			i = i + 1
			val := float64(173.33) + float64(i)
			temporary.Set(val)
			newtry.WithLabelValues("value1", "value2").Set(1.0)
			time.Sleep(65 * time.Second)
			newtry.WithLabelValues("value1", "value2").Set(0.0)
			time.Sleep(65 * time.Second)
			newtry.MetricVec.Reset()
			time.Sleep(65 * time.Second)
			newtry.WithLabelValues("none1", "none2").Set(0.0)

		}
	}()

	http.ListenAndServe(":8030", nil)

}
