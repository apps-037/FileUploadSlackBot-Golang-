// Step one is defining the package main
package main

//Step 2 is importing libraries required
import (
	"fmt" //for printing
	"os"  //get set the env

	"github.com/slack-go/slack" //external package for slackbot for slack apis
)

// Step 3 is creating a func
func main() {
	fmt.Printf("hi")

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5262879010003-5269716773379-o48nS6NCKUO7lljMpNnlAkjK")
	os.Setenv("CHANNEL_ID", "C057TES004C")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN")) // creating a new connection calling it slackbot token
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"postman_collection_v1.1.json", "database-export.csv"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}

}
