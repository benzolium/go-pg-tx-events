package txevents

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_PORT"),
		viper.GetString("POSTGRES_DB_NAME"),
	)
	minReconnectInterval := time.Second
	maxReconnectInterval := time.Minute
	listener := pq.NewListener(
		dsn,
		minReconnectInterval,
		maxReconnectInterval,
		func(event pq.ListenerEventType, err error) {
			if err != nil {
				panic(err)
			}
		},
	)
	err := listener.Listen("events")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case n := <-listener.Notify:
			if n == nil {
				continue
			}
			// TODO: emit with amqp here
			fmt.Println(n.Extra)
		}
	}

}
