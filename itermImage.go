package itermImage

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

// print inline image
// file includes full path
func Print(file string) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading image %s. Error: %s", file, err.Error())
	}
	b64FileName := base64.StdEncoding.EncodeToString([]byte(file))
	b64FileContents := base64.StdEncoding.EncodeToString(body)
	fmt.Printf("\033]1337;File=name=%s;inline=1:%s\a\n", b64FileName, b64FileContents)
}
