package main

import (
	"github.com/flweber/gurl/config"
	"github.com/flweber/gurl/db"
	"github.com/flweber/gurl/shortener"
	"github.com/flweber/gurl/web"
	"log"
	"os"
	"os/signal"
)

func main() {
	c := config.NewConfig()
	rdb := db.NewClient(c.Db.Addr, c.Db.Password, c.Db.Database)
	shorter := shortener.New(rdb)
	web.StartServer(shorter, c.Port, c.Page)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
