package lib

import (
	"errors"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-ping/ping"
)

// Does the actual ping, returns error if ping fails
func isConnected(timeout time.Duration) error {
	// Pings 1.1.1.1
	pinger, err := ping.NewPinger("1.1.1.1")
	if err != nil {
		return err
	}
	// Listen for Ctrl-C
	// Else if the ping gets stuck the whole application gets stuck
	// This is like real janky tho
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
			os.Exit(1)
		}
	}()

	// Config
	pinger.Count = 2
	pinger.Timeout = timeout
	// Required for windows, but may be a problem for linux
	pinger.SetPrivileged(true)

	err = pinger.Run() // Blocks until finished, that's why the ctrl-c listener is needed
	if err != nil {
		return err
	} else {
		return nil
	}

}

// The ping ""interface"" thing
func Pong(maxRetries int, pingSleep string, timeout string) error {
	// Parse sleep time from the string
	pingSleepDuration, err := time.ParseDuration(pingSleep)
	if err != nil {
		return err
	}

	// Parse timeout duration from the string
	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		return err
	}

	// Tries for maxRetries to check for connection, if it can connect, it will return nil, else "couldn't connect after x tries"
	for i := 0; i < maxRetries; i++ {
		err = isConnected(timeoutDuration)
		// if not, print "ping failed" and sleep for pingSleep seconds
		if err != nil {
			LogInColor.Error(err)
			LogInColor.Warn("ping failed, waiting", pingSleepDuration.String(), "seconds")
			time.Sleep(pingSleepDuration)
		} else {
			return nil
		}
	}
	// if it still fails, throw error
	return errors.New("Couldn't connect after " + strconv.Itoa(maxRetries) + " tries")
}
