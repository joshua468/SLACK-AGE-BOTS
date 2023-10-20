package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomalli11/slacker"
)

func printCommandEvents(analyticsChannel chan<- *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("CommandEvents")
		fmt.Println(event.Command)
		fmt.Println(event.TimeStamp)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main() {
	os.Setenv("BOT_SLACK_TOKEN", "")
	os.Setenv("APP_SLACK_TOKEN", "")

	bot := slacker.NewClient(os.Getenv("BOT_SLACK_TOKEN"), os.Getenv("BOT_SLACK_TOKEN"))
	go printCommandEvents(bot.CommandEvents)

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2021",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param["year"]
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
