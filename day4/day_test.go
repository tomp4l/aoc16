package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	assert.Equal(t, "abxyz", checksum("aaaaa-bbb-z-y-x"))
	assert.Equal(t, "abcde", checksum("a-b-c-d-e-f-g-h"))
	assert.Equal(t, "oarel", checksum("not-a-real-room"))
}

func TestParseRoom(t *testing.T) {
	room, err := parseRoom("aaaaa-bbb-z-y-x-123[abxyz]")
	assert.NoError(t, err)
	assert.Equal(t, "aaaaa-bbb-z-y-x", room.name)
	assert.Equal(t, "abxyz", room.checksum)
	assert.Equal(t, 123, room.number)
}

func TestExample(t *testing.T) {
	example := `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`

	p1, _, err := (Day{}).Run(example)

	assert.NoError(t, err)
	assert.Equal(t, "1514", p1)
}

func TestDecrypt(t *testing.T) {
	r := room{name: "qzmt-zixmtkozy-ivhz", number: 343}

	assert.Equal(t, "very encrypted name", decrypt(r))
}
