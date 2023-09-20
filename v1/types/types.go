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
}

type TrainImpl struct {
	id    TrainID
	route Route
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
	return errors.New("not implemented"), ""
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
