package main

import "fmt"

func main() {

	testArray()
	fmt.Println("--------------------")
	testSlices()
	fmt.Println("--------------------")
	testSlicesLiterals()
	fmt.Println("--------------------")
	navigatingSlices()
	fmt.Println("--------------------")
	navigatingSlices2()
	fmt.Println("--------------------")
	navigatingSlices3()
}

func testArray() {
	var words [5]string
	fmt.Println("Testing Array")
	words[0] = "Hello"
	words[1] = "How"
	words[2] = "Are"
	words[3] = "You"
	words[4] = "Clayman"
	fmt.Println(words)
}

func testSlices() {
	var words []string
	fmt.Println("Testing Slices")
	words = append(words, "Hey")
	words = append(words, "Yo")
	words = append(words, "Clayman")
	words = append(words, "Whatsup")
	words = append(words, "!")
	fmt.Println(words)
}

func testSlicesLiterals() {
	words := []string {"Hey", "Yo", "Whatsup", "Dude"}
	fmt.Println("Testing Slice literals")
	fmt.Println(words)
	fmt.Println(words[0])
	fmt.Println(words[1])
	fmt.Println(words[2])
	fmt.Println(words[3])
}

func navigatingSlices() {
	fmt.Println("Navigating Slices")
	words := getWords()
	for i := range words  {
		fmt.Println(words[i])
	}
}

func navigatingSlices2() {
	fmt.Println("Navigating Slices")
	words := getWords()
	for i, elem := range words  {
		fmt.Println("index", i, "Value", elem)
	}
}

func navigatingSlices3() {
	fmt.Println("Navigating Slices")
	words := getWords()
	for _, elem := range words  {
		fmt.Println("Value", elem)
	}
}

func getWords() []string {

	return []string{"Hello", "How", "are", "you", "Clayman"}
}