package main

import (
	"crypto/md5"
	"encoding/binary"
	"strconv"
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
				updateRPC("", "", 0)
			} else {
				sp := strings.SplitN(str, "-", 2)
				if len(sp) == 2 {
					updateRPC(sp[0], sp[1], getCD(str))
				} else {
					updateRPC("", "", 0)
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

func initDiscordRPC(discordAppId string) func(string, string, int) {
	title := ""
	artist := ""

	err := client.Login(discordAppId)
	if err != nil {
		panic(err)
	}

	isLogin := true

	update := func(newTitle, newArtist string, cd int) {
		if newTitle != title || newArtist != artist {
			title = newTitle
			artist = newArtist
			cdId := "cd_icon_" + strconv.Itoa(cd)
			println(cdId)

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
				LargeImage: cdId,
				LargeText:  "",
				SmallImage: "qqmusic",
				SmallText:  "Listening QQ Music",
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

func getCD(str string) int {
	size := 16
	hasher := md5.New()
	hasher.Write([]byte(str))
	hash := hasher.Sum(nil)
	return int(binary.BigEndian.Uint64((hash)) % uint64(size))
}
