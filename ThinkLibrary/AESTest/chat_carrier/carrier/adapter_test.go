package carrier

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"reflect"
	"testing"
	"think-library/AESTest/chat_carrier/carrier/change"
	"think-library/AESTest/chat_carrier/carrier/source"
)

func Test_Metadata(t *testing.T) {

	src := &source.Items{}
	src.Ext = "default ext type string"

	changed := &change.Items{}
	changed.Append(change.Item{ID: 1001, Cur: 10, Change: 20})
	changed.Append(change.Item{ID: 1001, Cur: 15, Change: 25})
	changed.Append(change.Item{ID: 1001, Cur: 1, Change: 2})
	changed.Append(change.Item{ID: 1002, Cur: 10, Change: 20})
	changed.Append(change.Item{ID: 1002, Cur: 15, Change: 20})

	md := Metadata{}
	md.With(src, changed)

	log.Println("output metadata -> ", zap.Inline(md))

	amd := &Metadata{}
	if err := amd.UnmarshalJSON(Data()); err != nil {
		t.Error(err)
		return
	}

	for _, k := range amd.apt {
		fmt.Println("source -> ", reflect.TypeOf(k.Source), " === ", k.Source)
		fmt.Println("changed -> ", reflect.TypeOf(k.Change), " === ", k.Change)
		fmt.Println()
	}

	t.Log("Ok.")
}

func Data() []byte {
	return []byte(`[{"source": "{\"src\":1,\"ext\":\"default ext type string\"}", "changed": "{\"typ\":1,\"Data\":[{\"id\":1001,\"cur\":10,\"change\":20},{\"id\":1001,\"cur\":15,\"change\":25},{\"id\":1001,\"cur\":1,\"change\":2},{\"id\":1002,\"cur\":10,\"change\":20},{\"id\":1002,\"cur\":15,\"change\":20}]}"}]`)
}
