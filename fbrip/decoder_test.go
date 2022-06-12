package fbrip

import (
    "testing"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
    payload, err := ReadRip("../examples/rip.json")
    if err != nil {
        t.Errorf("Error while trying to read json payload")
    }
	assert.Equal(t, "mock@email.com", payload.User.Email)
	assert.Equal(t, "mock_password", payload.User.Password)
	assert.Equal(t, true, payload.Actions.GetBasicInfo)
	assert.Equal(t, "1", payload.Actions.Reactions[0].ReactionId)
	assert.Equal(t, "https://www.facebook.com/", payload.Actions.Reactions[0].PostUrl)
	assert.Equal(t, "https://www.facebook.com/", payload.Actions.Scrap.Urls[0])
	assert.Equal(t, "./scraps", payload.Actions.Scrap.OutputFolderPath)
	assert.Equal(t, "mockPrefix", payload.Actions.Scrap.NamePrefix)
}