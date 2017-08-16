package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dvdmuckle/irkbot/lib/configure"
	"github.com/dvdmuckle/irkbot/lib/message"
	"net/http"
	"os"
	"strings"
)

type doomStruct struct {
	Type string `json:"type"`
}

const doomHost = "http://localhost:8000/api/player/actions"

func HelpDoom() []string {
	s := "doom <command> - play doom!"
	return []string{s}
}

//Sanitize input
func Doom(cfg *configure.Config, in *message.InboundMsg, actions *Actions) {
	doomCommand := "enter a command, dipstick"
	if len(in.MsgArgs[1:]) == 0 {
		actions.Say(doomCommand)
		return
	}
	//Ugly switch statement to sanitize input because RESTful Doom doesn't
	doomCommand = strings.Join(in.MsgArgs[1:], " ")
	switch doomCommand {
	case "shoot":
	case "forward":
	case "backward":
	case "left":
	case "right":
	case "use":
	default:
		actions.Say("invalid command, comands are: shoot, forward, backward, left, right, use")
		return
	}
	doomToPost := doomStruct{Type: doomCommand}
	jsonValue, _ := json.Marshal(doomToPost)
	resp, err := http.Post(doomHost, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		// handle err
		fmt.Fprintln(os.Stderr, err)
		actions.Say("something borked, try again")
		return
	}
	defer resp.Body.Close()
}
