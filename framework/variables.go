package framework

import (
	"time"

	bot "github.com/bwmarrin/discordgo"
)

var (
	StartTime      = time.Now()
	CountMsg       int
	OnlineActulise int
	Session        *bot.Session

	//Variable framework
	CountCommand int
	SlashCommand string
)
