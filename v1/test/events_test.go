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
		"Arlozoroff", "Abba Hillel", "Bialik", "Sha'ul HaMelech", "Yehudit", "Karlebach",
	}
	r3 := types.NewRoute("R3", stations)

	err, ch := types.GenerateEvents(ctx, r3, 10*time.Millisecond)
	require.NoError(t, err)

	var events []types.RouteEvent
	for loop := true; loop; {
		select {
		case e := <-ch:
			println("added event")
			events = append(events, e)
		case <-ctx.Done():
			println("done!")
			loop = false
			break
		}
	}

	fmt.Println(events)

	event0 := &types.RouteEvent{
		Status:  types.Arriving,
		Station: "Arlozoroff",
	}
	require.NotEmpty(t, events)
	require.EqualValues(t, event0, events[0])
}
