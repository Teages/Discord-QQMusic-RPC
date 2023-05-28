package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/Teages/go-autostart"
	"github.com/getlantern/systray"
)

var (
	appName    = "Discord QQMusic RPC"
	appVersion = "v1.0.1"
)

func initTray(onReady func(), onExit func()) {
	systray.Run(func() {
		systray.SetIcon(AppIcon)
		systray.SetTitle(appName)
		systray.SetTooltip(appName)

		mAbout := systray.AddMenuItem(fmt.Sprint(appName, " ", appVersion), "About")

		go func() {
			for {
				<-mAbout.ClickedCh
				openBrowser("https://github.com/Teages/Discord-QQMusic-RPC")
			}
		}()

		systray.AddSeparator()

		// Start with OS
		aStartWithOS := &autostart.App{
			Name:        appName,
			DisplayName: "Auto start " + appName,
			Exec:        []string{GetSelfPath()},
		}
		mStartWithOS := systray.AddMenuItemCheckbox("Start with OS", "Start with OS", aStartWithOS.IsEnabled())
		go func() {
			for {
				<-mStartWithOS.ClickedCh
				if aStartWithOS.IsEnabled() {
					aStartWithOS.Disable()
					mStartWithOS.Uncheck()
				} else {
					aStartWithOS.Enable()
					mStartWithOS.Check()
				}
			}
		}()

		mQuit := systray.AddMenuItem("Quit", "Quit the app")
		go func() {
			<-mQuit.ClickedCh
			systray.Quit()
		}()

		onReady()
	}, onExit)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func GetSelfPath() string {
	exePath, _ := os.Executable()
	return exePath
}
