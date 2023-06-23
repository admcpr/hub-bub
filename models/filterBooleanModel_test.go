package models

import (
	"hub-bub/structs"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewFilterBooleanModel(t *testing.T) {
	tests := []struct {
		title string
		value bool
	}{
		{"Test 1", true},
		{"Test 2", false},
		{"Test 3", true},
		{"Test 4", false},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			m := NewFilterBooleanModel(tt.title, tt.value)

			if m.Title != tt.title {
				t.Errorf("got %q, want %q", m.Title, tt.title)
			}

			if m.input.Value() != structs.YesNo(tt.value) {
				t.Errorf("got %q, want %q", m.input.Value(), structs.YesNo(tt.value))
			}
		})
	}
}

func TestFilterBooleanModel_Update(t *testing.T) {
	trueModel := NewFilterBooleanModel("True", true)
	falseModel := NewFilterBooleanModel("False", false)

	tests := []struct {
		model  FilterBooleanModel
		title  string
		msgKey rune
		want   bool
	}{
		{title: "'n' should set value to false", model: trueModel, msgKey: 'n', want: false},
		{title: "'N' should set value to false", model: trueModel, msgKey: 'N', want: false},
		{title: "'y' should set value to true", model: falseModel, msgKey: 'y', want: true},
		{title: "'Y' should set value to true", model: falseModel, msgKey: 'Y', want: true},
		{title: "'x' shouldn't change false value", model: falseModel, msgKey: 'x', want: false},
		{title: "'x' shouldn't change true value", model: trueModel, msgKey: 'x', want: true},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{tt.msgKey}}
			m, _ := tt.model.Update(msg)

			filterBooleanModel, _ := m.(FilterBooleanModel)
			got := filterBooleanModel.GetValue()

			if got != tt.want {
				t.Errorf("FilterBooleanModel.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
