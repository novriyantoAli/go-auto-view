package function

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/novriyantoAli/go-auto-view/model"
	"github.com/sirupsen/logrus"
)

// windowsSwitchWindow ...
func windowsSwitchWindow(profiles []model.Process) {
	for i := 0; i < len(profiles); i++ {
		robotgo.ActivePID(int32(profiles[i].ProcessID))
		title := robotgo.GetTitle(int32(profiles[i].ProcessID))
		fmt.Println("pid:", int32(profiles[i].ProcessID), "title:", title)

		time.Sleep(3 * time.Second)

		robotgo.KeyTap("t", "shift")
	}
}

func linuxSwitchWindow(applicationName string) {
	output, err := exec.Command("pgrep", applicationName).Output()
	if err != nil {
		logrus.Errorln(err)
	}

	s := strings.Split(strings.TrimSpace(fmt.Sprintf("%s", output)), "\n")
	
	outputWm, err := exec.Command("wmctrl", "-lp").Output()
	if err != nil {
		logrus.Error(err)
	}

	if len(s) > 0 {
		memoryAddress := make([]string, 0)
		x := strings.Split(fmt.Sprintf("%s", outputWm), "\n")
		for i := 0; i < len(x); i++ {
			y := strings.Split(x[i], " ")
			for j := 0; j < len(y); j++ {
				for k := 0; k < len(s); k++ {
					if strings.TrimSpace(s[k]) == strings.TrimSpace(y[j]) {
						memoryAddress = append(memoryAddress, y[0])
						break
					}
				}
			}
		}
		if len(memoryAddress) > 0 {
			for i := 0; i < len(memoryAddress); i++ {
				cmd := exec.Command("wmctrl", "-ia", memoryAddress[i])
				er := cmd.Run()
				if er != nil {
					logrus.Error(err)
				}

				time.Sleep(3 * time.Second)

				robotgo.KeyTap("t", "shift")
			}
		}
	}
	// logrus.Println(x[0])
}
