package function

import (
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"

	"github.com/novriyantoAli/go-auto-view/model"
	"github.com/sirupsen/logrus"

	cp "github.com/otiai10/copy"
)

/**
* Install
* 1. delete directory profile_default in system
* 2. open browser with profile_default
* 3. delete directory profile_default in local
* 4. copy directory profile_default from system to local
* 5. copy directory profile_default to profile_user and open it
* 6. delete directory profile_user in local
* 7. copy directory profile_user from system to local
 */
func Install(c *model.Config) {

	destinationProfile := c.DestinationProfile
	// delete directory profile_default in system
	err := os.RemoveAll((destinationProfile + c.SourceProfile[1:len(c.SourceProfile)]))
	if err != nil {
		logrus.Error(err)
	}

	// open browser with profile_default
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
		logrus.Panicln(err)
	}

	// get working directory
	path, err := os.Getwd()
	if err != nil {
		logrus.Panicln(err)
	}

	// delete directory profile default in working directory
	err = os.RemoveAll((path + c.SourceProfile))
	if err != nil {
		logrus.Panicln(err)
	}

	// copy profile from system to working directory
	err = cp.Copy((destinationProfile + c.SourceProfile[1:len(c.SourceProfile)]), (path + c.SourceProfile))
	if err != nil {
		logrus.Panicln(err)
	}

	for i := 0; i < len(c.Profile.Detail); i++ {
		fullname := c.Profile.Prefix + c.Profile.Split + c.Profile.Detail[i].ProfileName
		destination := destinationProfile + fullname

		logrus.Info("deleting destination...")
		err = os.RemoveAll(destination)
		if err != nil {
			logrus.Panicln(err)
		}

		logrus.Info("copy from profile default to chrome profile home...")
		err = cp.Copy((path + c.SourceProfile), destination)
		if err != nil {
			logrus.Panicln(err)
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
		err = cm.Run()
		if err != nil {
			logrus.Panicln(err)
		}

		logrus.Info("delete if profile exists in working directory...")
		profileWorkingDirectory := path + "/" + fullname
		err = os.RemoveAll(profileWorkingDirectory)
		if err != nil {
			logrus.Panicln(err)
		}

		logrus.Info("copy directory from system to working directory...")
		err = cp.Copy(destination, profileWorkingDirectory)
		if err != nil {
			logrus.Panicln(err)
		}

		writeCredentials(nil, nil)

		time.Sleep(3 * time.Second)
	}
}

/**
* InstallSomeProfile(c *model.Config)
* 1. check if profile_default not exits, and create it if not exists
 */
func InstallSomeProfile(c *model.Config) {
	// check profile_default if exists or not
	path, err := os.Getwd()
	if err != nil {
		logrus.Panicln(err.Error())
	}

	user, err := user.Current()
	if err != nil {
		logrus.Panicln(err.Error())
	}

	destinationProfile := strings.Replace(c.DestinationProfile, "$username", user.Username, -1)

	_, err = os.Stat((path + c.SourceProfile))
	if err != nil {
		if os.IsNotExist(err) {
			// delete directory profile_default in system

			err := os.RemoveAll((destinationProfile + c.SourceProfile[1:len(c.SourceProfile)]))
			if err != nil {
				logrus.Error(err)
			}

			// open browser with profile_default
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
				logrus.Panicln(err)
			}

			// copy profile from system to working directory
			err = cp.Copy((destinationProfile + c.SourceProfile[1:len(c.SourceProfile)]), (path + c.SourceProfile))
			if err != nil {
				logrus.Panicln(err)
			}
		} else {
			logrus.Panic(err.Error())
		}
	}

	for i := 0; i < len(c.Profile.Detail); i++ {
		if _, err := os.Stat(path + "/" + c.Profile.Prefix + c.Profile.Split + c.Profile.Detail[i].ProfileName); os.IsNotExist(err) {
			fullname := c.Profile.Prefix + c.Profile.Split + c.Profile.Detail[i].ProfileName
			destination := destinationProfile + fullname

			logrus.Info("deleting destination...")
			err = os.RemoveAll(destination)
			if err != nil {
				logrus.Panicln(err)
			}

			logrus.Info("copy from profile default to chrome profile home...")
			err = cp.Copy((path + c.SourceProfile), destination)
			if err != nil {
				logrus.Panicln(err)
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
			err = cm.Run()
			if err != nil {
				logrus.Panicln(err)
			}

			logrus.Info("delete if profile exists in working directory...")
			profileWorkingDirectory := path + "/" + fullname
			err = os.RemoveAll(profileWorkingDirectory)
			if err != nil {
				logrus.Panicln(err)
			}

			logrus.Info("copy directory from system to working directory...")
			err = cp.Copy(destination, profileWorkingDirectory)
			if err != nil {
				logrus.Panicln(err)
			}

			writeCredentials(nil, nil)

			time.Sleep(3 * time.Second)
		}
	}
}
