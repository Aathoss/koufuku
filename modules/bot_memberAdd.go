package modules

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/koufuku/framework"
	"gitlab.com/koufuku/logger"
)

var (
	titre string
)

func GuildMemberAdd(s *discordgo.Session, join *discordgo.GuildMemberAdd) {
	//ajoute d'un grade lors de l'arrivé
	/* s.GuildMemberRoleAdd(viper.GetString("GuildID"), join.User.ID, viper.GetString("RoleID.Bienvenue")) */

	countMembers, err := s.State.Guild(viper.GetString("GuildID"))
	if err != nil {
		logger.DebugLogger.Println(err)
		return
	}

	t1 := time.Now()
	if t1.Hour() >= 6 && t1.Hour() <= 18 {
		titre = "Bonjour"
	} else {
		titre = "Bonsoir"
	}

	//affichage d'un message de bienvenue
	embed := framework.NewEmbed().
		SetTitle(":inbox_tray: " + titre + ", " + join.User.Username).
		SetColor(viper.GetInt("EmbedColor.Background")).
		SetDescription("Bienvenue sur le discord **Koufuku_Shop**, \n\nMerci de bien vouloir vous respecter les un, les autres !\nÊtre conviviale et poli.").MessageEmbed

	s.ChannelMessageSendEmbed(viper.GetString("ChannelID.Trafic"), embed)
	framework.LogsChannel(":inbox_tray: [**" + strconv.Itoa(countMembers.MemberCount) + "**] **" + join.User.Username + "** Viens de nous rejoindre.")
}
