package sender

import (
	"github.com/DarthPestilane/easytcp"
	"github.com/labstack/gommon/log"
	"github.com/valerius21/scap/pkg/common"
	"net"
	"time"
)

// Sender is an interface for sending messages to the webservers
type Sender struct {
	C      *net.Conn
	Packer *easytcp.DefaultPacker
}

// CreateSender creates a tcp connection to the specified host and port
func CreateSender(host, port string) Sender {
	c, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	packer := easytcp.NewDefaultPacker()
	return Sender{
		C:      &c,
		Packer: packer,
	}
}

// Send sends a message to the receiver and waits for a response
func (s *Sender) Send(msg string, respond func() interface{}) (int, string, error) {
	conn := *s.C
	packer := s.Packer
	go func() {
		// write loop
		for {
			time.Sleep(time.Second)
			msg := easytcp.NewMessage(common.MsgIdPingReq, []byte("ping, ping, ping"))
			packedBytes, err := packer.Pack(msg)
			if err != nil {
				panic(err)
			}
			if _, err := conn.Write(packedBytes); err != nil {
				panic(err)
			}
			log.Infof("snd >>> | id:(%d) size:(%d) data: %s", msg.ID(), len(msg.Data()), msg.Data())
			//break
		}
	}()
	go func() {
		// read loop
		for {
			msg, err := packer.Unpack(conn)
			if err != nil {
				panic(err)
			}
			log.Infof("rec <<< | id:(%d) size:(%d) data: %s", msg.ID(), len(msg.Data()), msg.Data())
			//respond()
			//break
		}
	}()
	select {}
}
