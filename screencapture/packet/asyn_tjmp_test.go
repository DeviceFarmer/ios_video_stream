package packet_test

import (
	"github.com/danielpaulus/quicktime_video_hack/screencapture/packet"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

var expectedBytes = []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

func TestTjmp(t *testing.T) {
	dat, err := ioutil.ReadFile("fixtures/asyn-tjmp")
	if err != nil {
		log.Fatal(err)
	}
	tjmpPacket, err := packet.NewAsynTjmpPacketFromBytes(dat)
	if assert.NoError(t, err) {
		assert.Equal(t, uint64(0x11123bc18), tjmpPacket.ClockRef)
		assert.Equal(t, packet.AsynPacketMagic, tjmpPacket.AsyncMagic)
		assert.Equal(t, packet.TJMP, tjmpPacket.MessageType)
		assert.Equal(t, expectedBytes, tjmpPacket.Unknown)
	}
}