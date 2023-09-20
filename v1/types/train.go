package types

import (
	"context"
	"time"
)

type StationID string

type TrainID string

type Train interface {
	GetID() TrainID
	GetEvents(context.Context) (error, []RouteEvent)
	GetLastEvent(context.Context) (error, RouteEvent)
	GetRoute(context.Context) (error, Route)
	StartRoute(context.Context, GetRouteEvents) error
}

type TrainStatus int64

const (
	Stopped TrainStatus = iota
	EnRoute
	// Delayed FIXME implement later
	Arriving
	Waiting
	Departing
)

type TrainImpl struct {
	id     TrainID
	route  Route
	events []RouteEvent

	routeContext       context.Context
	routeContextCancel context.CancelFunc
}

func NewTrain(id TrainID, route Route) Train {
	return &TrainImpl{
		id:    id,
		route: route,
		events: []RouteEvent{{
			Status:  Stopped,
			Station: "",
		}},
	}
}

func (t *TrainImpl) GetRoute(context.Context) (error, Route) {
	return nil, t.route
}

func (t *TrainImpl) GetEvents(context.Context) (error, []RouteEvent) {
	return nil, t.events
}

func (t *TrainImpl) GetLastEvent(context.Context) (error, RouteEvent) {
	return nil, t.events[len(t.events)-1]
}

type GetRouteEvents func(ctx context.Context, route Route, period time.Duration) (error, chan (RouteEvent))

func (t *TrainImpl) StartRoute(ctx context.Context, getRouteEvents GetRouteEvents) error {
	t.routeContext, t.routeContextCancel = context.WithCancel(ctx)

	// FIXME temporary solution
	err, ch := getRouteEvents(ctx, t.route, 1*time.Millisecond)
	if err != nil { // FIXME should wrap here
		return err
	}

	go func(routeContext context.Context, routeContextCancel context.CancelFunc) {
		for loop := true; loop; {
			select {
			case e, ok := <-ch:
				if !ok {
					loop = false
					break
				}
				t.events = append(t.events, e)
			case <-ctx.Done():
				loop = false
				break
			}
		}
	}(t.routeContext, t.routeContextCancel)

	return nil
}

func (t *TrainImpl) GetID() TrainID {
	return t.id
}
