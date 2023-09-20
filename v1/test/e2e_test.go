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
	route := types.NewRoute("R3", stations)

	err, stationList := route.GetStationList(ctx)

	require.NoError(t, err)
	require.EqualValues(t, stations, stationList)

}
