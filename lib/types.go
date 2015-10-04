package lib

import (
    goirc "github.com/thoj/go-ircevent"
)

// TODO: give modules their own sections in config?
type Config struct {
    User struct {
        Nick string
        User string
    }
    Server struct {
        Host string
        Port string
    }
    Channel struct {
        Channelname string
        Greeting string
    }
    Module struct {
        Insult_swearfile string
    }
}

type Privmsg struct {
    Msg string
    MsgArgs []string
    Dest string
    Event *goirc.Event
    Conn *goirc.Connection
    SayChan chan Say
}

type Say struct {
    Conn *goirc.Connection
    Dest string
    Msg string
}
