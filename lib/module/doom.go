package module

import (
	"fmt"
	"github.com/dvdmuckle/irkbot/lib/configure"
	"github.com/dvdmuckle/irkbot/lib/message"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const doomHost = "localhost:8000/api/player/actions"

func Helpdoom() []string {
	s := "doom <command> - play doom!"
	return []string{s}
}

//Sanitize input
func doom(cfg *configure.Config, in *message.InboundMsg, actions *Actions) {
	doomCommand := "enter a command, dipstick"
	if len(in.MsgArgs[1:]) == 0 {
		actions.Say(doomCommand)
		return
	}
	if in.MsgArgs[1:] == "shoot"|"forward"|"backward"|"left"|"right" {
		actions.Say("invalid command, comands are: shoot, forward, backward, left, right")
		return
	}
	doomCommand = in.MsgArgs[1:]
	post(doomCommand)
}

//Perform actual POST
func post(doomCommand string) (string, error) {
	//There has got to be a better way to do this
	switch doomCommand {
	case "shoot":
		body := strings.NewReader(`{"type":"shoot"}`)
	case "forward":
		body := strings.NewReader(`{"type":"forward"}`)
	case "backward":
		body := strings.NewReader(`{"type":"backward"}`)
	case "left":
		body := strings.NewReader(`{"type":"turn-left"}`)
	case "right":
		body := strings.NewReader(`{"type":"turn-right"}`)
	case "open":
		body := strings.NewReader(`{"type":"open"}`)
	default:
		return //This shouldn't actually happen since we sanitize earlier, but just in case
	}
	req, err := http.NewRequest("POST", doomHost, body)
	if err != nil {
		// handle err
		fmt.Fprintln(os.Stderr, err)
		actions.Say("something borked, try again")
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Fprintln(os.Stderr, err)
		actions.Say("something borked, try again")
		return
	}
	defer resp.Body.Close()
}
