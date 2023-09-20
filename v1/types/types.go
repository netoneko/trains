package types

import (
	"context"
)

type StationID string
type StationList []StationID

type Route interface {
	GetStationList(context.Context) (error, StationList)
}

type RouteImpl struct {
	stationList StationList
}

func (r *RouteImpl) GetStationList(context.Context) (error, StationList) {
	return nil, r.stationList
}

func NewRoute(stationList StationList) Route {
	return &RouteImpl{
		stationList: stationList,
	}
}

type Train interface {
	GetCurrentStation(context.Context) (error, StationID)
	GetLastStation(context.Context) (error, StationID)
	GetRoute(context.Context) (error, Route)
}

type SensorPingResponse interface {
	GetStationID() StationID
}

type Sensor interface {
	Ping(context.Context) (error, SensorPingResponse)
}
