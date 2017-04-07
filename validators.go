package main

import (
	"errors"
)

func ValidateStatsdProtocol(statsdProtocol string) (err error) {
	err_msg := "Statsd protocol needs to be one of: ['tcp', 'udp']"
	switch statsdProtocol {
 		case "udp", "tcp":
 			err = nil
		default:
			err = errors.New(err_msg)
	}

	return err
}