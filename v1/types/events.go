package types

import (
	"context"
	"time"
)

type RouteEvent struct {
	Status  TrainStatus
	Station StationID
}

func GenerateEvents(ctx context.Context, route Route, period time.Duration) (error, chan (RouteEvent)) {
	ch := make(chan RouteEvent)

	return nil, ch
}
