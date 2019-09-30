package dict_test

import (
	"github.com/danielpaulus/quicktime_video_hack/usb/dict"
	"github.com/stretchr/testify/assert"
	"testing"
)

//I took these from hexdumps
var typeSix = []byte{0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9E, 0x40}
var typeFour = []byte{0x04, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var typeThree = []byte{0x03, 0x1E, 00, 00, 00}

const typeSixDecoded float64 = 1920
const typeThreeDecoded uint32 = 30
const typeFourDecoded uint64 = 5

func TestErrors(t *testing.T) {
	typeSix[0] = 3
	_, err := dict.NewNSNumber(typeSix)
	assert.Error(t, err)

	typeThree[0] = 6
	_, err = dict.NewNSNumber(typeThree)
	assert.Error(t, err)

	typeThree[0] = 4
	_, err = dict.NewNSNumber(typeThree)
	assert.Error(t, err)

	typeThree[0] = 56
	_, err = dict.NewNSNumber(typeThree)
	assert.Error(t, err)
}

func TestNumberValue(t *testing.T) {

	float64Num, err := dict.NewNSNumber(typeSix)
	if assert.NoError(t, err) {
		assert.Equal(t, typeSixDecoded, float64Num.FloatValue)
	}

	uint64Num, err := dict.NewNSNumber(typeFour)
	if assert.NoError(t, err) {
		assert.Equal(t, typeFourDecoded, uint64Num.LongValue)
	}

	uint32Num, err := dict.NewNSNumber(typeThree)
	if assert.NoError(t, err) {
		assert.Equal(t, typeThreeDecoded, uint32Num.IntValue)
	}
}
