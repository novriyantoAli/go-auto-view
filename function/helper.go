package function

import (
	"crypto/aes"
	"crypto/cipher"
<<<<<<< HEAD
	"encoding/json"
	"encoding/pem"
=======
	"encoding/hex"
	"encoding/json"
>>>>>>> b90dbcc04d388a38f6aaf9e1dd7e81bae9c6050e
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/novriyantoAli/go-auto-view/model"
	"github.com/sirupsen/logrus"

	cp "github.com/otiai10/copy"
)

const (
	keyFile       = "key.apin"
	encryptedfile = "config.rein"
)

var abc = []byte("5419682441671276")

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

<<<<<<< HEAD
func rKey(filename string) ([]byte, error) {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return key, err
	}
	block, _ := pem.Decode(key)
	return block.Bytes, nil
}

func cKey() []byte {
	genkey := make([]byte, 16)
	_, err := rand.Read(genkey)
	if err != nil {
		logrus.Panicln("failed to read key: %s", err.Error())
	}

	return genkey
}

func sKey(filename string, key []byte) {
	block := &pem.Block{
		Type:  "AES KEY",
		Bytes: key,
	}
	err := ioutil.WriteFile(filename, pem.EncodeToMemory(block), 0755)
	if err != nil {
		logrus.Panicln("Failed tio save the key %s: %s", filename, err)
	}
}

func aesKey() []byte {
	file := fmt.Sprintf(keyFile)
	key, err := rKey(file)
	if err != nil {
		logrus.Infoln("Create a new AES KEY")
		key = cKey()
		sKey(file, key)
	}

	return key
}

func createCipher() cipher.Block {
	c, err := aes.NewCipher(aesKey())
	if err != nil {
		logrus.Panic(err.Error())
	}

	return c
}

func Decryption() []byte {
	bytes, err := ioutil.ReadFile(fmt.Sprintf(encryptedfile))
=======
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
>>>>>>> b90dbcc04d388a38f6aaf9e1dd7e81bae9c6050e
	if err != nil {
		panic(err.Error())
	}

<<<<<<< HEAD
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, abc)
	stream.XORKeyStream(bytes, bytes)

	return bytes
=======
	return fmt.Sprintf("%s", plainText)
>>>>>>> b90dbcc04d388a38f6aaf9e1dd7e81bae9c6050e
}

// func removeProcessIndex(s []model.Process, index int) []model.Process {
// 	ret := make([]model.Process, 0)
// 	ret = append(ret, s[:index]...)
// 	return append(ret, s[index+1:]...)
// }
