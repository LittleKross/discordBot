package plugins
import(
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Load(s *discordgo.Session) {
	s.AddHandler(messageScramble)
	s.AddHandler(pingPong)
	s.AddHandler(shank)
	for i:=0;i<1;i++ {
		//s.AddHandler()
		log.Println("The plugins have been loaded! ✓ ")
	}
	return
}

func scramble (m string) string {
	rand.Seed(time.Now().UnixNano())
	var res strings.Builder
	temp := strings.Split(m,"")
	for i := 0; i < len(m); i++ {
		res.WriteString(temp[rand.Intn(len(m))])
	}
	return res.String()
}

func messageScramble (s *discordgo.Session, m *discordgo.MessageCreate) {
	//log.Println("The messageScramble plugin has loaded ✓")
	if m.ChannelID == "474960608207568896" {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content != "ping" && m.Content != "pong" && !strings.Contains(m.Content,"Botshirt, shank "){
			s.ChannelMessageSend(m.ChannelID,scramble(m.Content))
		}
	}
}
func pingPong (s *discordgo.Session, m *discordgo.MessageCreate) {
	//log.Println("The pingPong plugin has loaded ✓")
	//personal: 886821410595561492 //allegiant: 474960608207568896
	if m.ChannelID == "474960608207568896"{
		if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content == "ping" {
			s.ChannelMessageSend(m.ChannelID,"Pong!")
		}
		if m.Content == "pong" {
			s.ChannelMessageSend(m.ChannelID,"Ping!")
		}
	}
}
func shank (s *discordgo.Session, m *discordgo.MessageCreate) {
	//log.Println("The shank plugin has loaded ✓")
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.ChannelID == "474960608207568896" {
		if strings.Contains(m.Content, "Botshirt, shank ") {
			s.ChannelMessageSend(m.ChannelID,m.Content[16:] + " has been shanked!")
		}
	}
}