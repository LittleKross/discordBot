package main
import(
	"log"
	"os"

	"githum.com/bwmarrin/discordgo"
)
int plugins
func load(s *discordgo.Session) {
	for i:=0;i<plugins;i++ {
		//s.AddHandler()
		log.Println("The plugin loader has been loaded")
	}
}
