package dapr

import (
	"fmt"
	"strings"
	"time"

	"github.com/dapr/components-contrib/bindings"
	"github.com/dapr/components-contrib/bindings/twitter"
	dapr_logger "github.com/dapr/kit/logger"
)

type DaprInputBindingConnector struct {
	name    string
	binding bindings.InputBinding
}

func NewDaprInputBindingConnector(name string, logger dapr_logger.Logger) (*DaprInputBindingConnector, error) {
	componentNameParts := strings.SplitN(name, "/", 2)
	if len(componentNameParts) != 2 {
		return nil, fmt.Errorf("unknown dapr input binding '%s'", name)
	}

	componentName := componentNameParts[1]

	var binding bindings.InputBinding
	switch componentName {
	case "twitter":
		binding = twitter.NewTwitter(logger)
	}

	if binding == nil {
		return nil, fmt.Errorf("unknown dapr input binding '%s'", componentName)
	}

	return &DaprInputBindingConnector{
		name:    componentName,
		binding: binding,
	}, nil
}

func (c *DaprInputBindingConnector) Init(Epoch time.Time, Period time.Duration, Interval time.Duration, params map[string]string) error {
	metadata := bindings.Metadata{
		Name:       c.name,
		Properties: params,
	}
	return c.binding.Init(metadata)
}

func (c *DaprInputBindingConnector) Read(handler func(data []byte, metadata map[string]string) ([]byte, error)) error {
	daprHandler := func(response *bindings.ReadResponse) ([]byte, error) {
		return handler(response.Data, response.Metadata)
	}
	return c.binding.Read(daprHandler)
}