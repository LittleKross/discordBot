package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	
	
	//"github.com/littlekross/discordBot/bot/config"
	"github.com/littlekross/discordBot/plugins"
	"github.com/bwmarrin/discordgo"
)

type bot struct {
	client *discordgo.Session
	//config *config
}

func New(token string) *bot {
	//c := config.New(token)
	s, err := discordgo.New(token)
	if err != nil {
		log.Fatal("Error authenticating to the discord API\nError:\n", err)
	}
	b := &bot {
		client: s,
	}
	return b
}

func (b *bot) Run() {
	plugins.Load(b.client)
	b.client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	err := b.client.Open()
	if err != nil {
		log.Fatal("Error opening connection to the discord API\nError:\n", err)
	}
	
	plugins.CreateCommands(b.client)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	
	b.client.Close()
	
}
