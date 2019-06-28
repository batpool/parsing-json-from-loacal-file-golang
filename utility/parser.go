package utility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"play/model"
)

// ReadJson function
func ReadJson(data string) model.Users {
	jsonFile, err := os.Open(data)
	if err != nil {
		fmt.Print(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users model.Users

	json.Unmarshal(byteValue, &users)

	defer jsonFile.Close()

	return users
}
