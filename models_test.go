package main

import (
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestUserModel_Update(t *testing.T) {
	testUser := User{Login: "test"}
	type args struct {
		msg tea.Msg
	}
	tests := []struct {
		name      string
		m         UserModel
		args      args
		wantModel tea.Model
		wantCmd   tea.Cmd
	}{
		// TODO: Add test cases.
		{"Authentication Success", UserModel{}, args{AuthenticationMsg{User: testUser}}, UserModel{Authenticated: true, User: testUser}, getOrganisations},
		{"Authentication Failure", UserModel{}, args{AuthenticationErrorMsg{}}, UserModel{Authenticated: false}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotModel, gotCmd := tt.m.Update(tt.args.msg)
			if !reflect.DeepEqual(gotModel, tt.wantModel) {
				t.Errorf("UserModel.Update() gotModel = %v, want %v", gotModel, tt.wantModel)
			}
			if reflect.ValueOf(gotCmd) != reflect.ValueOf(tt.wantCmd) {
				t.Errorf("UserModel.Update() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}
		})
	}
}

func TestOrganisationModel_Update(t *testing.T) {
	type args struct {
		msg tea.Msg
	}
	tests := []struct {
		name      string
		m         OrganisationModel
		args      args
		wantModel tea.Model
		wantCmd   tea.Cmd
	}{
		// TODO: Add more test cases.
		{"Quit KeyMsg", OrganisationModel{}, args{tea.KeyMsg{Type: tea.KeyCtrlC}}, OrganisationModel{}, tea.Quit},
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
