package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"github.com/abiosoft/ishell"
	"fmt"
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
	var helpText string
	response,  err := http.Get(defaultURL + "/help")
    if err != nil {
        errorCheck(err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        errorCheck(err)
				helpText += string(contents)
			}
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

func errorCheck(err error) {
	if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
	}
}
