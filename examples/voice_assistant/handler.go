package main

import (
	"fmt"
	"log/slog"

	"github.com/davecgh/go-spew/spew"
	"github.com/mycontroller-org/esphome_api/pkg/api"
	"github.com/mycontroller-org/esphome_api/pkg/client"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/proto"
)

type Listener interface {
	GetPort() int
	Stop() error
	Read(num int) ([]byte, error)
}

type Stt interface {
	Transcribe(data []byte) (string, error)
}

type commands map[string]func() error

type EventHandler struct {
	li   Listener
	stt  Stt
	port int
	cs   commands
}

func NewEventHandler(li Listener, stt Stt, port int, cs commands) *EventHandler {
	return &EventHandler{li, stt, port, cs}
}

func (eh EventHandler) Handle(c *client.Client, msg proto.Message) {
	switch msg.(type) {
	case *api.VoiceAssistantRequest:
		if msg.(*api.VoiceAssistantRequest).Start {
			slog.Info("VoiceAssistantRequest.Start")
			go func() {
				c.Send(&api.VoiceAssistantEventResponse{EventType: api.VoiceAssistantEvent_VOICE_ASSISTANT_INTENT_START})
				if err := eh.run(); err != nil {
					slog.Error("pipeline failed", "error", err)
					c.Send(&api.VoiceAssistantEventResponse{EventType: api.VoiceAssistantEvent_VOICE_ASSISTANT_ERROR})
					// return
				}
				c.Send(&api.VoiceAssistantEventResponse{EventType: api.VoiceAssistantEvent_VOICE_ASSISTANT_INTENT_END})
			}()
			c.Send(&api.VoiceAssistantResponse{Port: uint32(eh.port), Error: false})
		} else {
			slog.Info("VoiceAssistantRequest.Stop")
		}
	case *api.VoiceAssistantEventResponse:
		slog.Info("VoiceAssistantEventResponse." + msg.(*api.VoiceAssistantEventResponse).EventType.String())
	}
}

func (eh EventHandler) run() error {
	bs, err := eh.li.Read(100)
	if err != nil {
		return err
	}
	text, err := eh.stt.Transcribe(bs)
	if err != nil {
		return err
	}
	spew.Dump(text)
	matchedCommand := match(text, maps.Keys(eh.cs))
	cmd, ok := eh.cs[matchedCommand]
	if !ok {
		return fmt.Errorf("matched command %q not configured", matchedCommand)
	}
	if err := cmd(); err != nil {
		return fmt.Errorf("command %q failed: %w", matchedCommand, err)
	}
	return nil
}
