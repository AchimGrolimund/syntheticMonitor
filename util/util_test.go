package util

import (
	"grolimundSolutions.ch/syntheticMonitor/data"
	"testing"
)

func Test_Count(t *testing.T) {
	testData := data.SyntheticSettings{
		Location: "Test",
		SyntheticUrls: []data.SyntheticUrls{
			data.SyntheticUrls{
				URL:    "www.google1.ch",
				Name:   "Google1",
				Expect: "google1",
			},
			data.SyntheticUrls{
				URL:    "www.google2.ch",
				Name:   "Google2",
				Expect: "google2",
			},
		},
	}
	if num := Count(&testData); num != 2 {
		t.Fail()
	}
}

func Test_Count2(t *testing.T) {
	testData := data.SyntheticSettings{
		Location: "Test",
		SyntheticUrls: []data.SyntheticUrls{
			data.SyntheticUrls{
				URL:    "www.google1.ch",
				Name:   "Google1",
				Expect: "google1",
			},
			data.SyntheticUrls{
				URL:    "www.google2.ch",
				Name:   "Google2",
				Expect: "google2",
			},
			data.SyntheticUrls{
				URL:    "www.google2.ch",
				Name:   "Google2",
				Expect: "google2",
			},
		},
	}
	if num := Count(&testData); num != 3 {
		t.Fail()
	}
}
