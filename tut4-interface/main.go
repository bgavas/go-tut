package main

// getGreeting'i olan her struct bot interface'ini de implement ediyor
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	// printGreetingEn(eb)
	// printGreetingSp(sb)
	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	println(b.getGreeting())
}

// func printGreetingEn(eb englishBot) {
// 	println(eb.getGreeting())
// }

// func printGreetingSp(sb spanishBot) {
// 	println(sb.getGreeting())
// }
