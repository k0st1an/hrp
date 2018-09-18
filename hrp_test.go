package hrp_test

import (
	"testing"

	"github.com/k0st1an/hrp"

	"github.com/stretchr/testify/assert"
)

var samples = []struct {
	sample       string
	expValue     string
	expValueType string
	expError     bool
}{
	{sample: "", expValue: "", expValueType: "", expError: true},
	{sample: "2d", expValue: "2", expValueType: "d", expError: false},
	{sample: "120m", expValue: "120", expValueType: "m", expError: false},
	{sample: "640", expValue: "640", expValueType: "", expError: false},
	{sample: "1", expValue: "1", expValueType: "", expError: false},
	{sample: "d", expValue: "", expValueType: "", expError: true},
	{sample: "20k wps", expValue: "20000", expValueType: "wps", expError: false},
	{sample: "20 wps", expValue: "20", expValueType: "wps", expError: false},
	{sample: "42.1k rps", expValue: "", expValueType: "", expError: true},
}

func TestConvert(t *testing.T) {
	for _, item := range samples {
		value, vType, err := hrp.Convert(item.sample)

		assert.Equal(t, item.expValueType, vType)
		assert.Equal(t, item.expValue, value)

		if item.expError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
