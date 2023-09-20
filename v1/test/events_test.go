package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/netoneko/trains/v1/types"
	"github.com/stretchr/testify/require"
)

func Test_GenerateRouteTransitions(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	stations := []types.StationID{
		"Bialik", "Abba Hillel", "Arlozoroff", "Sha'ul HaMelech", "Yehudit", "Karlebach",
	}
	r3 := types.NewRoute("R3", stations)

	err, ch := types.GenerateRouteEvents(ctx, r3, 3*time.Millisecond)
	require.NoError(t, err)

	var events []types.RouteEvent
	for loop := true; loop; {
		select {
		case e, ok := <-ch:
			if !ok {
				loop = false
				break
			}
			println("added event")
			events = append(events, e)
		case <-ctx.Done():
			println("done!")
			loop = false
			break
		}
	}

	fmt.Println(events)

	require.NotEmpty(t, events)
	require.EqualValues(t, types.RouteEvent{
		Status:  types.EnRoute,
		Station: "Bialik",
	}, events[0])
	require.EqualValues(t, types.RouteEvent{
		Status:  types.Arriving,
		Station: "Bialik",
	}, events[1])
	require.EqualValues(t, types.RouteEvent{
		Status:  types.Waiting,
		Station: "Bialik",
	}, events[2])
	require.EqualValues(t, types.RouteEvent{
		Status:  types.Departing,
		Station: "Bialik",
	}, events[3])
	require.EqualValues(t, types.RouteEvent{
		Status:  types.EnRoute,
		Station: "Abba Hillel",
	}, events[4])
}
