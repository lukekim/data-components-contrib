package dataconnectors

import (
	"fmt"
	"testing"

	"github.com/spiceai/data-components-contrib/dataconnectors/file"
	"github.com/spiceai/data-components-contrib/dataconnectors/influxdb"
	"github.com/stretchr/testify/assert"
)

func TestNewDataConnector(t *testing.T) {
	t.Run("NewDataConnector() - Invalid connector", testNewDataConnectorUnknownFunc())

	connectorNamesToTest := []string{
		"dapr/twitter",
		file.FileConnectorName,
		influxdb.InfluxDbConnectorName,
	}

	for _, name := range connectorNamesToTest {
		t.Run(fmt.Sprintf("NewDataConnector() - %s", name), testNewDataConnectorFunc(name))
	}
}

func testNewDataConnectorUnknownFunc() func(*testing.T) {
	return func(t *testing.T) {
		_, err := NewDataConnector("does-not-exist")
		assert.Error(t, err)
	}
}

func testNewDataConnectorFunc(name string) func(*testing.T) {
	return func(t *testing.T) {
		c, err := NewDataConnector(name)
		assert.NoError(t, err)
		assert.NotNil(t, c)
	}
}