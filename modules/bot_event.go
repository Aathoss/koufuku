package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/koufuku/framework"
	"gitlab.com/koufuku/logger"
)

var (
	ready bool
)

func Ready(s *discordgo.Session, Event *discordgo.Event) {
	framework.Session = s

	if Event.Type == "READY" && ready == false {
		ready = true
		s.UpdateGameStatus(0, viper.GetString("Motd"))
		logger.InfoLogger.Println("Le bot est dispo. [Appuyez sur CTRL+C pour l'arrêter !]")
		framework.LogsChannel("[:tools:] **Koufuku** à correctement démarré")

		/* commands := "[]*discordgo.ApplicationCommand{" + framework.SlashCommand + "}"
		for _, v := range commands {
			fmt.Println(v)
			_, err := s.ApplicationCommandCreate(s.State.User.ID, viper.GetString("GuildID"), v)
			if err != nil {
				log.Panicf("Cannot create '%v' command: %v", v.Name, err)
			}
		}
		s.ApplicationCommandCreate() */
	}
}
