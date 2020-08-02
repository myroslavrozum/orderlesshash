package orderlesshash

//package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

/*
func TestGetIoSha(t *testing.T) {
	type args struct {
		r io.Reader
	}
	f, _ := os.Open(`./one.json`)
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{`one`,
			args{
				f,
			},
			[]byte{126, 94, 159, 89, 79, 75, 64, 45, 157, 207, 254, 9, 94, 34, 169, 45, 186, 103, 41, 5},
			false,
		},
	}

	for _, tt := range tests {
		var hashes Hashes
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashes.IoUnorderedSha(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIoSha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIoSha() = %v, want %v", got, tt.want)
			}
			t.Log(hashes)
		})
	}
}
*/

func TestHashes_JSONUnorderedSha(t *testing.T) {
	type args struct {
		json []byte
	}

	b, _ := ioutil.ReadFile(`./one.json`)

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{`one`,
			args{
				b,
			},
			[]byte{165, 206, 205, 165, 219, 88, 137, 15, 67, 22, 65, 183, 173, 107, 60, 152},
			false,
		},
	}

	for _, tt := range tests {
		var hashes Hashes
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashes.JSONMapUnorderedSha(tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hashes.JSONUnorderedSha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hashes.JSONUnorderedSha() = %v, want %v", got, tt.want)
			}
		})
	}
}
