package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	bot "github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/koufuku/command"
	"gitlab.com/koufuku/framework"
	"gitlab.com/koufuku/logger"
	"gitlab.com/koufuku/modules"
)

// Variable
var (
	CmdHandler *framework.CommandHandler
)

func main() {

	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	dg, err := bot.New("Bot " + viper.GetString("ID"))
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de la session discord,", err)
		return
	}

	dg.AddHandler(modules.Ready)
	dg.AddHandler(modules.GuildMemberAdd)
	dg.AddHandler(modules.GuildMemberLeave)
	dg.AddHandler(modules.LevelingMessages)
	dg.AddHandler(commandHandler)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	err = dg.Open()
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de la connexion,", err)
		return
	}

	go func() {
		for {
			consoleReader := bufio.NewReader(os.Stdin)

			input, _ := consoleReader.ReadString('\n')

			input = strings.ToLower(input)

			if strings.HasPrefix(input, "bye") {
				framework.LogsChannel("[:tools:] **Koufuku** s'est déconnecté de l'univers !")

				fmt.Println("\nUptime : " + framework.Calculetime(framework.StartTime.Unix(), 0) +
					"\nMessage total : " + strconv.Itoa(framework.CountMsg) +
					"\nRoutine : " + strconv.Itoa(runtime.NumGoroutine()) +
					"\n\nAllez bonne route ++ \n")
				os.Exit(100)
			}
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

out:
	for {
		select {
		case <-sc:
			framework.LogsChannel("[:tools:] **Koufuku** s'est déconnecté de l'univers !")
			break out
		}
	}
}

func commandHandler(s *bot.Session, m *bot.MessageCreate) {

	framework.Session = s
	user := m.Author

	if user.ID == s.State.User.ID || user.Bot {
		return
	}

	if viper.GetBool("Dev.PrintMessage") == true {
		log.Println(m.Content)
	}

	framework.CountMsg = framework.CountMsg + 1

	content := m.Content
	if len(content) <= len(viper.GetString("PrefixMsg")) {
		return
	}
	if content[:len(viper.GetString("PrefixMsg"))] != viper.GetString("PrefixMsg") {
		return
	}
	content = content[len(viper.GetString("PrefixMsg")):]
	if len(content) < 1 {
		return
	}
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de l'obtention du channel,", err)
		return
	}
	staff := 0
	if channel.Type != discordgo.ChannelTypeDM {
		staff = framework.VerifStaff(m.Member.Roles)
	}
	checkCmdName := CmdHandler.CheckCmd(content)
	command, found, permission := CmdHandler.Get(checkCmdName, staff)
	if !found {
		return
	}
	if permission == false {
		s.ChannelMessageSendEmbed(m.ChannelID, framework.EmbedPermissionFalse)
		return
	}
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de l'obtention de la guilde,", err)
		return
	}
	ctx := framework.NewContext(s, guild, channel, user, m, CmdHandler, checkCmdName, staff)
	messageSplit := strings.Fields(content)
	if len(strings.Fields(checkCmdName)) == 1 {
		ctx.Args = messageSplit[1:]
	}
	if len(strings.Fields(checkCmdName)) == 2 {
		ctx.Args = messageSplit[2:]
	}
	c := *command
	c(*ctx)
}

func registerCommands() {
	CmdHandler.Register("info", []string{"insta", "instagram", "facebook", "twitch", "tiktok", "discord", "informations"}, 0, command.Informations, "Message d'informations contenant les divers réseaux sociaux")
}
