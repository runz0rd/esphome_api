package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

type Client struct {
	*spotify.Client
	deviceName           string
	likedSongsContextUri string
}

func NewClient(deviceName, contextUri string) (*Client, error) {
	f, err := os.Open("token.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var token oauth2.Token
	json.NewDecoder(f).Decode(&token)
	return &Client{
		spotify.New(auth.Client(context.Background(), &token)),
		deviceName,
		contextUri,
	}, nil
}

func (c Client) PlayLiked(ctx context.Context) error {
	deviceId, err := c.getDeviceId(ctx, c.deviceName)
	if err != nil {
		return err
	}
	contextUri := spotify.URI(c.likedSongsContextUri)
	return c.PlayOpt(ctx, &spotify.PlayOptions{
		DeviceID:        &deviceId,
		PlaybackContext: &contextUri,
	})
}

func (c Client) getDeviceId(ctx context.Context, name string) (spotify.ID, error) {
	ds, err := c.PlayerDevices(ctx)
	if err != nil {
		return "", err
	}
	for _, d := range ds {
		if d.Name == name {
			return d.ID, nil
		}
	}
	return "", fmt.Errorf("no device named %q found", name)
}
