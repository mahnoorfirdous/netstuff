package sample

import (
	"fmt"
	"grsidecar/pbgen"
	"math/rand"
	"os"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var URLlist []string = []string{
	"https://www.example.com/?apparel=board",
	"https://www.example.net/",
	"https://example.org/",
	"https://example.com/base/ba",
	"http://example.com/aunt",
	"https://www.example.com/aunt",
}

var K8sdistro []string = []string{
	"MicroK8s",
	"K3s", "K3d",
	"Minikube", "EKS",
	"GKE", "AKS", "Openshift",
	"Rafay", "RKE",
}

var contexts []string = []string{
	"random-context", "mycluster204", "cocoa_943", "minikube", "kubeadm",
}

func randomUpdate() bool {
	return rand.Intn(2) == 1
}

func randomURL() []string {
	length := len(URLlist)
	end := rand.Intn(length)
	begin := rand.Intn(length)
	if end > begin {
		return URLlist[begin:end]
	} else {
		return URLlist[end:begin]
	}

	return nil
}

var m map[string]int64 = map[string]int64{}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func getHostname() (hostname string) {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	} else {
		return hostname
	}
}

func randomAlert() *pbgen.AlertDetail {
	return &pbgen.AlertDetail{
		Updateflag:  randomUpdate(),
		Alerttime:   timestamppb.Now(),
		URLS:        randomURL(),
		Callinghost: getHostname(),
		K8Sorigin:   randomK8sinstance(),
	}
}

func randomK8sdistro() string {
	return K8sdistro[rand.Intn(len(K8sdistro))]
}

func randomK8scontext() string {
	return contexts[rand.Intn(len(contexts))]
}
func randomK8sinstance() *pbgen.K8SInstance {
	return &pbgen.K8SInstance{
		K8Sver:    fmt.Sprintf("%v", rand.Float32()*22),
		K8Sdistro: randomK8sdistro(),
		Context:   randomK8scontext(),
	}
}
