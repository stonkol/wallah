package cli

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func Test_settingsMenuModel_View(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want tea.View
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var m settingsMenuModel
			got := m.View()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("View() = %v, want %v", got, tt.want)
			}
		})
	}
}
