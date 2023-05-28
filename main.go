package main

import (
	"strings"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func main() {
	initTray(func() {
		updateRPC := initDiscordRPC("1112302853944508416")
		update := func() {
			time.Sleep(time.Second * 1)
			str := findQQMusic()
			if str == "" {
				updateRPC("", "")
			} else {
				sp := strings.SplitN(str, "-", 2)
				if len(sp) == 2 {
					updateRPC(sp[0], sp[1])
				} else {
					updateRPC("", "")
				}
			}
		}
		go func() {
			for {
				update()
			}
		}()
	}, func() {

	})
}

func findQQMusic() string {
	return GetDesktopWindowName("QQMusic_Daemon_Wnd")
}

func initDiscordRPC(discordAppId string) func(string, string) {
	title := ""
	artist := ""

	err := client.Login(discordAppId)
	if err != nil {
		panic(err)
	}

	isLogin := true

	update := func(newTitle, newArtist string) {
		if newTitle != title || newArtist != artist {
			title = newTitle
			artist = newArtist

			if title == "" && isLogin {
				client.Logout()
				isLogin = false
				return
			} else if !isLogin {
				client.Login(discordAppId)
				isLogin = true
			}
			now := time.Now()

			client.SetActivity(client.Activity{
				State:      artist,
				Details:    title,
				LargeImage: "qqmusic",
				LargeText:  "",
				SmallImage: "",
				SmallText:  "",
				Party:      nil,
				Timestamps: &client.Timestamps{
					Start: &now,
				},
				Buttons: []*client.Button{},
			})
		}
	}
	return update
}
