package main

import (
	"fmt"
)

type Animal struct {
	name   string
	speech string
}

func speak(anim Animal, ch chan string) {
	// Send a string message to the channel
	ch <- fmt.Sprintf("%s says '%s'.", anim.name, anim.speech)
}

func main() {
	dog := Animal{name: "Dog", speech: "Woof"}
	cat := Animal{name: "Cat", speech: "Meow"}

	// Create the channel which will accept string messages
	animalSpeechChannel := make(chan string)

	// We call speak() twice, passing the channel
	go speak(dog, animalSpeechChannel)
	go speak(cat, animalSpeechChannel)

	// We block until we receive two messages via the channel
	a1, a2 := <-animalSpeechChannel, <-animalSpeechChannel

	fmt.Println(a1, a2)
}
