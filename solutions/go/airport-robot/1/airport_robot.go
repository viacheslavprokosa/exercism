package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}
type German struct {
}

func (gg German) LanguageName() string {
	return "German"
}
func (gg German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}

type Portuguese struct {
}

func (pg Portuguese) LanguageName() string {
	return "Portuguese"
}
func (pg Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

type Italian struct {
}

func (ig Italian) LanguageName() string {
	return "Italian"
}
func (ig Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}
func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}