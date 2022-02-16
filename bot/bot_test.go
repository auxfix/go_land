package bot

import (
	"testing"
)

func TestRuEnBot(t *testing.T) {

	enb := EngBot{}
	rub := RuBot{}

	if TypeGreetings(enb) != "Hi man!" {
		t.Errorf("Expected bot greetings Hi man! but got %v", TypeGreetings(enb))
	}

	if TypeGreetings(rub) != "Привет!" {
		t.Errorf("Expected bot greetings ПРивет! but got %v", TypeGreetings(rub))
	}
}
