package types

import (
	"context"
	"errors"
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
	GetCurrentStation(context.Context) (error, StationID)
	GetRoute(context.Context) (error, Route)
	StartRoute(context.Context) error
	GetStatus() TrainStatus
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
	id    TrainID
	route Route

	status  TrainStatus
	station StationID

	routeContext       context.Context
	routeContextCancel context.CancelFunc
}

func NewTrain(id TrainID, route Route) Train {
	return &TrainImpl{
		id:    id,
		route: route,
	}
}

func (t *TrainImpl) GetRoute(context.Context) (error, Route) {
	return nil, t.route
}

func (t *TrainImpl) GetCurrentStation(context.Context) (error, StationID) {
	if t.station == "" {
		return errors.New("unknown location"), ""
	}

	switch t.status {
	case Stopped:
		return nil, t.station
	}

	return errors.New("not implemented"), ""
}

func (t *TrainImpl) StartRoute(ctx context.Context) error {
	t.routeContext, t.routeContextCancel = context.WithCancel(ctx)

	go func(routeContext context.Context, routeContextCancel context.CancelFunc) {

	}(t.routeContext, t.routeContextCancel)

	// return errors.New("not implemented")
	return nil
}

func (t *TrainImpl) GetStatus() TrainStatus {
	return t.status
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
