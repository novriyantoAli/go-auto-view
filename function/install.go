package function

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/novriyantoAli/go-auto-view/model"
	"github.com/sirupsen/logrus"

	cp "github.com/otiai10/copy"
)

func Install(c *model.Config) {
	// delete directory profile default in system
	err := os.RemoveAll((c.DestinationProfile + c.SourceProfile[1:len(c.SourceProfile)]))
	if err != nil {
		logrus.Error(err)
	}

	args := []string{}

	for i := 0; i < len(c.InstallationBrowserCommand.Arguments); i++ {
		if strings.Contains(c.InstallationBrowserCommand.Arguments[i], "profile-directory") {
			args = append(args, (c.InstallationBrowserCommand.Arguments[i] + c.SourceProfile[1:len(c.SourceProfile)]))
		} else {
			args = append(args, c.InstallationBrowserCommand.Arguments[i])
		}
	}

	cmd := exec.Command(c.InstallationBrowserCommand.ApplicationName, args...)
	err = cmd.Run()
	if err != nil {
		logrus.Errorln(err)
	}

	logrus.Errorln(cmd.ProcessState.Pid())

	// get working directory
	path, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
	}

	// delete directory profile default in working directory
	err = os.RemoveAll((path + c.SourceProfile))
	if err != nil {
		logrus.Error(err)
	}

	// copy profile from system to working directory
	err = cp.Copy((c.DestinationProfile + c.SourceProfile[1:len(c.SourceProfile)]), (path + c.SourceProfile))
	if err != nil {
		logrus.Error(err)
	}

	for i := 0; i < len(c.Profile.Detail); i++ {
		fullname := c.Profile.Prefix + c.Profile.Split + c.Profile.Detail[i].ProfileName
		destination := c.DestinationProfile + fullname

		logrus.Info("deleting destination...")
		err = os.RemoveAll(destination)
		if err != nil {
			logrus.Error(err)
		}

		logrus.Info("copy from profile default to chrome profile home...")
		err = cp.Copy((path + c.SourceProfile), destination)
		if err != nil {
			logrus.Error(err)
		}

		logrus.Info("write credentials...")
		writeCredentials(&c.Profile.Detail[i].Username, &c.Profile.Detail[i].Password)

		argsOpenBrowser := []string{}
		for i := 0; i < len(c.OpenBrowserCommand.Arguments); i++ {
			if i == 0 {
				argsOpenBrowser = append(argsOpenBrowser, "https://accounts.google.com")
			} else if strings.Contains(c.OpenBrowserCommand.Arguments[i], "profile-directory") {
				argsOpenBrowser = append(argsOpenBrowser, (c.OpenBrowserCommand.Arguments[i] + fullname))
			} else {
				argsOpenBrowser = append(argsOpenBrowser, c.OpenBrowserCommand.Arguments[i])
			}
		}

		logrus.Info("open browser with argumens...")
		cm := exec.Command(c.OpenBrowserCommand.ApplicationName, argsOpenBrowser...)
		err = cm.Start()
		if err != nil {
			logrus.Error(err)
		}

		logrus.Info("delete if profile exists in working directory...")
		profileWorkingDirectory := path + "/" + fullname
		err = os.RemoveAll(profileWorkingDirectory)
		if err != nil {
			logrus.Error(err)
		}

		err = cp.Copy(destination, profileWorkingDirectory)
		if err != nil {
			logrus.Error(err)
		}

		writeCredentials(nil, nil)

		time.Sleep(3 * time.Second)
	}
}
