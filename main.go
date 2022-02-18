package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/novriyantoAli/go-auto-view/function"
	"github.com/novriyantoAli/go-auto-view/model"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	var c model.Config

	json.Unmarshal(function.Decryption(), &c)

	ch := make(chan model.Stp)

	// create echo framework
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(ec echo.Context) error {
		return ec.JSON(http.StatusOK, c.JavascriptConfig)
	})

	e.GET("/installation", func(c echo.Context) error {
		jsonFile, err := os.Open("credentials.json")
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		defer jsonFile.Close()

		bv, err := ioutil.ReadAll(jsonFile)

		if err != nil {
			logrus.Error(err)

			return c.JSON(http.StatusInternalServerError, err)
		}

		var cre model.Credentials
		err = json.Unmarshal([]byte(bv), &cre)

		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, cre)
	})

	// get url
	e.POST("/step", func(c echo.Context) error {
		var stp model.Stp
		if err := c.Bind(&stp); err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		logrus.Println("username: ", stp.Username)
		logrus.Println("command: ", stp.Command)

		ch <- stp

		return c.JSON(http.StatusCreated, stp)
	})

	go func() {
		e.Logger.Fatal(e.Start(c.Port))
	}()

	defer e.Shutdown(context.Background())

	reader := bufio.NewReader(os.Stdin)
	logrus.Println("STARTING AILA BOT 0.1")

	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			logrus.Errorln(err)
			break
		}

		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)

		if strings.Compare("1", text) == 0 {
			logrus.Println("installation started...")
			function.Install(&c)
		} else if strings.Compare("0", text) == 0 {
			logrus.Println("thanks...")
			os.Exit(0)
		} else if strings.Compare("2", text) == 0 {
			logrus.Println("checking directory...")
			directoryReady := true
			for i := 0; i < len(c.Profile.Detail); i++ {
				path, err := os.Getwd()
				if err != nil {
					logrus.Error(err)
					directoryReady = false
					break
				}

				if _, err := os.Stat(path + "/" + c.Profile.Prefix + c.Profile.Split + c.Profile.Detail[i].ProfileName); os.IsNotExist(err) {
					logrus.Error("you dont have default configuration directory, try to install it...")
					directoryReady = false
					break
				}
			}

			if directoryReady {
				logrus.Println("getting started application function...")
				function.Run(&c, ch)
			}
		} else {
			logrus.Warningln("unknown command")
		}
	}
}
