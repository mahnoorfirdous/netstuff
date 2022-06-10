package main

import (
	"log"
	"os"
	"prototry/samplepb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	stream := samplepb.AlertRequest{
		URLS:       []string{"https://www.google.com.pk", "https://www.msn.com"},
		Updateflag: true,
		Stallcap:   samplepb.URGENCY_IMMEDIATE,
		AlertTime:  timestamppb.Now(),
	}
	var err error
	stream.Callinghost, err = os.Hostname()
	if err != nil {
		log.Printf("Could not get hostname %v/n", err)
	}

	log.Printf("Result of direct access %v", stream.URLS)
	log.Println("Result of getter method ", stream.GetURLS())

	bytedatas, err := proto.Marshal(&stream)
	if err != nil {
		log.Default().Fatalln("Could not marshal: ", err)
	}
	log.Default().Println("The bytedata as string: ", string(bytedatas))

	recycle := &samplepb.AlertRequest{}
	err = proto.Unmarshal(bytedatas, recycle)
	if err != nil {
		log.Default().Fatalln("Could not umarshal: ", err)
	}
	log.Printf("Bytedatas unmarshaled %v", recycle)

	cloudflare_alert := samplepb.AlertRequest{
		URLS:       []string{"https://www.cloudflare.com", "https://1.1.1.1"},
		Updateflag: true,
		Stallcap:   samplepb.URGENCY_IMPORTANT,
		AlertTime:  timestamppb.Now(),
	}

	cloudflare_alert.Callinghost, err = os.Hostname()
	if err != nil {
		log.Printf("Could not get hostname %v/n", err)
	}
	/* Full Alert */
	fullstream := samplepb.Alert{Request: []*samplepb.AlertRequest{&stream, recycle, &cloudflare_alert}}
	//	[]*samplepb.AlertRequest: request is a slice of pointers to AlertRequest objects
	// we need addresses of AlertRequest objects here

	bytedatas, err = proto.Marshal(&fullstream)
	if err != nil {
		log.Default().Fatalln("Could not marshal: ", err)
	}
	log.Default().Println("The bytedata as string: ", string(bytedatas))

	/* and we get eject it and unmarshal it similar to the above */
	eject := samplepb.Alert{}
	err = proto.Unmarshal(bytedatas, &eject)
	if err != nil {
		log.Default().Fatalln("Could not umarshal: ", err)
	}
	log.Printf("Bytedatas unmarshaled %v", eject.Request)

}
