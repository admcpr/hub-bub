package models

import (
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestOrgModel_Update(t *testing.T) {
	type args struct {
		msg tea.Msg
	}
	tests := []struct {
		name      string
		m         OrgModel
		args      args
		wantModel tea.Model
		wantCmd   tea.Cmd
	}{
		// TODO: Add more test cases.
		{"Quit KeyMsg", OrgModel{}, args{tea.KeyMsg{Type: tea.KeyCtrlC}}, OrgModel{}, tea.Quit},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotModel, gotCmd := tt.m.Update(tt.args.msg)
			if !reflect.DeepEqual(gotModel, tt.wantModel) {
				t.Errorf("OrganisationModel.Update() gotModel = %v, want %v", gotModel, tt.wantModel)
			}
			if reflect.ValueOf(gotCmd) != reflect.ValueOf(tt.wantCmd) {
				t.Errorf("OrganisationModel.Update() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}
		})
	}
}

func TestOrgModel_NewOrgModel(t *testing.T) {
	type args struct {
		title  string
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test NewOrgModel", args: args{title: "hub-bub", width: 1024, height: 768}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOrgModel(tt.args.title, tt.args.width, tt.args.height)
			if got.Title != tt.args.title {
				t.Errorf("NewOrgModel title got = %v, want %v", got.Title, tt.args.title)
			}
			if got.width != tt.args.width {
				t.Errorf("NewOrgModel width got = %v, want %v", got.width, tt.args.width)
			}
			if got.height != tt.args.height {
				t.Errorf("NewOrgModel height got = %v, want %v", got.height, tt.args.height)
			}
			if got.keys.Enter.Enabled() != true {
				t.Errorf("NewOrgModel keys.Enter got = %v, want %v", got.keys.Enter.Enabled(), true)
			}
		})
	}
}
