package types

import (
	"context"
	"time"
)

type StationID string
type StationList []StationID

type RouteID string

type Route interface {
	GetID(context.Context) RouteID
	GetStationList(context.Context) (error, StationList)
}

type RouteImpl struct {
	id          RouteID
	stationList StationList
}

func (r *RouteImpl) GetStationList(context.Context) (error, StationList) {
	return nil, r.stationList
}

func (r *RouteImpl) GetID(context.Context) RouteID {
	return r.id
}

func NewRoute(id RouteID, stationList StationList) Route {
	return &RouteImpl{
		id:          id,
		stationList: stationList,
	}
}

type TrainID string

type Train interface {
	GetID() TrainID
	GetEvents(context.Context) (error, []RouteEvent)
	GetLastEvent(context.Context) (error, RouteEvent)
	GetRoute(context.Context) (error, Route)
	StartRoute(context.Context) error
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

func (t *TrainImpl) StartRoute(ctx context.Context) error {
	t.routeContext, t.routeContextCancel = context.WithCancel(ctx)

	// FIXME temporary solution
	err, ch := GenerateRouteEvents(ctx, t.route, 1*time.Millisecond)
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

type SensorPingResponse interface {
	GetStationID() StationID
}

type Sensor interface {
	Ping(context.Context) (error, SensorPingResponse)
}
