package main

import (
	"fmt"
	"log"
	"time"

	"github.com/c-bata/go-prompt"
	minecraft "github.com/willroberts/minecraft-client"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	client, err := minecraft.NewClient(minecraft.ClientOptions{
		Hostport: "127.0.0.1:25575",
		Timeout:  5 * time.Second, // Optional, this is the default value.
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Authenticate("password"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("To quit, type `q`")

	for {
		t := prompt.Input("> ", completer)

		if t == "q" {
			break
		}

		resp, err := client.SendCommand(t)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.Body)
	}
}
