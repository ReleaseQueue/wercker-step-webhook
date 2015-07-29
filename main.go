package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func postData(url string, dataMap map[string]interface{}) error {

	body, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return nil
}

func getWerckerEnvVars() map[string]interface{} {

	checkEnvVar := func(key string) bool {
		if strings.HasPrefix(key, "WERCKER") {
			return true
		}
		KeysToInclude := []string{"CI", "DEPLOY"}
		for _, keyToInclude := range KeysToInclude {
			if key == keyToInclude {
				return true
			}
		}
		return false
	}

	envVarMap := make(map[string]interface{})

	for _, eVar := range os.Environ() {
		splits := strings.SplitN(eVar, "=", 2)
		key := splits[0]
		val := splits[1]
		if checkEnvVar(key) {
			envVarMap[key] = val
		}
	}
	return envVarMap
}

func main() {

	if len(os.Args) < 2 {
		panic("Missing url argument")
	}

	url := os.Args[1]

	envVarMap := getWerckerEnvVars()

	err := postData(url, envVarMap)
	if err != nil {
		panic(err)
	}

}
