package perf

import (
	"fmt"
	"testing"
	"time"

	"github.com/robatussum/kpis/model"
)

func TestMTBF(t *testing.T) {
	cases := []struct {
		c        model.Component
		expected int64
	}{
		{
			model.Component{
				Failures: []model.Failure{},
			},
			0,
		},
		{
			model.Component{
				Failures: []model.Failure{
					{
						StartTime: time.Date(2021, time.April, 17, 5, 1, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 5, 6, 0, 0, time.UTC),
						Cause:     "",
					},
				},
			},
			0,
		},
		{
			model.Component{
				Failures: []model.Failure{
					{
						StartTime: time.Date(2021, time.April, 17, 6, 1, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 6, 11, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						StartTime: time.Date(2021, time.April, 17, 7, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 1, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						StartTime: time.Date(2021, time.April, 17, 7, 43, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 49, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						StartTime: time.Date(2021, time.April, 17, 9, 10, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 9, 51, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						StartTime: time.Date(2021, time.April, 17, 11, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 12, 31, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
				},
			},
			0,
		},
	}

	for _, c := range cases {
		r := MTBF(c.c)
		t.Logf("I: Calculated MTBF of - %v", r)
		if r != c.expected {
			t.Errorf("Error: Failed to calculate MTBF %v\nExpected %v", r, c)
		}
	}
}

func TestMDT(t *testing.T) {
	cases := []struct {
		c        model.Component
		expected int64
	}{
		{
			model.Component{
				Failures: []model.Failure{},
			},
			0,
		},
		{
			model.Component{
				Failures: []model.Failure{
					{
						StartTime: time.Date(2021, time.April, 17, 5, 1, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 5, 6, 0, 0, time.UTC),
						Cause:     "",
					},
				},
			},
			300,
		},
		{
			model.Component{
				Failures: []model.Failure{
					{
						// 10 mins
						StartTime: time.Date(2021, time.April, 17, 6, 1, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 6, 11, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 1 min
						StartTime: time.Date(2021, time.April, 17, 7, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 1, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 6 mins
						StartTime: time.Date(2021, time.April, 17, 7, 43, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 49, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 41 mins
						StartTime: time.Date(2021, time.April, 17, 9, 10, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 9, 51, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 91 mins
						StartTime: time.Date(2021, time.April, 17, 11, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 12, 31, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
				},
			},
			1788,
		},
	}

	for _, c := range cases {
		r := MDT(c.c)
		t.Logf("I: Calculated MDT of - %v", r)
		if r != c.expected {
			t.Errorf("Error: Failed to calculate MDT %v\nExpected %v", r, c)
		}
	}
}

func TestMDTS(t *testing.T) {
	cases := []struct {
		c1       model.Component
		c2       model.Component
		expected int64
	}{
		{
			model.Component{
				Failures: []model.Failure{},
			},
			model.Component{
				Failures: []model.Failure{},
			},
			0,
		},
		{
			model.Component{
				Failures: []model.Failure{
					{
						StartTime: time.Date(2021, time.April, 17, 5, 1, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 5, 6, 0, 0, time.UTC),
						Cause:     "",
					},
				},
			},
			model.Component{
				Failures: []model.Failure{
					{
						StartTime: time.Date(2021, time.April, 17, 5, 1, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 5, 6, 0, 0, time.UTC),
						Cause:     "",
					},
				},
			},
			300,
		},
		{
			model.Component{
				Failures: []model.Failure{
					{
						// 10 mins
						StartTime: time.Date(2021, time.April, 17, 6, 1, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 6, 11, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 1 min
						StartTime: time.Date(2021, time.April, 17, 7, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 1, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 6 mins
						StartTime: time.Date(2021, time.April, 17, 7, 43, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 49, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 41 mins
						StartTime: time.Date(2021, time.April, 17, 9, 10, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 9, 51, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 91 mins
						StartTime: time.Date(2021, time.April, 17, 11, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 12, 31, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
				},
			},
			model.Component{
				Failures: []model.Failure{
					{
						// 6 mins
						StartTime: time.Date(2021, time.April, 17, 6, 1, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 6, 7, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 30 secs
						StartTime: time.Date(2021, time.April, 17, 7, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 0, 30, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 2 mins
						StartTime: time.Date(2021, time.April, 17, 7, 43, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 45, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 21 mins
						StartTime: time.Date(2021, time.April, 17, 9, 10, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 9, 31, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						// 81 mins
						StartTime: time.Date(2021, time.April, 17, 11, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 12, 21, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
				},
			},
			1788,
		},
	}

	for _, c := range cases {
		r := MDTS(c.c1, c.c2)
		t.Logf("I: Calculated MDT of - %v", r)
		if r != c.expected {
			t.Errorf("Error: Failed to calculate MDT %v\nExpected %v", r, c)
		}
	}
}

func TestCalcGap(t *testing.T) {
	cases := []struct {
		c        model.Component
		expected int64
	}{
		{
			model.Component{
				Failures: []model.Failure{
					{
						StartTime: time.Date(2021, time.April, 17, 6, 1, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 6, 11, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						StartTime: time.Date(2021, time.April, 17, 7, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 1, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						StartTime: time.Date(2021, time.April, 17, 7, 43, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 7, 49, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						StartTime: time.Date(2021, time.April, 17, 9, 10, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 9, 51, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
					{
						StartTime: time.Date(2021, time.April, 17, 11, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2021, time.April, 17, 12, 31, 0, 0, time.UTC),
						Cause:     "Glitch",
					},
				},
			},
			0,
		},
	}

	for _, c := range cases {
		gaps := int64(0)
		for i := 0; i < len(c.c.Failures)-1; i++ {
			gap := c.c.Failures[i].EndTime.Unix() - c.c.Failures[i+1].StartTime.Unix()
			gaps += gap
			fmt.Printf("Gap %v\n", gap)
			t.Logf("Gap %v\n", gap)
		}

		avg := gaps / int64(len(c.c.Failures)-1)
		fmt.Printf("I: average gap %v\n", avg)
		t.Logf("I: Calculated MTBF of - %v\n", avg)
	}
}
