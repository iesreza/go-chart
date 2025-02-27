package main

import (
	"net/http"

	"github.com/iesreza/go-chart"
	"github.com/iesreza/go-chart/drawing"
	"github.com/iesreza/go-chart/seq"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:    50,
				Left:   25,
				Right:  25,
				Bottom: 10,
			},
			FillColor: drawing.ColorFromHex("efefef"),
		},
		XAxis: chart.XAxis{
			Style: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: seq.Range(1.0, 100.0),
				YValues: seq.RandomValuesWithMax(100, 512),
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func drawChartDefault(res http.ResponseWriter, req *http.Request) {
	graph := chart.Chart{
		Background: chart.Style{
			FillColor: drawing.ColorFromHex("efefef"),
		},
		XAxis: chart.XAxis{
			Style: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: seq.Range(1.0, 100.0),
				YValues: seq.RandomValuesWithMax(100, 512),
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.HandleFunc("/default", drawChartDefault)
	http.ListenAndServe(":8080", nil)
}
