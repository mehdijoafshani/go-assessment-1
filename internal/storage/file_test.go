package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileName(t *testing.T) {
	initTestEnv()

	dir := "/some/path"
	id := 0
	expectedFileName := "/some/path/0" + config.Data.BalanceFileExtension
	actualFileName := fileName(dir, id)
	assert.Equal(t, expectedFileName, actualFileName, "/some/path", "0", config.Data.BalanceFileExtension)

	dir = "/some/path/"
	id = 0
	expectedFileName = "/some/path/0" + config.Data.BalanceFileExtension
	actualFileName = fileName(dir, id)
	assert.Equal(t, expectedFileName, actualFileName, "/some/path/", "0", config.Data.BalanceFileExtension)

	dir = "/some/path///"
	id = 0
	expectedFileName = "/some/path/0" + config.Data.BalanceFileExtension
	actualFileName = fileName(dir, id)
	assert.Equal(t, expectedFileName, actualFileName, "/some/path///", "0", config.Data.BalanceFileExtension)

}
