package twitter_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	dapr_logger "github.com/dapr/kit/logger"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spiceai/data-components-contrib/dataconnectors/dapr"
	"github.com/stretchr/testify/assert"
)

func TestDaprTwitterConnector(t *testing.T) {
	logger := dapr_logger.NewLogger("test-logger")
	logger.SetOutputLevel(dapr_logger.DebugLevel)
	
	params := map[string]string{
		"consumerKey": "",
		"consumerSecret": "",
		"accessToken": "",
		"accessSecret": "",
	}

	t.Run("Init()", testInitFunc(logger, params))
	t.Run("Read()", testReadFunc(logger, params))
}

func testInitFunc(logger dapr_logger.Logger, params map[string]string) func(*testing.T) {
	return func(t *testing.T) {
		c, err := dapr.NewDaprInputBindingConnector("dapr/twitter", logger)
		assert.NoError(t, err)
		assert.NotNil(t, c)

		var epoch time.Time
		period := 7 * 24 * time.Hour
		interval := time.Hour

		err = c.Init(epoch, period, interval, params)
		assert.NoError(t, err)
	}
}

func testReadFunc(logger dapr_logger.Logger, params map[string]string) func(*testing.T) {
	return func(t *testing.T) {
		c, err := dapr.NewDaprInputBindingConnector("dapr/twitter", logger)
		assert.NoError(t, err)
		assert.NotNil(t, c)

		params["query"] = "#spiceai"
		var epoch time.Time
		period := 7 * 24 * time.Hour
		interval := time.Hour

		err = c.Init(epoch, period, interval, params)
		assert.NoError(t, err)

		err = c.Read(func(data []byte, metadata map[string]string) ([]byte, error) {
			fmt.Println("READ")
			var tweet twitter.Tweet
			err := json.Unmarshal(data, &tweet)
			if err != nil {
				return nil, err
			}
			assert.NotNil(t, tweet)
			fmt.Printf("%+v", tweet)

			return nil, nil
		})
		assert.NoError(t, err)
	}
}