package main

import (
	"database/sql"
	"log"
	"order/application/command"
	eventhandler "order/application/event_handler"
	"order/infra/event"
	"order/infra/repository"
	"order/infra/smtp"
	"order/infra/web"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://rabbit:rabbit_password@localhost:5672/")
	if err != nil {
		log.Fatal("a", err)
		return
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
		return
	}
	db, err := sql.Open("sqlite3", "./order.sql")
	if err != nil {
		log.Fatal(err)
		return
	}
	smtp := smtp.NewOrderSMTP(
		os.Getenv("EMAIL"),
		os.Getenv("KEY"),
	)
	r := gin.Default()

	eventBus, err := event.NewEventBus(channel, "events")
	if err != nil {
		log.Fatal(err)
		return
	}
	repository := repository.NewSQLRepository(db)

	placeOrder := command.NewPlaceOrder(eventBus)

	sendEmail := eventhandler.NewSendEmail(smtp, eventBus)
	saveOrder := eventhandler.NewSaveOrder(repository)

	eventBus.Subscribe("OrderPlaced", sendEmail)
	eventBus.Subscribe("EmailSent", saveOrder)

	controller := web.NewOrderController(
		*placeOrder,
	)
	web.InitRoutes(&r.RouterGroup, controller)
	r.Run(":8080")
}
