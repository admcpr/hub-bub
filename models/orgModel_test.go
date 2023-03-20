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
