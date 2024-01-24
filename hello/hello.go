package main

import "fmt"

const (
	french     = "French"
	portuguese = "Portuguese"
	russian    = "Russian"

	englishPrefix    = "Hello, "
	portuguesePrefix = "Olá,"
	frenchPrefix     = "Bonjour, "
	russianPrefix    = "Privet, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case portuguese:
		prefix = portuguesePrefix
	case french:
		prefix = frenchPrefix
	case russian:
		prefix = russian
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Nicolas", portuguese))
}
