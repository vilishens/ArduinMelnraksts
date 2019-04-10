package pointConfig

type IntervalOnOff struct {
	Pin      string `json:"pin"`
	State    string `json:"state"`
	Interval string `json:"interval"`
}

type IntervalArr []IntervalOnOff

type Intervals map[string]IntervalArr
