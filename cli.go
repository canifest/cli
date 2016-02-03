package main

import (
	"github.com/abiosoft/ishell"
)

func main() {
	shell := ishell.New()
	shell.Println("--[ canifest ]-----")
	shell.SetPrompt(">")

	shell.Register("hello", func(args ...string) (string, error) {
		shell.ShowPrompt(false)
		defer shell.ShowPrompt(true)  // interesting use of defer for reverting state
		
		shell.Print("What is your name? ")
		name := shell.ReadLine()

		return "Hello " + name, nil
	})

	shell.Start()
}