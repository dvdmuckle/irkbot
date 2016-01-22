package modpm

import (
    "fmt"
    "strings"
    urbandict "github.com/davidscholberg/go-urbandict"
    "github.com/davidscholberg/irkbot/lib"
)

func Urban(p *lib.Privmsg) bool {
    if ! strings.HasPrefix(p.Msg, "..urban") {
        return false
    }

    if (strings.HasPrefix(p.Msg, "..urban-trending")) {
        showTrending(p)
    } else {
        showDefinition(p)
    }

    return true
}

func showDefinition(p *lib.Privmsg) {
    var def *urbandict.Definition
    var err error
    nick := p.Event.Nick
    isWotd := strings.HasPrefix(p.Msg, "..urban-wotd")
    if isWotd {
        def, err = urbandict.WordOfTheDay()
    } else if len(p.MsgArgs) == 1 {
        def, err = urbandict.Random()
    } else {
        def, err = urbandict.Define(strings.Join(p.MsgArgs[1:], " "))
    }
    if err != nil {
        lib.Say(p, fmt.Sprintf("%s: %s", nick, err.Error()))
        return
    }

    // TODO: implement max message length handling

    if isWotd {
        lib.Say(p, fmt.Sprintf("%s: Word of the day: \"%s\"", nick, def.Word))
    } else {
        lib.Say(p, fmt.Sprintf("%s: Top definition for \"%s\"", nick, def.Word))
    }
    for _, line := range strings.Split(def.Definition, "\r\n") {
        lib.Say(p, fmt.Sprintf("%s: %s", nick, line))
    }
    lib.Say(p, fmt.Sprintf("%s: Example:", nick))
    for _, line := range strings.Split(def.Example, "\r\n") {
        lib.Say(p, fmt.Sprintf("%s: %s", nick, line))
    }
    lib.Say(p, fmt.Sprintf("%s: permalink: %s", nick, def.Permalink))
}

func showTrending(p *lib.Privmsg) {
    nick := p.Event.Nick

    trendingWords, err := urbandict.Trending()
    if err != nil {
        lib.Say(p, fmt.Sprintf("%s: %s", nick, err.Error()))
        return
    }

    lib.Say(p, fmt.Sprintf("%s: Top %d trending words:",
        nick,
        len(trendingWords)))

    for i, word := range trendingWords {
        lib.Say(p, fmt.Sprintf("%s: %d. %s", nick, i + 1, word))
    }

    return
}
