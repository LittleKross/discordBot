package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	
	
	"github.com/littlekross/discordBot/plugins"
	"github.com/bwmarrin/discordgo"
)

type bot struct {

	client *discordgo.Session

}

func New(token string) *bot {
	dg, err := discordgo.New(token)
	if err != nil {
		log.Println("Error authenticating to the discord API\nError: ", err)
	}
	b := &bot {
		client: dg,
	}
	return b
}

func (b *bot) Run() {
	plugins.Load(b.client)
	b.client.Identify.Intents = discordgo.IntentsGuildMessages
	err := b.client.Open()
	if err != nil {
		log.Println("Error opening connection to the discord API\nError: ", err)
	}
	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	
	b.client.Close()
	
}
