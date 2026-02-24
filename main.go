package main

import (
	"database/sql"
	"log"
	"order/application/command"
	eventhandler "order/application/event_handler"
	"order/application/query"
	"order/infra/adapter"
	"order/infra/event"
	"order/infra/repository"
	"order/infra/smtp"
	"order/infra/web"
	"os"

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
	server := adapter.NewGinAdapter()

	eventBus, err := event.NewEventBus(channel, "events")
	if err != nil {
		log.Fatal(err)
		return
	}
	repository := repository.NewSQLRepository(db)

	placeOrder := command.NewPlaceOrder(eventBus)
	getOrderByID := query.NewGetOrderByID(db)

	sendEmail := eventhandler.NewSendEmail(smtp, eventBus)
	saveOrder := eventhandler.NewSaveOrder(repository)

	eventBus.Subscribe("OrderPlaced", sendEmail)
	eventBus.Subscribe("EmailSent", saveOrder)

	controller := web.NewOrderController(
		*placeOrder,
		*getOrderByID,
	)
	web.InitRoutes(server, controller)
	server.Start(":8080")
}
