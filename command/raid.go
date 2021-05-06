package command

import (
	"gitlab.com/koufuku/framework"
)

func raid(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	/* if len(ctx.Args) != 0 {

		list := strings.Split(ctx.Args[0], "?")
		count := len(list)

		for _, value := range list {
		}
	} */
}
