package fbrip

import (
    "testing"
	"github.com/stretchr/testify/assert"
)


func TestTransformUrlToBasicFacebook(t *testing.T) {
	facebookUrl := *FacebookUrl
	assert.Equal(t, "www.facebook.com", facebookUrl.Host)
	transformUrlToBasicFacebook(&facebookUrl)
	assert.Equal(t, "mbasic.facebook.com", facebookUrl.Host)
}

func TestTransformUrlToFacebook(t *testing.T) {
	basicFacebookUrl := *BasicFacebookUrl
	assert.Equal(t, "mbasic.facebook.com", basicFacebookUrl.Host)
	transformUrlToFacebook(&basicFacebookUrl)
	assert.Equal(t, "www.facebook.com", basicFacebookUrl.Host)
}

func TestParseUrls(t *testing.T){
	rawUrls := []string{
		"www.facebook.com",
		"www.facebook.com/this-is-a-random-path",
		"www.facebook.com/this-is-a-random-path/?random=query-parameter",
		"mbasic.facebook.com",
		"mbasic.facebook.com/this-is-a-random-path",
		"mbasic.facebook.com/this-is-a-random-path/?random=query-parameter",
	} 
	parsedUrls := parseUrls(rawUrls)
	assert.Equal(t, len(rawUrls), len(parsedUrls))
}
