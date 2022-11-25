package lib

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/go-ping/ping"
)

// checks if the pc is connected
func isConnected() error {
	// does the magic setup for the ping
	// the log.Println and return are there so i can retry
	pinger, err := ping.NewPinger("bing.com")
	if err != nil {
		return err
	}
	pinger.Count = 2 // just to be sure
	// Required for windows
	pinger.SetPrivileged(true)
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return err
	}
	return nil
	//stats := pinger.Statistics()
}

// the actual ping thing
func Pong(maxRetries int, pingSleep string) error {
	pingSleepDuration, err := time.ParseDuration(pingSleep)
	if err != nil {
		return err
	}
	// checks if the pc is connected
	for i := 0; i < maxRetries; i++ {
		err := isConnected()
		// if not, print "ping failed" and sleep for pingSleep seconds
		if err != nil {
			log.Println(err)
			log.Println("ping failed, waiting", pingSleepDuration.String(), "seconds")
			time.Sleep(pingSleepDuration)
		} else {
			// if yes, return
			return nil
		}
	}
	// if it still fails, panik
	return errors.New("couldn't connect after " + strconv.Itoa(maxRetries) + " tries")
}
