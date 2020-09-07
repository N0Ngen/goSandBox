package main

import "fmt"

// Consts
const spanish string = "Spanish"
const french string = "French"
const arabic string = "Arabic"
const englishHelloPrefix string = "Hello, "
const spanishHelloPrefix string = "Hola, "
const frenchHelloPrefix string = "Bonjour, "
const arabicHelloPrefix string = "مرحبا, "

// one piece of code
// func main() {
// 	fmt.Println("Hello, World")
// }

// Hello separate the "domain" code from the outside world (side-effects)
func Hello(name, lang string) string {
	if name == "" {
		name = "World" // Return "Hello, World" as default if no argument passed
	}

	return greetingPrefix(lang) + name
}

// greetingprefix is a Private function
func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case arabic:
		prefix = arabicHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

// The main function
func main() {
	fmt.Println(Hello("Alix", ""))
}
