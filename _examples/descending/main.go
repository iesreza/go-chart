package main

import (
	"net/http"

	"github.com/iesreza/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {

	/*
	   The below will draw the same chart as the `basic` example, except with both the x and y axes turned on.
	   In this case, both the x and y axis ticks are generated automatically, the x and y ranges are established automatically,
	   the canvas "box" is adjusted to fit the space the axes occupy so as not to clip.
	   Additionally, it shows how you can use the "Descending" property of continuous ranges to change the ordering of
	   how values (including ticks) are drawn.
	*/

	graph := chart.Chart{
		Height: 500,
		Width:  500,
		XAxis: chart.XAxis{
			Style: chart.StyleShow(),
			/*Range: &chart.ContinuousRange{
				Descending: true,
			},*/
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Descending: true,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
