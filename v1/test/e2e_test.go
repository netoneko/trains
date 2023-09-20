package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/netoneko/trains/v1/types"
	"github.com/stretchr/testify/require"
)

func Test_E2E(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	stations := []types.StationID{
		"Bialik", "Abba Hillel", "Arlozoroff", "Sha'ul HaMelech", "Yehudit", "Karlebach",
	}
	r3 := types.NewRoute("R3", stations)

	err, stationList := r3.GetStationList(ctx)

	require.NoError(t, err)
	require.EqualValues(t, "R3", r3.GetID(ctx))
	require.EqualValues(t, stations, stationList)

	train := types.NewTrain("R3_01", r3)
	require.EqualValues(t, "R3_01", train.GetID())

	err, trainRoute := train.GetRoute(ctx)
	require.NoError(t, err)
	require.EqualValues(t, r3, trainRoute)

	err, events := train.GetEvents(ctx)
	require.NoError(t, err)
	require.Len(t, events, 1)

	err, lastEvent := train.GetLastEvent(ctx)
	require.NoError(t, err)
	require.EqualValues(t, types.RouteEvent{
		Status:  types.Stopped,
		Station: "",
	}, lastEvent)

	// start the route

	err = train.StartRoute(ctx, GenerateRouteEvents)
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	err, events = train.GetEvents(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, events)

	fmt.Println(events)

	err, lastEvent = train.GetLastEvent(ctx)
	require.NoError(t, err)
	require.EqualValues(t, types.RouteEvent{
		Status:  types.Departing,
		Station: "Karlebach",
	}, lastEvent)
}
