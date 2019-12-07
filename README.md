#GoLang Basics


-	Go is a compiled statically typed language with a C - like syntax and garbage collection.


##Building and running go

-	go build filename.go
-	./filename


##Running with go run
-	It builds and run the code straigt from source file.
-	go run filename.go

##The gofmt command

-	The gofmt command ships with go and can be used to automatically format source code.

##Formatting source code with gofmt
-	gofmt could be configured to run on every file save or before each commit
-	command: gofmt -w filename.go


##Reading Arguments:

-	The os.Args array contains all arguments passed to the running program, including user-supplied arguments from the command line.
-	os.Args - It's an array containing program arguments, starting with the name of the executables and followed by any user-supplied arguments.


##The goimports command
-	It ships with go and detects missing packages and automatically updates import statements in the source code.
-	command: goimports -w filename.go

##Declaring Variables with Type Inferences

-	The := operator tells Go to automatically find out the data type on the right being assigned to the newly declared variables on the left. This is known as type inference.


##Storing Values as Data Types

-	For every value, there's always a proper data type that determines how the value should be stored and the operations that can be done on values of that data type

##Type Inference vs Manual Type Declaration

##Type Inference:
-	Data type is infered during assignment.
	Ex: message := "Hello World"

##Manual Type Declaration:
-	Data type is declared prior to assignment
	Ex: var message string
		message = "Hello World"

##Built-in data types in GO
-	int
-	string
-	bool
-	[]string


##Functions:

##Declaration:

func getGreeting(hour int) string {}
	
	-	getGreeting - function name
	-	hour	-	paramter
	-	string 	-	return type

##Working with Errors

##Declararing Multiple Return Values:

-	In Go, we communicate errors via a separate  return value. Lets update our function signature to return two different values: a string and an error.

##Zero Values:

-	A zero value in Go is the default value assigned to variables declared without an explicit initial value.
	-	string - ""
	-	initial - 0
	-	bool - false
	-	float - 0.0
	-	byte - 0
	-	function - nil

##Loops
- use for keyword for all type of loops

##Arrays:

 -	Arrays are of fixed size
 -	Ex: var words [2]string
		words[0] = "Hello"
		words[1] = "How"

##Slices
 -	Slices are dynamic
 -	A nil Slice in Go behaves same as an empty slice. It can be appended to using the built-in function append, which takes two arguments: a slice and a variable number of elements
 -	Ex: var words []string
		words = append(words, "Hey")
		words = append(words, "Yo")

##Slice Literals

-	A slice literal is a quick way to create slices with initial elements via type inference. We can pass elements between curly braces {}.
-	Ex: words := []string {"Hey", "Yo", "Whatsup", "Dude"}
		fmt.Println(words)

##Reading from a Slice with unknown size

-	While reading element by index works, it doesn't scale well once our slice grows or when the exact number of elements is not known until the program is run.
-	Ex:
	words := getWords()
	for i, elem := range words  {
		fmt.Println("index", i, "Value", elem)
	}

##Navigating a slice with for and range

-	The range clause provides a way to iterate over slices. When only one value is used on the left of the range, then this value will be the index for each element on the slice, one at a time.

##Structs:
-	A Struct is a built in data type that allows us to group properties under a single name.

-	Ex: 
	type person struct {
		name string
		age int
	}

##Pass By Value

A> Pass By Value (default behaviour)

-	language := "Scala"  [memory address - 0X123456]
-	jvmLanguage := language  [memory address - 0X654321] (Copy of original data)

##Pass By Reference

B> Pass by reference

-	language := "Scala"  [memory address - 0X123456]
-	jvmLanguage := &language  [memory address - 0X123456] (reference of original data)

##Interfaces:

-	type playcser interface {
		playcs() string
	}

##goroutines

-	A goroutine is a special function that executes concurrently with other functions. We create them with go keyword.

-	Ex 1:
	func1()
	func2()
-	Here, func2 must wait until func1 is finished

-	Ex 2:
	go func1()	- creates goroutines
	func2()
-	Here, func1 and func2 are executed concurrently

Note: In single core machine, concurrent core is unlikely to perform better than sequential code.


















