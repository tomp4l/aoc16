package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSupportTls(t *testing.T) {
	ip, err := parseIp("abba[mnop]qrst")
	assert.NoError(t, err)
	assert.True(t, ip.supportsTls())

	ip, err = parseIp("abcd[bddb]xyyx")
	assert.NoError(t, err)
	assert.False(t, ip.supportsTls())

	ip, err = parseIp("aaaa[qwer]tyui")
	assert.NoError(t, err)
	assert.False(t, ip.supportsTls())

	ip, err = parseIp("ioxxoj[asdfgh]zxcvbn")
	assert.NoError(t, err)
	assert.True(t, ip.supportsTls())
	t.Run("support multiple hypernets", func(t *testing.T) {
		ip, err = parseIp("ioxxoj[asdfgh]zxcvbn[abba]asdadfas")
		assert.NoError(t, err)
		assert.False(t, ip.supportsTls())
	})
}

func TestSupportSs(t *testing.T) {
	ip, err := parseIp("aba[bab]xyz")
	assert.NoError(t, err)
	assert.True(t, ip.supportsSsl())

	ip, err = parseIp("xyx[xyx]xyx")
	assert.NoError(t, err)
	assert.False(t, ip.supportsSsl())

	ip, err = parseIp("aaa[kek]eke")
	assert.NoError(t, err)
	assert.True(t, ip.supportsSsl())

	ip, err = parseIp("zazbz[bzb]cdb")
	assert.NoError(t, err)
	assert.True(t, ip.supportsSsl())
}
