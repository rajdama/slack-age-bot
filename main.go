package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)

	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5888766340278-5903175195333-gDniPQ7dzpgCq87A7sagAh9w")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05RWRT6PB9-5929085276576-dd626e225abd9133554c39f03398a23c54428ba5b0d0f6a98d0d31b756440cfb")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
