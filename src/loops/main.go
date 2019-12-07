package main

import "fmt"

func main() {
	fmt.Println("For Loop")
	forLoopExample()
	fmt.Println("While Loop")
	whileLoopExample()
	fmt.Println("While Loop 2")
	whileLoopExample2()
}

func forLoopExample()  {
	for i := 0; i< 5; i++ {
		fmt.Println(i)
	}
}

func whileLoopExample()  {
	i := 0
	isLessThanFive := true

	for isLessThanFive {
		if i >= 5 {
			isLessThanFive = false
		}
		fmt.Println(i)
		i++
	}
}

func whileLoopExample2()  {
	i := 0

	for {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}

/*func inifiniteLoop()  {

	for {
		//Some function listening for connections from other programs
		someListeningFunction()
	}
}*/