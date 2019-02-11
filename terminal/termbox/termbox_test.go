package termbox

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/mum4k/termdash/terminalapi"
)

func TestNewTerminal(t *testing.T) {
	tests := []struct {
		desc string
		opts []Option
		want *Terminal
	}{
		{
			desc: "default options",
			want: &Terminal{
				colorMode: terminalapi.ColorMode256,
			},
		},
		{
			desc: "sets color mode",
			opts: []Option{
				ColorMode(terminalapi.ColorModeNormal),
			},
			want: &Terminal{
				colorMode: terminalapi.ColorModeNormal,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got := newTerminal(tc.opts...)

			// Ignore these fields.
			got.events = nil
			got.done = nil

			if diff := pretty.Compare(tc.want, got); diff != "" {
				t.Errorf("newTerminal => unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
