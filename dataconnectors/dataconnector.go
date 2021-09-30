package dataconnectors

import (
	"fmt"
	"strings"
	"time"

	"github.com/spiceai/data-components-contrib/dataconnectors/dapr"
	"github.com/spiceai/data-components-contrib/dataconnectors/file"
	"github.com/spiceai/data-components-contrib/dataconnectors/influxdb"

	dapr_logger "github.com/dapr/kit/logger"
)

var (
	// TODO: Replace with Spice AI daprComponentLogger
	daprComponentLogger dapr_logger.Logger = dapr_logger.NewLogger("dapr-data-connectors-logger")
)

type DataConnector interface {
	Init(Epoch time.Time, Period time.Duration, Interval time.Duration, params map[string]string) error
	Read(handler func(data []byte, metadata map[string]string) ([]byte, error)) error
}

func NewDataConnector(name string) (DataConnector, error) {
	if strings.HasPrefix(name, "dapr/") {
		return dapr.NewDaprInputBindingConnector(name, daprComponentLogger)
	}
	
	switch name {
	case file.FileConnectorName:
		return file.NewFileConnector(), nil
	case influxdb.InfluxDbConnectorName:
		return influxdb.NewInfluxDbConnector(), nil
	}

	return nil, fmt.Errorf("unknown data connector '%s'", name)
}
