package function

import (
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/novriyantoAli/go-auto-view/model"
	"github.com/sirupsen/logrus"
)

func Run(c *model.Config, ch chan model.Stp) {

	destinationProfile := c.DestinationProfile

	logrus.Info("regenerate all cache folder...")
	for i := 0; i < len(c.Profile.Detail); i++ {
		regenerateCache(destinationProfile, (c.Profile.Prefix + c.Profile.Split + c.Profile.Detail[i].ProfileName))
	}

	// kill all process
	logrus.Info("kill all browser ...")
	cmd := exec.Command(c.KillBrowserCommand.ApplicationName, c.KillBrowserCommand.Arguments...)
	err := cmd.Run()
	if err != nil {
		logrus.Error(err)
	}

	// create profiles just for kill process if url is not contains playlist
	profiles := make([]model.Process, 0)

	logrus.Info("open browser by random selection ... ")

	var listRandomIndex []int
	for i := 0; i < c.CloneBrowser; i++ {
		var randomInt int
		for {
			rand.Seed(time.Now().UnixNano())
			randomInt = rand.Intn(len(c.Profile.Detail))
			if len(listRandomIndex) == 0 {
				listRandomIndex = append(listRandomIndex, randomInt)
				break
			} else {
				localCheck := false
				for v := 0; v < len(listRandomIndex); v++ {
					if listRandomIndex[v] == randomInt {
						localCheck = true
						break
					}
				}
				if !localCheck {
					listRandomIndex = append(listRandomIndex, randomInt)
					break
				}
			}
		}

		// open browser
		args := []string{}
		profileNamePref := c.Profile.Prefix + c.Profile.Split
		for i := 0; i < len(c.OpenBrowserCommand.Arguments); i++ {
			if strings.Contains(c.OpenBrowserCommand.Arguments[i], "profile-directory") {
				args = append(args, (c.OpenBrowserCommand.Arguments[i] + profileNamePref + c.Profile.Detail[randomInt].ProfileName))
			} else {
				args = append(args, c.OpenBrowserCommand.Arguments[i])
			}
		}

		command := exec.Command(c.OpenBrowserCommand.ApplicationName, args...)
		err := command.Start()
		if err != nil {
			logrus.Panic(err)
		}

		procss := model.Process{}
		procss.ProfileDetail = c.Profile.Detail[randomInt]
		procss.ProcessID = command.Process.Pid

		profiles = append(profiles, procss)

		time.Sleep(10 * time.Second)
	}

	listOfKill := make([]int, 0)
	listOfUsername := make([]string, 0)
	firstWindow := true
	for {
		res, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			break
		}

		inList := false
		for i := 0; i < len(c.JavascriptConfig.Step); i++ {
			if c.JavascriptConfig.Step[i].Url == res.Command {
				inList = true
				break
			}
		}

		if strings.Contains(res.Command, c.JavascriptConfig.PlaylistCode) {
			// listOfUsername menampung nama-nama yang telah di lakukan tab dan di aktivkan PIDnya
			if len(listOfUsername) == 0 {
				listOfUsername = append(listOfUsername, res.Username)
			} else {
				usernameInList := false
				for i := 0; i < len(listOfUsername); i++ {
					if listOfUsername[i] == res.Username {
						usernameInList = true
						break
					}
				}
				if !usernameInList {
					listOfUsername = append(listOfUsername, res.Username)
				}
			}
		}

		if len(listOfUsername) == c.CloneBrowser && firstWindow {
			switch runtime.GOOS {
			case "windows":
				windowsSwitchWindow(profiles)
			case "linux":
				linuxSwitchWindow(c.OpenBrowserCommand.ApplicationName)
			default:
				logrus.Warningln("os not regonized")
			}
			firstWindow = false
		}

		// check if command not contains url playlist
		if !strings.Contains(res.Command, c.JavascriptConfig.PlaylistCode) && !inList {
			// check if not in list
			// search username and add to list
			for i := 0; i < len(profiles); i++ {
				if profiles[i].ProfileDetail.Username == res.Username {
					listOfKill = append(listOfKill, profiles[i].ProcessID)
					break
				}
			}

			if len(listOfKill) >= c.CloneBrowser {
				logrus.Info("kill browser...")
				cmd := exec.Command(c.KillBrowserCommand.ApplicationName, c.KillBrowserCommand.Arguments...)
				err := cmd.Run()
				if err != nil {
					logrus.Error(err)
				}

				go Run(c, ch)

				break
			}
		}
	}
	// // initialize progress container, with custom width
	// p := mpb.New(mpb.WithWidth(64))

	// // create a single bar, which will inherit container's width
	// bar := p.New(c.MaxTime,
	// 	// BarFillerBuilder with custom style
	// 	mpb.BarStyle().Lbound("╢").Filler("▌").Tip("▌").Padding("░").Rbound("╟"),
	// 	mpb.PrependDecorators(
	// 		// display our name with one space on the right
	// 		decor.Name(c.JobName, decor.WC{W: len(c.JobName) + 1, C: decor.DidentRight}),
	// 		// replace ETA decorator with "done" message, OnComplete event
	// 		decor.OnComplete(
	// 			decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
	// 		),
	// 	),
	// 	mpb.AppendDecorators(decor.Percentage()),
	// )
	// // simulating some work
	// var i int64 = 0
	// for {
	// 	if i < c.MaxTime {
	// 		i++
	// 	} else {
	// 		break
	// 	}
	// 	time.Sleep(1 * time.Second)
	// 	bar.Increment()
	// }

	// // wait for our bar to complete and flush
	// p.Wait()

	// logrus.Info("kill browser...")
	// cmd := exec.Command(c.KillBrowserCommand.ApplicationName, c.KillBrowserCommand.Arguments...)
	// err := cmd.Run()
	// if err != nil {
	// 	logrus.Error(err)
	// }

	// time.Sleep(5 * time.Second)
	// Run(c, ch)
}
