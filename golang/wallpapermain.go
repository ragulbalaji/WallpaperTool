/*                  ____                               __              __
 *  _      ______ _/ / /___  ____ _____  ___  _____   / /_____  ____  / /
 * | | /| / / __ `/ / / __ \/ __ `/ __ \/ _ \/ ___/  / __/ __ \/ __ \/ / 
 * | |/ |/ / /_/ / / / /_/ / /_/ / /_/ /  __/ /     / /_/ /_/ / /_/ / /  
 * |__/|__/\__,_/_/_/ .___/\__,_/ .___/\___/_/      \__/\____/\____/_/   
 *                 /_/         /_/                                       
 *
 * It is what it is. Ragul Balaji (C) 2018 All Rights Reserved.
 */

package main

import (
	"fmt"
	"./icon"

	"github.com/reujab/wallpaper"
	"github.com/getlantern/systray"
)

func main() {
	background, err := wallpaper.Get()

	if err != nil {
		panic(err)
	}

	fmt.Println("Current wallpaper:", background)

	systray.Run(onReady, onExit)

	//wallpaper.SetFromURL("https://unsplash.it/1440/900/?random")
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Wallpaper Tool")
	systray.SetTooltip("RB2k18")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	
}

func onExit() {
	fmt.Println("I Quit!")// clean up here
}