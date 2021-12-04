package main
import(
	"log"
	"os"
	"os/signal"
	"io/ioutil"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Authentication Token pulled from environment variable DGU_TOKEN
	TokenStream,err := ioutil.ReadFile("token.txt")
	Token := string(TokenStream)
	if Token == "" {
		return
	}

	// Create a new Discordgo session
	dg, err := discordgo.New(Token)
	if err != nil {
		log.Println("Error authenticating to the discord API\nError: ", err)
		return
	}
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil {
		log.Println("Error opening connection to the discord API\nError: ", err)
	}
	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate (s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content != "ping" {
		return
	}
	channel,err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		log.Println("Error opening the DM!\nError: ", err)
		s.ChannelMessageSend(m.ChannelID,"Something went wrong while opening the DM!",
		)
		return
	}
	_, err = s.ChannelMessageSend(channel.ID,"Pong!")
	if err != nil {
		log.Println("Error sending DM!\nError: ", err)
		s.ChannelMessageSend(m.ChannelID,"Something went wrong while sending the DM!")
	}
}
