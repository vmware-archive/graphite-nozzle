package metrics

import (
	"errors"
)

func CreateClient(conf map[string]string) (client StatsdClient, err error) {

	switch protocol, _ := conf["protocol"]; protocol {
	case "udp":
		client, err = NewStatsdUDPClient(conf)
	case "tcp":
		client, err = NewStatsdTCPClient(conf)
	default:
		return nil, errors.New("Invalid Client name. Must be one of: 'udp', 'tcp'")
	}
	// Run the factory with the configuration.
	return client, err
}
