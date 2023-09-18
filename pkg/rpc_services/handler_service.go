package rpc_services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/dto"
	"github.com/valerius21/scap/pkg/fns"
	"github.com/valerius21/scap/pkg/utils"
)

type HandlerService struct{}

func (*HandlerService) HandleMessage(msg *dto.Message, reply *dto.Message) error {
	switch msg.Name {
	case "empty":
		{
			log.Info().Msg("Empty message received")
			now := time.Now()
			fns.EmptyFn()
			ts := utils.TimeTrack(now, "empty")
			*reply = dto.Message{
				Name:     "node:" + ts.Instance,
				Duration: ts.Duration,
				Data:     "",
			}
		}
	case "math":
		{
			log.Info().Msg("Math message received")
			number, err := strconv.Atoi(msg.Data)
			if err != nil {
				log.Error().Err(err).Msg("Failed to convert string to int")
			}
			now := time.Now()
			data := fns.MathFn(number)
			ts := utils.TimeTrack(now, "math")
			*reply = dto.Message{
				Name:     "node:" + ts.Instance,
				Data:     fmt.Sprintf("%.10f", data),
				Duration: ts.Duration,
			}
		}
	case "image":
		{
			log.Info().Msg("Image message received")
			now := time.Now()

			// Read the file into a byte slice
			data, err := os.ReadFile(msg.Data)
			if err != nil {
				log.Error().Err(err).Msg("Failed to read the file")
				return err
			}

			err = fns.TransformImage(data)
			if err != nil {
				log.Error().Err(err).Msg("Failed to transform the image")
				return err
			}

			ts := utils.TimeTrack(now, "image")
			*reply = dto.Message{
				Name:     "node:" + ts.Instance,
				Data:     "",
				Duration: ts.Duration,
			}
		}
	case "sleep":
		{
			log.Info().Msg("Sleep message received")
			now := time.Now()
			fns.SleeperFn(1)
			ts := utils.TimeTrack(now, "sleep")
			*reply = dto.Message{
				Name:     "node:" + ts.Instance,
				Data:     "",
				Duration: ts.Duration,
			}
		}
	default:
		{
			return fmt.Errorf("invalid message name")
		}
	}

	return nil
}
