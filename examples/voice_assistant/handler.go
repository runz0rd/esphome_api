package main

import (
	"log/slog"

	"github.com/davecgh/go-spew/spew"
	"github.com/mycontroller-org/esphome_api/pkg/api"
	"github.com/mycontroller-org/esphome_api/pkg/client"
	"google.golang.org/protobuf/proto"
)

type EventHandler struct {
	li   Listener
	stt  Stt
	port int
}

func NewEventHandler(li Listener, stt Stt, port int) *EventHandler {
	return &EventHandler{li, stt, port}
}

func (eh EventHandler) Handle(c *client.Client, msg proto.Message) {
	switch msg.(type) {
	case *api.VoiceAssistantRequest:
		if msg.(*api.VoiceAssistantRequest).Start {
			slog.Info("VoiceAssistantRequest.Start")
			go func() {
				c.Send(&api.VoiceAssistantEventResponse{EventType: api.VoiceAssistantEvent_VOICE_ASSISTANT_INTENT_START})
				if err := eh.run(); err != nil {
					slog.Error("pipeline failed", err)
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
	// classify text to command
	// run command
	return nil
}
