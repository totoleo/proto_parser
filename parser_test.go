package proto_parser

import (
	"math"
	"testing"

	"github.com/gogo/protobuf/proto"

	"github.com/totoleo/proto_parser/sample"
)

func TestDecode(t *testing.T) {
	request := &sample.RecallRequest{Dtu: 1, Uid: "2233", DeviceCode: "232222112222222", Extend: map[string]string{
		"hello": "world",
	}, TUid: "2233112"}

	buff, err := proto.Marshal(request)
	if err != nil {
		t.Error(err)
		return
	}
	fields, err := Decode(buff)
	if err != nil {
		t.Error(err)
		return
	}

	for i, field := range fields {
		t.Log(i, field, string(field.Data))
	}
}

func BenchmarkDecode(b *testing.B) {
	request := &sample.RecallRequest{Dtu: 1, Uid: "2233", DeviceCode: "232222112222222", Extend: map[string]string{
		"hello": "world",
	}, TUid: "2233112", Lat: 2.3, RegTime: math.MaxInt64}

	buff, err := proto.Marshal(request)
	if err != nil {
		b.Error(err)
		return
	}
	b.ReportAllocs()
	var fields []Field
	for i := 0; i < b.N; i++ {
		fields, err = Decode(buff)
		if err != nil {
			b.Error(err)
			return
		}
	}
	_ = fields
}

func BenchmarkUnmarshal(b *testing.B) {
	request := &sample.RecallRequest{Dtu: 1, Uid: "2233", DeviceCode: "232222112222222", Extend: map[string]string{
		"hello": "world",
	}, TUid: "2233112"}

	buff, err := proto.Marshal(request)
	if err != nil {
		b.Error(err)
		return
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		pb := &sample.RecallRequest{}
		_ = proto.Unmarshal(buff, pb)
	}
}
