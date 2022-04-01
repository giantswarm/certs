package certstest

import (
	"testing"

	"github.com/giantswarm/certs/v4/pkg/certs"
)

func Test_CertsTest_NewSearcher(t *testing.T) {
	config := Config{}
	s := NewSearcher(config)
	_, ok := interface{}(s).(certs.Interface)
	if !ok {
		t.Fatal("searcher does not implement correct interface")
	}
}
