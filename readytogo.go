package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Router for the REST API
func Router() *mux.Router {
	r := mux.NewRouter()
	return r
}

// StartServer starts the server
func StartServer() error {
	r := Router()
	n := negroni.New(negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("web/dist")))
	n.UseHandler(r)
	n.Run(":7777")
	return nil
}

// StartClient starts the client
func StartClient() error {
	return nil
}

func topHelp() {
	fmt.Printf("Usage of readytogo: readytogo <command> [<args>]\n\n")
	fmt.Printf("These are the available commands:\n")
	fmt.Println("  server   Start server")
	fmt.Println("  client   Start client")
}

func start(args []string) {
	serverCommand := flag.NewFlagSet("server", flag.ContinueOnError)
	serverCommand.String("addr", "0.0.0.0:7777", "Server address")
	clientCommand := flag.NewFlagSet("client", flag.ContinueOnError)
	clientCommand.String("addr", "0.0.0.0:7777", "Server address")
	if len(args) == 0 {
		topHelp()
		return
	}
	switch args[0] {
	case "-h":
		topHelp()
	case "server":
		serverCommand.Parse(args[1:])
		err := StartServer()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	case "client":
		clientCommand.Parse(args[1:])
		err := StartClient()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	start(os.Args[1:])
}
