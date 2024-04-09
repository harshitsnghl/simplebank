package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// This is done because gin logs are difficult to read because it is in debug mode by default, in TestMode logs are less
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
