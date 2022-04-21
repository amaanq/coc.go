package coc

import "time"

type GoldPassSeason struct {
	Start string `json:"startTime,omitempty"`
	End   string `json:"endTime,omitempty"`
}

func (g *GoldPassSeason) StartTime() time.Time {
	parsed, _ := time.Parse(TimestampFormat, g.Start)
	return parsed
}

func (g *GoldPassSeason) EndTime() time.Time {
	parsed, _ := time.Parse(TimestampFormat, g.End)
	return parsed
}
