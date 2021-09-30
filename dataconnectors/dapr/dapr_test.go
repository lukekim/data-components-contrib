package dapr_test

import (
	"fmt"
	"testing"

	dapr_logger "github.com/dapr/kit/logger"
	"github.com/spiceai/data-components-contrib/dataconnectors/dapr"
	"github.com/stretchr/testify/assert"
)

func TestDaprInputBindingConnector(t *testing.T) {
	componentsToTest := []string{
		"dapr/twitter",
	}

	for _, name := range componentsToTest {
		t.Run(fmt.Sprintf("NewDaprInputBindingConnector() - %s", name), testDaprInputBindingConnectorFunc(name))
	}
}

func testDaprInputBindingConnectorFunc(name string) func(*testing.T) {
	logger := dapr_logger.NewLogger("test-logger")
	return func(t *testing.T) {
		c, err := dapr.NewDaprInputBindingConnector(name, logger)
		assert.NoError(t, err)
		assert.NotNil(t, c)
	}
}