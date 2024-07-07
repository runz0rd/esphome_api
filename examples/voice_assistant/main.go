package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/mycontroller-org/esphome_api/examples/voice_assistant/spotify"
	"github.com/mycontroller-org/esphome_api/pkg/api"
	"github.com/mycontroller-org/esphome_api/pkg/client"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		slog.Error("something went wrong", "error", err)
		return
	}
	// block
	select {}
}

func run(ctx context.Context) error {
	port := 12346
	ul, err := NewUDPListerner(port)
	if err != nil {
		return err
	}
	stt := NewSttApiClient("http://localhost:8000")
	sc, err := spotify.NewClient("", "")
	if err != nil {
		return err
	}
	eh := NewEventHandler(ul, stt, port, commands{
		// TODO
		// "lights": to_main.toggle,
		// "bedroom": to_bedroom.toggle,
		// "sleep": to_bedroom.delayed_on_off,
		"spotify": func() error {
			return sc.PlayLiked(ctx)
		},
		"play": func() error {
			return sc.Play(ctx)
		},
		"pause": func() error {
			return sc.Pause(ctx)
		},
		"stop": func() error {
			return sc.Pause(ctx)
		},
		"next": func() error {
			return sc.Next(ctx)
		},
		"previous": func() error {
			return sc.Previous(ctx)
		},
	})
	c, err := client.New("test", "192.168.2.34:6053", "", 10*time.Second, eh.Handle)
	if err != nil {
		return err
	}
	if err := c.Login(""); err != nil {
		return err
	}
	return c.Send(&api.SubscribeVoiceAssistantRequest{Subscribe: true, Flags: 0})
}
