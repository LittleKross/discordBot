package main
import(
	"log"
	"os"
	"os/signal"
	"io/ioutil"
	"syscall"
	
	"github.com/littlekross/discordBot/plugins"
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
	plugins.Load(dg)
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

