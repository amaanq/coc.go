package coc

import (
	"reflect"
	"testing"
)

func Test_getFromCache(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "hello", args: args{key: "hello"}, want: []byte("world"), wantErr: false},
		{name: "#2PP", args: args{key: "#2PP"}, want: []byte("#2PP"), wantErr: false},
		{name: "goodbye", args: args{key: "goodbye"}, want: []byte("night night"), wantErr: false},
	}

	// add 3 things first to cache
	writeToCache("hello", []byte("world"), 0)
	writeToCache("#2PP", []byte("#2PP"), 0)
	writeToCache("goodbye", []byte("night night"), 0)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFromCache(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFromCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFromCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeToCache(t *testing.T) {
	type args struct {
		key      string
		data     []byte
		duration int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "hello", args: args{key: "hello", data: []byte("world"), duration: 0}, wantErr: false},
		{name: "#2PP", args: args{key: "#2PP", data: []byte("#2PP"), duration: 0}, wantErr: false},
		{name: "goodbye", args: args{key: "goodbye", data: []byte("night night"), duration: 0}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeToCache(tt.args.key, tt.args.data, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("writeToCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
