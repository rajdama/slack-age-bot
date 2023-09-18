package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

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
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5888766340278-5903175195333-Tz0Bq48y9ckYKxUS2A6C17jt")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05RWRT6PB9-5906042780307-0ab34007a1d58138d3ea5175124ba6083601196f17311f277802713f410bdfc7")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			year := r.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println(err)
			}
			age := 2023 - yob
			res := fmt.Sprintf("age is %d", age)
			w.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
