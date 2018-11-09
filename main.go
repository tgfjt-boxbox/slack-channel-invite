package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tgfjt-boxbox/slack-channel-invite/config"
	"github.com/tgfjt-boxbox/slack-channel-invite/models"
	"github.com/urfave/cli"
)

func main() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}

	app := cli.NewApp()
	app.Name = "scit: slack channel invite tool"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{}
	app.Action = func(c *cli.Context) error {
		conf := config.GetConfig()

		var userIds []string
		var channelIds []string

		members := models.GetMembers()
		channels := models.GetChannels()

		getUserUid := func(name string) {
			uid := members.GetUidByName(name)
			userIds = append(userIds, uid)
		}

		for _, name := range conf.TargetUsers {
			getUserUid(name)
		}

		getChanelUid := func(cn string) {
			uid := channels.GetUidByName(cn)
			channelIds = append(channelIds, uid)
		}

		for _, c := range conf.TargetChannels {
			getChanelUid(c)
		}

		f := func(cID string) {
			for _, uID := range userIds {
				fmt.Println(uID, cID)
				models.Invite(uID, cID)
			}
		}

		for _, cID := range channelIds {
			go f(cID)
		}

		time.Sleep(time.Second)

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
