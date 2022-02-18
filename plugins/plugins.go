package plugins
import(
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Load(s *discordgo.Session) {
	s.AddHandler(func(c *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(c, i)
		}
	})

	//s.AddHandler(memes)
	s.AddHandler(defaultCommands)
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	log.Println("The plugins have been loaded! ✓ ")
	return
}

func CreateCommands(s *discordgo.Session) {
	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "470408764995141632", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}

func buildHelp () string {
	res := "**Botshirt Help!**\n```Commands used in format [B,b]otshirt<,> <command>"
	res += "\nCommands:"
	res += "\nHelp:\t\t Sends this help message."
	res += "\nScramble:\t Scrambles your message."
	res += "\nShank:\t\tShank someone! Tag someone or just type their name in normally."
	res += "\nPing/pong:\tYou say ping, I say pong, or vice versa.```"
	return res
}

func scramble (m string) string {
	rand.Seed(time.Now().UnixNano())
	var res strings.Builder
	temp := strings.Split(m,"")
	for i := 0; i < len(m); i++ {
		index := rand.Intn(len(temp))
		res.WriteString(temp[index])
		temp = append(temp[:index], temp[(index+1):]...)
	}
	return res.String()
}

func isCommand (m string) bool {
	re, err := regexp.Compile(`([Bb]otshirt|<@!909138800482058372>)(\s|,\s)(\w+)`)
	if err != nil {
		log.Println("Error compiling regex object\nError: ", err)
	}
	return re.MatchString(m)
}

func getCommand (m string) string {
	re, err := regexp.Compile(`([Bb]otshirt|<@!909138800482058372>)(\s|,\s)(\w+)`)
	if err != nil {
		log.Println("Error compiling regex object\nError: ", err)
	}
	temp := re.FindSubmatch([]byte(m))
	return string(temp[len(temp)-1])
}

func trimCommand (m string) string {
	re, err := regexp.Compile(`([Bb]otshirt|<@!909138800482058372>)(\s|,\s)(\w+)\s`)
	if err != nil {
		log.Println("Error compiling regex object\nError: ", err)
	}
	return re.ReplaceAllString(m,"")
}

func defaultCommands (s *discordgo.Session, m *discordgo.MessageCreate) {
	channelID := "474960608207568896"
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.ChannelID == channelID && isCommand(m.Content) {
		command := getCommand(m.Content)
		//adds cases for default commands
		switch command {
			case "help":
				s.ChannelMessageSend(m.ChannelID,buildHelp())
			case "ping":
				s.ChannelMessageSend(m.ChannelID,"Pong!")
			case "pong":
				s.ChannelMessageSend(m.ChannelID,"Ping!")
			case "scramble":
				s.ChannelMessageSend(m.ChannelID,scramble(trimCommand(m.Content)))
			case "shank":
				s.ChannelMessageSend(m.ChannelID, trimCommand(m.Content) + " has been shanked!")
			default:
				s.ChannelMessageSend(m.ChannelID,"Error: **" + command + "** is not a registered command! For more information, please ask botshirt for help (i.e. Botshirt, help).")
		}
	}
}

func memes (s *discordgo.Session, t *discordgo.TypingStart) {
	channelID := "474960608207568896"
	if t.ChannelID == channelID {
		s.ChannelMessageSend(channelID,"Hurry up and finish typing <@!" + t.UserID + ">!!!")
	}
}




//Begin Slash Command Construction
var(
	commands = []*discordgo.ApplicationCommand {
		{
			Name: "ping",
			Description: "You say ping, I say pong.",
		},
		{
			Name: "pong",
			Description: "You say pong, I say ping.",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
		},
		"pong": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Ping!",
				},
			})
		},
	}
)