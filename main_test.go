package main

import (
	"testing"

	models "github.com/admcpr/hub-bub/models"
)

func TestYesNo(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Yes", args{true}, "Yes"},
		{"No", args{false}, "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := models.YesNo(tt.args.b); got != tt.want {
				t.Errorf("YesNo() = %v, want %v", got, tt.want)
			}
		})
	}
}
