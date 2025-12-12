package commands

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCommandListInitialization(t *testing.T) {
	t.Helper()

	expectedCommands := []string{"say", "/exit", "/help"}

	for _, name := range expectedCommands {
		cmd, ok := CommandList[name]
		if !ok {
			t.Fatalf("command %s not found in CommandList", name)
		}

		if cmd.Name != name {
			t.Errorf("command %s has name %s", name, cmd.Name)
		}

		if cmd.Execute == nil {
			t.Errorf("command %s has nil Execute function", name)
		}
	}
}

func TestSayCommand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ret, err := Say.Execute(ParsedCommand{Args: []string{"hello", "world"}})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if ret.Control != ControlNone {
			t.Errorf("unexpected control: %s", ret.Control)
		}

		if ret.Message != "hello world" {
			t.Errorf("unexpected message: %s", ret.Message)
		}
	})

	t.Run("error when missing args", func(t *testing.T) {
		ret, err := Say.Execute(ParsedCommand{})

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if ret.Control != ControlNone {
			t.Errorf("expected ControlNone, got %s", ret.Control)
		}

		if ret.Message != "" {
			t.Errorf("expected empty message, got %s", ret.Message)
		}
	})
}

func TestExitCommand(t *testing.T) {
	ret, err := Exit.Execute(ParsedCommand{})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ret.Control != ControlExit {
		t.Errorf("expected ControlExit, got %s", ret.Control)
	}

	if ret.Message != "" {
		t.Errorf("expected empty message, got %s", ret.Message)
	}
}

func TestHelpCommandOutputsCommands(t *testing.T) {
	output := captureOutput(t, func() {
		ret, err := Help.Execute(ParsedCommand{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if ret != (CommandReturn{}) {
			t.Fatalf("expected zero CommandReturn, got %+v", ret)
		}
	})

	for _, expected := range []string{"Available Commands:", "say:", "/exit:", "/help:"} {
		if !strings.Contains(output, expected) {
			t.Errorf("help output missing %q in %q", expected, output)
		}
	}
}

func captureOutput(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("creating pipe: %v", err)
	}

	os.Stdout = w
	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("closing writer: %v", err)
	}
	os.Stdout = oldStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("reading pipe: %v", err)
	}
	if err := r.Close(); err != nil {
		t.Fatalf("closing reader: %v", err)
	}

	return buf.String()
}
