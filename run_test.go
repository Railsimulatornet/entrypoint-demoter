//go:build linux || darwin

package entrypoint_demoter

import (
	"io"
	"strings"
	"testing"
)

type collectingOutputListener struct {
	lines []string
}

func (l *collectingOutputListener) UseStdin(io.Writer) {}
func (l *collectingOutputListener) HandleStdout(line string) {
	l.lines = append(l.lines, line)
}

func TestRunCommandDrainsListenerOutput(t *testing.T) {
	for i := 0; i < 20; i++ {
		listener := &collectingOutputListener{}
		err := RunCommandWithListeners(0, 0, "", "", 0, []string{"sh", "-c", "printf 'first\\nlast'"}, listener)
		if err != nil {
			t.Fatal(err)
		}
		if len(listener.lines) != 2 || listener.lines[0] != "first" || listener.lines[1] != "last" {
			t.Fatalf("incomplete output: %q", listener.lines)
		}
	}
}

func TestRunCommandHandlesLongListenerLine(t *testing.T) {
	listener := &collectingOutputListener{}
	err := RunCommandWithListeners(0, 0, "", "", 0, []string{"sh", "-c", "head -c 131072 /dev/zero | tr '\\000' x"}, listener)
	if err != nil {
		t.Fatal(err)
	}
	if len(listener.lines) != 1 || listener.lines[0] != strings.Repeat("x", 131072) {
		t.Fatal("long output was not delivered in full")
	}
}
