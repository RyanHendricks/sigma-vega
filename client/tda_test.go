package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFetchOptionChain(t *testing.T) {
	tests := []struct {
		name    string
		want    *OptionChain
		wantErr bool
	}{
		{"TestFetchOptionChain", &OptionChain{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchOptionChain()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchOptionChain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("FetchOptionChain() got = %v, want %v", got, tt.want)
			}
			fmt.Println(len(got.PutMap), len(got.CallMap))
			PrettyPrint(got)
		})
	}
}
