package main

import (
	"net/http"
	"os"

	"github.com/abiosoft/ishell"
)

//TODO someone explain the different between localhost:993 and :993 to me, please
var defaultURL = "http://localhost:9993"
var shell *ishell.Shell

func main() {
	createNewShell()

	writeIntroductoryStuff()

	verifyServerStatusAndInformUser()

	configureShell()
	registerHandlers()

	startShell()
}

func createNewShell() {
	shell = ishell.New()
}

func writeIntroductoryStuff() {
	//TODO make this more interesting
	shell.Println("--[ canifest ]-----")
}

func verifyServerStatusAndInformUser() {
	showCheckServerStatus()
	var serverError error
	if serverError = checkRestServerStatus(); serverError != nil {
		//TODO make this more graceful or a better experience for the user
		shell.Println(serverError.Error())
		exitApplication(1) //TODO - non-zero might be useful for the wrapper script we need to write
	}
	showServerSuccessfulConnection()
}

func configureShell() {
	shell.SetPrompt(">")
}

func registerHandlers() {
	shell.Register("help", help)
	shell.Register("quit", quit)
	shell.Register("exit", quit)
}

func showCheckServerStatus() {
	shell.Println("checking server status... (" + defaultURL + ")")
}

func checkRestServerStatus() error {
	var response, error = http.Get(defaultURL + "/status")
	if error != nil {
		return error
	}
	defer response.Body.Close()
	return nil
}

func showServerSuccessfulConnection() {
	shell.Println("ok")
}

func help(args ...string) (string, error) {
	//TODO either make help text read from a file or from the rest server
	//TODO go to the server and get the list of commands to show the user
	var helpText string
	helpText += "Welcome to Canifest. To make this work properly, make sure you "
	helpText += "start the core rest server before you run the CLI.\n"
	helpText += "./bin/core from your GOPATH"
	return helpText, nil
}

func quit(args ...string) (string, error) {
	var _, _ = http.Get(defaultURL + "/quit")
	exitApplication(0)
	return "", nil //TODO golang says I can't return nil as a string, what else could I do instead of a zero length string???
}

func startShell() {
	shell.Start()
}

func exitApplication(status int) {
	os.Exit(status)
}
