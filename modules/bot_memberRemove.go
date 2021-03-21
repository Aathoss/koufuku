package modules

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/koufuku/framework"
	"gitlab.com/koufuku/logger"
)

func GuildMemberLeave(s *discordgo.Session, leave *discordgo.GuildMemberRemove) {
	//Trafic des membres du discord [leave]

	countMembers, err := s.State.Guild(viper.GetString("GuildID"))
	if err != nil {
		logger.DebugLogger.Println(err)
		return
	}

	framework.LogsChannel(":outbox_tray: [**" + strconv.Itoa(countMembers.MemberCount) + "**] **" + leave.User.Username + "** vient de nous quitté.")
}
