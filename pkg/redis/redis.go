package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/rueidis"
)

func New(cfg *Config) (rueidis.Client, error) {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{cfg.Host + ":" + strconv.Itoa(cfg.Port)},
		Password:    cfg.Password, // no password set
		SelectDB:    cfg.Db,       // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client.Do(ctx, client.B().Ping().Build())

	client.Do(ctx, client.B().Set().Key("foo").Value("bar").Build())

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = client.Do(context.Background(), client.B().Get().Key("foo").Build()).ToString()
	if err != nil {
		return nil, err
	}

	return client, nil
}
