package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFilterDateModel(t *testing.T) {
	const title = "Title"
	fromString := "2022-01-01"
	toString := "2022-12-31"
	from, _ := time.Parse("2006-01-02", fromString)
	to, _ := time.Parse("2006-01-02", toString)

	t.Run("NewFilterDateModel", func(t *testing.T) {
		m := NewFilterDateModel(title, from, to)
		assert.Equal(t, m.Title, title)
		assert.Equal(t, m.fromInput.Placeholder, fromString)
		assert.Equal(t, m.toInput.Placeholder, toString)

		gotFrom, gotTo := m.GetValue()

		assert.Equal(t, gotFrom, from)
		assert.Equal(t, gotTo, to)
	})
}

// func TestFilterDateModel_View(t *testing.T) {
// 	// 	type args struct {
// 	// 		title string
// 	// 		from  time.Time
// 	// 		to    time.Time
// 	// 	}
// 	// 	tests := []struct {
// 	// 		name string
// 	// 		args args
// 	// 		want FilterDateModel
// 	// 	}{
// 	// 		{name: "1", args: {"Title", time.Now(), time.Now() }, FilterDateModel{title: "Title", from: time.Now(), to: time.Now()}}
// 	// 	}
// 	// 	for _, tt := range tests {
// 	// 		t.Run(tt.name, func(t *testing.T) {
// 	// 			if got := NewFilterDateModel(tt.args.title, tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
// 	// 				t.Errorf("NewFilterDateModel() = %v, want %v", got, tt.want)
// 	// 			}
// 	// 		})
// 	// 	}

// 	t.Run("View", func(t *testing.T) {
// 		m := NewFilterDateModel(title, from, to)
// 		m.View()
// 	})

// 	t.Run("Focus", func(t *testing.T) {
// 		m := NewFilterDateModel(title, from, to)
// 		m.Focus()
// 	})
// }
