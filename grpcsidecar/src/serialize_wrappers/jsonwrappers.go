package serialize_wrappers

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto" //the old package used by protobuf-compiler starts with github.com/
	//maybe we won't need to do anything if we just use a newer version of protobuf
)

func ProtobuftoJson(msg proto.Message) string {
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers:  false,
		Indent:          " ",
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}

	return marshaler.Format(msg)
}

func JsontoProtobuf[Raw []byte | string](data Raw, msg proto.Message) (proto.Message, error) {
	rawdata := []byte(data)
	err := protojson.Unmarshal(rawdata, msg)
	return msg, err
}
