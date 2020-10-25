package mock

import "errors"

// fake data
var urls map[string]string = map[string]string{
	"Url 1": `The book you are reading right now is Mastering Go Second Edition, and is all about helping
	you become a better Go developer! There are many exciting new topics, including an entirely new chapter that talks about
	Machine Learning in Go as well as information and code examples relating to the Viper and Cobra Go packages, gRPC, working 
	with Docker images, working with YAML files, working with the go/scannerand go/tokenpackages, and generating WebAssembly
	code from Go. In total, there are more than 130 new pages in this second edition of Mastering Go`,

	"Url 2": `This book is for amateur and intermediate Go programmers who want to take their Go knowledge to the next level, 
	as well as experienced developers in other programming languages who want to learn Go without learning again how a forloop works.
	Some of the information found in this book can be also found in my other book, Go System Programming. The main difference 
	between these two books is that Go Systems Programmimg is about developing system tools using the capabilities of Go, whereas 
	Mastering Gois explaining the capabilities and the internals of Go in order to become a better Go develop Both books can be.`,

	"Url 3": `Chapter 1, Go and the Operating System, begins by talking about the history of Go and the advantages of Go before 
	describing the godocutility and explaining how you can compile and execute Go programs. After that, it talks about printing 
	output and getting user input, working with the command-line arguments of a program and using log files. The final topic
	in the first chapter is error handling, which plays a key role in Go. Chapter 2, Understanding Go Internals, discusses the 
	Go garbage collector and the way it	operates. Then, it talks about unsafe code and the unsafepackage.`,

	"Url 4": `Chapter 5, How to Enhance Go Code with Data Structures, is about developing your own data structures when the structures 
	offered by Go do not fit a particular problem. This includes developing binary trees, linked lists, hash tables, stacks, and queues, 
	and learning about their advantages. This chapter also showcases the use of the structures found in the	containerstandard Go package, 
	as well as how to use Go to verify Sudoku puzzles and generate random numbers. Chapter 7, Reflection and Interfaces for All Seasons, 
	discusses three advanced Go concepts: reflection, interfaces, and type methods. Additionally, it discusses the object-oriented
	capabilities of Go and how to debug Go programs using Delve!`,
}

// TotalGoAppearingNumber is a total number of "Go" appearing in the mock URLs responses
const TotalGoAppearingNumber = 29

// RequestURL mocks HTTP GET request to the url
func RequestURL(url string) (string, error) {
	body, ok := urls[url]
	if ok == false {
		return "", errors.New("Bad request")
	}

	return body, nil
}
