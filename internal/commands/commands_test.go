package commands

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/MichaelRBond/go-db/internal/locations"
	"github.com/MichaelRBond/go-db/internal/player"
)

func TestCommandListInitialization(t *testing.T) {
	t.Helper()

	expectedCommands := []string{"look", "move", "say", "/exit", "/help"}

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
		ret, err := Say.Execute(player.InitPlayer(), buildRooms(), ParsedCommand{Args: []string{"hello", "world"}})

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
		ret, err := Say.Execute(player.InitPlayer(), buildRooms(), ParsedCommand{})

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
	ret, err := Exit.Execute(player.InitPlayer(), buildRooms(), ParsedCommand{})

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
		ret, err := Help.Execute(player.InitPlayer(), buildRooms(), ParsedCommand{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if ret != (CommandReturn{}) {
			t.Fatalf("expected zero CommandReturn, got %+v", ret)
		}
	})

	for _, expected := range []string{"Available Commands:", "look:", "move:", "say:", "/exit:", "/help:"} {
		if !strings.Contains(output, expected) {
			t.Errorf("help output missing %q in %q", expected, output)
		}
	}
}

func TestLookCommand(t *testing.T) {
	p := player.InitPlayer()
	rooms := buildRooms()

	ret, err := Look.Execute(p, rooms, ParsedCommand{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(ret.Message, "Room One") {
		t.Fatalf("expected room name in output, got %q", ret.Message)
	}

	plainMessage := stripANSI(ret.Message)
	if !strings.Contains(plainMessage, "Exits: [north]") {
		t.Fatalf("expected exits list in output, got %q", plainMessage)
	}
}

func TestMoveCommand(t *testing.T) {
	t.Run("success moves player and describes new room", func(t *testing.T) {
		p := player.InitPlayer()
		rooms := buildRooms()

		ret, err := Move.Execute(p, rooms, ParsedCommand{Args: []string{"north"}})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if p.Location != "room_two" {
			t.Fatalf("expected player to move to room_two, got %s", p.Location)
		}

		if !strings.Contains(ret.Message, "You move north.") {
			t.Errorf("missing move acknowledgement in %q", ret.Message)
		}

		if !strings.Contains(ret.Message, "Room Two") {
			t.Errorf("missing new room description in %q", ret.Message)
		}
	})

	t.Run("error when invalid direction", func(t *testing.T) {
		p := player.InitPlayer()
		rooms := buildRooms()

		_, err := Move.Execute(p, rooms, ParsedCommand{Args: []string{"west"}})
		if err == nil {
			t.Fatalf("expected error for invalid direction")
		}

		if p.Location != "grandmas_kitchen" {
			t.Fatalf("player location should not change on error, got %s", p.Location)
		}
	})

	t.Run("error when current location missing", func(t *testing.T) {
		p := player.InitPlayer()
		p.Location = "unknown-room"
		rooms := buildRooms()

		_, err := Move.Execute(p, rooms, ParsedCommand{Args: []string{"north"}})
		if err == nil {
			t.Fatalf("expected error for missing location")
		}
	})
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

var ansiSequence = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(input string) string {
	return ansiSequence.ReplaceAllString(input, "")
}

func buildRooms() *locations.RoomsById {
	roomOne := locations.Room{
		Id:          "grandmas_kitchen",
		Name:        "Room One",
		Description: "First room",
		Exits: []locations.RoomExit{
			{
				Direction: locations.North,
				RoomID:    "room_two",
			},
		},
	}

	roomTwo := locations.Room{
		Id:          "room_two",
		Name:        "Room Two",
		Description: "Second room",
		Exits: []locations.RoomExit{
			{
				Direction: locations.South,
				RoomID:    "grandmas_kitchen",
			},
		},
	}

	rooms := locations.RoomsById{
		roomOne.Id: &roomOne,
		roomTwo.Id: &roomTwo,
	}

	return &rooms
}
