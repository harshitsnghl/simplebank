package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/harshitsnghl/simplebank/db/sqlc"
	"github.com/harshitsnghl/simplebank/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	// This is done because gin logs are difficult to read because it is in debug mode by default, in TestMode logs are less
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
