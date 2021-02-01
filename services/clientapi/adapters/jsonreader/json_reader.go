package jsonreader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"ndanamedtt/services/clientapi/domain"
)

var ErrInvalidStream = errors.New("invalid opening token")

type JsonObjectStreamer struct {
	dec *json.Decoder
}

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func NewJsonObjectStreamer(r io.Reader) (JsonObjectStreamer, error) {
	handleErr := func(err error) (JsonObjectStreamer, error) {
		return JsonObjectStreamer{}, fmt.Errorf("NewJsonObjectStreamer: %w", err)
	}
	dec := json.NewDecoder(r)
	t, err := dec.Token()
	if err != nil {
		return handleErr(err)
	}
	if d, ok := t.(json.Delim); !ok || d.String() != "{" {
		return handleErr(ErrInvalidStream)
	}
	return JsonObjectStreamer{dec: dec}, nil
}

func (stream JsonObjectStreamer) GetStream(ctx context.Context) <-chan domain.Port {
	logErr := func(err error) {
		log.Println(fmt.Errorf("JsonObjectStreamer.GetStream error : %w", err))
	}
	ch := make(chan domain.Port)
	go func() {
		var p domain.Port
		defer close(ch)
		for stream.dec.More() {
			//read key. i assume that we don't quite need key itself, since it already exists in Port data
			if _, err := stream.dec.Token(); err != nil {
				logErr(fmt.Errorf("read toke: %w", err))
				continue
			}
			if err := stream.dec.Decode(&p); err != nil {
				logErr(fmt.Errorf("decode : %w", err))
				continue
			}
			select {
			case ch <- p:
				continue
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}
