package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/scgolang/osc"
)

type Config struct {
	Addr string
}

func main() {
	var config Config
	flag.StringVar(&config.Addr, "l", "0.0.0.0:57120", "listening address")
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", config.Addr)
	if err != nil {
		panic(err)
	}
	srv, err := osc.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	app := &App{Config: config}

	fmt.Printf("listening for OSC data on %s\n", config.Addr)

	if err := srv.Serve(1, app); err != nil {
		panic(err)
	}
}

type App struct {
	Config
}

func (app *App) Dispatch(bundle osc.Bundle, exactMatch bool) error {
	return nil
}

func (app *App) Invoke(msg osc.Message, exactMatch bool) error {
	args := make([]any, len(msg.Arguments)+1)
	args[0] = msg.Address
	formatString := "%s"
	for i, arg := range msg.Arguments {
		args[i+1] = arg
		formatString += " %s"
	}

	fmt.Printf(formatString+"\n", args...)

	return nil
}
