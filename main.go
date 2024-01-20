package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)


func main() {
	botToken := os.Getenv("BOT_TOKEN") // heroku envar
	launchpadChannel := "1091545478576996453"
	carlBot := "235148962103951360"
	emuEmoji := "<:emu:1198109556987936909>"

	sess, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	sess.Identify.Intents = discordgo.IntentsGuildMessages

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Author.ID == carlBot {
			return
		}

		if m.ChannelID == launchpadChannel && strings.Contains(strings.ToLower(m.Content), "/verifyme") {

			customEmoji := &discordgo.Emoji{Name: emuEmoji}
			response := fmt.Sprintf("<@&1161405957029515364> Verification Requested! Please review <@%s>'s application %s", m.Author.ID, customEmoji.MessageFormat())
			
			_, err := s.ChannelMessageSend(m.ChannelID, response)
			if err != nil {
				log.Println("Error sending message:", err)
			} 
		}
	})

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("EMU online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}