package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/code-you/GoProjects/discord-ping/config"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	println("Started bot", goBot.Token)
	if err != nil {
		fmt.Println("Now", err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println("Error requesting", err.Error())
		return
	}
	BotID = u.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot started successfully")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
