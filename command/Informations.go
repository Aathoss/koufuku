package command

import (
	"gitlab.com/koufuku/framework"
)

//DynmapDropURL retourne un message avec les informations de la dynmap
func Informations(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	embed := framework.NewEmbed().
		SetTitle("Vous pouvez me retrouver ici aussi :").
		SetColor(0xAD9F91).
		SetDescription("**Discord :** <https://discord.gg/qxVRVcPV3N>" +
			"\n**Instagram :** <https://www.instagram.com/koufuku_shop/>" +
			"\n**Facebook :** à venir" +
			"\n**TikTok :** à venir" +
			"\n**Twitch :** à venir" +
			"\n").MessageEmbed

	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}
