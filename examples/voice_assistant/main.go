package main

import (
	"log/slog"
	"time"

	"github.com/mycontroller-org/esphome_api/pkg/api"
	"github.com/mycontroller-org/esphome_api/pkg/client"
)

func main() {
	if err := run(); err != nil {
		slog.Error("something went wrong", err)
		return
	}
	// block
	select {}
}

func run() error {
	port := 12346
	ul, err := NewUDPListerner(port)
	if err != nil {
		return err
	}
	stt := NewSttApiClient("http://localhost:8000")
	eh := NewEventHandler(ul, stt, port)
	c, err := client.New("testis", "192.168.2.34:6053", "", 10*time.Second, eh.Handle)
	if err != nil {
		return err
	}
	if err := c.Login(""); err != nil {
		return err
	}
	return c.Send(&api.SubscribeVoiceAssistantRequest{Subscribe: true, Flags: 0})
}
