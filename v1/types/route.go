package types

import (
	"context"
)

type RouteID string
type StationList []StationID

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

type RouteEvent struct {
	Status  TrainStatus
	Station StationID
}

type SensorPingResponse interface {
	GetStationID() StationID
}

type Sensor interface {
	Ping(context.Context) (error, SensorPingResponse)
}
