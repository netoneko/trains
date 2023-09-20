package types

import (
	"context"
	"time"
)

type RouteEvent struct {
	Status  TrainStatus
	Station StationID
}

func GenerateRouteEvents(ctx context.Context, route Route, period time.Duration) (error, chan (RouteEvent)) {
	ch := make(chan RouteEvent)

	err, stations := route.GetStationList(ctx)
	if err != nil {
		return err, nil
	}

	defaultTransitions := []TrainStatus{
		EnRoute, Arriving, Waiting, Departing,
	}

	go func() {
		for i := 0; i < len(stations); i++ {
			for _, status := range defaultTransitions {
				ch <- RouteEvent{
					Status:  status,
					Station: stations[i],
				}
				time.Sleep(period)
			}
		}
		close(ch)
	}()

	return nil, ch
}
