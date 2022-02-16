package bot

import "fmt"

type EngBot struct {}
type RuBot struct {}

type Bot interface {
	 GetGreetings() string
}

func TypeGreetings(b Bot) string {
	fmt.Println(b.GetGreetings())
	return b.GetGreetings()
}


func(EngBot) GetGreetings() string {
	return "Hi man!"
}

func(RuBot) GetGreetings() string {
	return "Привет!"
}
