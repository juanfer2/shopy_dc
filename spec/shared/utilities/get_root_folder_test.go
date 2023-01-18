package utilities

import (
	"regexp"
	"testing"

	"github.com/juanfer2/shopy_dc_golang/src/shared/utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetRootFolder(t *testing.T) {
	assert := assert.New(t)
	folder := utilities.GetRootFolder()
	match, _ := regexp.MatchString("github.com/juanfer2/shopy_dc_golang", folder)

	assert.Equal(match, true, "they should be truthy")
}
