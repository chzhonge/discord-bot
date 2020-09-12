package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"math/rand"
	"time"
	"github.com/bwmarrin/discordgo"
)

var haruImageMap = map[string]string{
	"haruHappy" : "https://imgur.com/PBrykp3",
	"haruFight" : "https://imgur.com/kTCco71",
	"haruCry" : "https://imgur.com/KrDgeXS",
	"haruShock" : "https://imgur.com/A3d709U",
	"haruAmazing" : "https://imgur.com/vSAuUk2",
	"haruNose" : "https://imgur.com/9E8HKJF",
	"haruHi" : "https://imgur.com/XinYTZ8",
	"haruMilk" : "https://imgur.com/d2t9VTc",
	"haruSad" : "https://imgur.com/HHfqUBf",
	"haruFlower" : "https://imgur.com/wvJQLCL",
	"haruMoney" : "https://imgur.com/AnX5Cnt",
	"haruJob" : "https://imgur.com/8KmaPQX",
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if -1 == strings.Index(m.Content, "!q ") {
		return
	}

	action := m.Content[3:]

	switch action {
	case "help":
		msg := fmt.Sprintf(
			"```%s\n%s\n%s\n```",
			"please input '!q [command]' ",
			"sexy : show %%%% .",
			"haru : display haru gif from random.",
		)
		s.ChannelMessageSend(m.ChannelID, msg)
	case "sexy":
		s.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=okRkOdJXH_E&feature=youtu.be&t=14")
		return
	case "haru":
		rand.Seed(time.Now().UnixNano())
		j := 0
		keys := make([]string, len(haruImageMap))
		for k := range haruImageMap {
			keys[j] = k
			j++
		}
		randomKey := keys[rand.Intn(len(keys))]
		s.ChannelMessageSend(m.ChannelID, haruImageMap[randomKey])
		return
	}

	if -1 != strings.Index(action, "haru") && 4 < len(action) {
		feeling := "haru" + strings.Title(action[4:]);
		if val, ok:= haruImageMap[feeling]; ok {
			s.ChannelMessageSend(m.ChannelID, val)
			return
		}
	}
}