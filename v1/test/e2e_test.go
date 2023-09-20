package test

import (
	"context"
	"testing"
	"time"

	"github.com/netoneko/trains/v1/types"
	"github.com/stretchr/testify/require"
)

func Test_E2E(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	stations := []types.StationID{
		"Arlozoroff", "Abba Hillel", "Bialik", "Sha'ul HaMelech", "Yehudit", "Karlebach",
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

	err, _ = train.GetCurrentStation(ctx)
	require.EqualError(t, err, "unknown location")
	require.EqualValues(t, types.Stopped, train.GetStatus())

	// start the route

	err = train.StartRoute(ctx)
	require.NoError(t, err)
	require.EqualValues(t, types.Arriving, train.GetStatus())

	time.Sleep(20 * time.Millisecond)

	err, currentStation := train.GetCurrentStation(ctx)
	require.NoError(t, err)
	require.EqualValues(t, "Arlozoroff", currentStation)
}
