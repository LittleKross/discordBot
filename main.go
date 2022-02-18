package main
import(
	"io/ioutil"
	"log"
	
	"github.com/littlekross/discordBot/bot"
)

func main() {
	// Authentication Token pulled from local file token.txt
	TokenStream,err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatal("Error opening token file\nError:\n", err)
		return
	}
	Token := string(TokenStream)
	if Token == "" {
		return
	}

	b := bot.New(Token)
	b.Run()
}

