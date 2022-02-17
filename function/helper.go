package function

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/novriyantoAli/go-auto-view/model"
	"github.com/sirupsen/logrus"

	cp "github.com/otiai10/copy"
)

func writeCredentials(username *string, password *string) {
	data := model.Credentials{
		Username: username,
		Password: password,
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		logrus.Error(err)
	}

	err = ioutil.WriteFile("credentials.json", file, 0644)
	if err != nil {
		logrus.Error(err)
	}
}

func regenerateCache(destination string, fullname string) {

	logrus.Info("remove profile in system chrome config directory...")
	logrus.Info("profile system directory: ", (destination + fullname))
	err := os.RemoveAll((destination + fullname))
	if err != nil {
		logrus.Error(err)
	}

	// get working directory
	logrus.Info("get working directory...")
	path, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("copy directory to destination...")
	err = cp.Copy((path + "/" + fullname), (destination + fullname))
	if err != nil {
		logrus.Error(err)
	}
}

func Decrypt(encryptedString string, k string) string {
	key, err := hex.DecodeString(k)
	if err != nil {
		panic(err.Error())
	}

	enc, err := hex.DecodeString(encryptedString)
	if err != nil {
		panic(err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphetext := enc[:nonceSize], enc[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, ciphetext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plainText)
}

// func removeProcessIndex(s []model.Process, index int) []model.Process {
// 	ret := make([]model.Process, 0)
// 	ret = append(ret, s[:index]...)
// 	return append(ret, s[index+1:]...)
// }
