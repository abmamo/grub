package main

import (
	"context"
	"log"
	"strconv"

	"github.com/shomali11/slacker"
)

func InitSlack() {
	defer wg.Done()
	botToken := getEnvironment("BOT_TOKEN", ".env")
	bot := slacker.NewClient(botToken)

	definition := &slacker.CommandDefinition{
		Description: "Repeat a word a number of times!",
		Example:     "repeat hello 10",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := request.StringParam("word", "Hello!")
			number := request.IntegerParam("number", 1)
			for i := 0; i < number; i++ {
				response.Reply(word)
			}
		},
	}

	definitionTwo := &slacker.CommandDefinition{
		Description: "Multiply two numbers",
		Example:     "multiply 2 5",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			numberOne := request.IntegerParam("numberOne", 1)
			numberTwo := request.IntegerParam("numberTwo", 1)
			response.Reply(strconv.Itoa(numberOne * numberTwo))
		},
	}

	bot.Command("repeat <word> <number>", definition)

	bot.Command("multiply <numberOne> <numberTwo>", definitionTwo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
