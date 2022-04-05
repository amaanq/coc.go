package client

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/amaanq/coc.go/clan"
	"github.com/amaanq/coc.go/labels"
	"github.com/amaanq/coc.go/location"
)

type Base64String string

type searchOptions struct {
	limit  int
	before Base64String
	after  Base64String
}

func SearchOptions() *searchOptions {
	return &searchOptions{}
}

func (s *searchOptions) SetLimit(limit int) *searchOptions {
	if limit < 0 {
		return s
	}
	s.limit = limit
	return s
}

// Automatically converts to clash base64 json
func (s *searchOptions) SetBefore(before int) *searchOptions {
	if before < 0 || s.after != "" {
		return s
	}
	s.before = Base64String(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"pos": %d}`, before))))
	return s
}

// Automatically converts to clash base64 json
func (s *searchOptions) SetAfter(after int) *searchOptions {
	if after < 0 || s.before != "" {
		return s
	}
	s.after = Base64String(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"pos": %d}`, after))))
	return s
}

func (s *searchOptions) ToQuery() string {
	if s.hasNoOptions() {
		return ""
	}
	strbld := strings.Builder{}
	strbld.WriteString("?")
	if s.limit > 0 {
		strbld.WriteString(fmt.Sprintf("limit=%d&", s.limit))
	}
	if s.before != "" {
		strbld.WriteString(fmt.Sprintf("before=%s&", s.before))
	}
	if s.after != "" {
		strbld.WriteString(fmt.Sprintf("after=%s&", s.after))
	}
	return strbld.String()[:len(strbld.String())-1]
}

func (s *searchOptions) hasNoOptions() bool {
	return s.limit == 0 && s.before == "" && s.after == ""
}

type clanSearchOptions struct {
	name          string
	warFrequency  clan.WarFrequency
	locationID    location.LocationID
	minMembers    int
	maxMembers    int
	minClanPoints int
	minClanLevel  int
	limit         int
	before        Base64String
	after         Base64String
	labelIds      string
}

func ClanSearchOptions() *clanSearchOptions {
	return &clanSearchOptions{}
}

func (c *clanSearchOptions) SetName(name string) *clanSearchOptions {
	if len(name) < 3 {
		return c
	}
	c.name = name
	return c
}

func (c *clanSearchOptions) SetWarFrequency(warFrequency clan.WarFrequency) *clanSearchOptions {
	if !warFrequency.Valid() {
		return c
	}
	c.warFrequency = warFrequency
	return c
}

func (c *clanSearchOptions) SetLocation(locationID location.LocationID) *clanSearchOptions {
	if !locationID.Valid() {
		return c
	}
	c.locationID = locationID
	return c
}

func (c *clanSearchOptions) SetMinMembers(minMembers int) *clanSearchOptions {
	if minMembers < 2 || minMembers > 50 {
		return c
	}
	c.minMembers = minMembers
	return c
}

func (c *clanSearchOptions) SetMaxMembers(maxMembers int) *clanSearchOptions {
	if maxMembers < 1 || maxMembers > 50 {
		return c
	}
	c.maxMembers = maxMembers
	return c
}

func (c *clanSearchOptions) SetMinClanPoints(minClanPoints int) *clanSearchOptions {
	if minClanPoints < 0 {
		return c
	}
	c.minClanPoints = minClanPoints
	return c
}

func (c *clanSearchOptions) SetMinClanLevel(minClanLevel int) *clanSearchOptions {
	if minClanLevel < 2 {
		return c
	}
	c.minClanLevel = minClanLevel
	return c
}

func (c *clanSearchOptions) SetLimit(limit int) *clanSearchOptions {
	if limit < 0 {
		return c
	}
	c.limit = limit
	return c
}

// Automatically converts to clash base64 json
func (c *clanSearchOptions) SetBefore(before int) *clanSearchOptions {
	if before < 0 || c.after != "" {
		return c
	}
	c.before = Base64String(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"pos": %d}`, before))))
	return c
}

// Automatically converts to clash base64 json
func (c *clanSearchOptions) SetAfter(after int) *clanSearchOptions {
	if after < 0 || c.before != "" {
		return c
	}
	c.after = Base64String(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"pos": %d}`, after))))
	return c
}

func (c *clanSearchOptions) SetLabelIds(labelID labels.ClanLabelID) *clanSearchOptions {
	if !labelID.Valid() {
		return c
	}
	if len(strings.Split(c.labelIds, ",")) > 3 {
		return c
	}
	if len(c.labelIds) > 0 {
		c.labelIds += ","
	}
	c.labelIds += fmt.Sprint(labelID)
	return c
}

func (c *clanSearchOptions) ToQuery() string {
	if c.hasNoOptions() {
		return ""
	}

	strbld := strings.Builder{}
	strbld.Write([]byte("?"))
	if c.name != "" {
		strbld.Write([]byte(fmt.Sprintf("name=%s&", c.name)))
	}
	if c.warFrequency != "" {
		strbld.Write([]byte(fmt.Sprintf("warFrequency=%s&", c.warFrequency)))
	}
	if c.locationID != 0 {
		strbld.Write([]byte(fmt.Sprintf("locationId=%d&", c.locationID)))
	}
	if c.minMembers != 0 {
		strbld.Write([]byte(fmt.Sprintf("minMembers=%d&", c.minMembers)))
	}
	if c.maxMembers != 0 {
		strbld.Write([]byte(fmt.Sprintf("maxMembers=%d&", c.maxMembers)))
	}
	if c.minClanPoints != 0 {
		strbld.Write([]byte(fmt.Sprintf("minClanPoints=%d&", c.minClanPoints)))
	}
	if c.minClanLevel != 0 {
		strbld.Write([]byte(fmt.Sprintf("minClanLevel=%d&", c.minClanLevel)))
	}
	if c.limit != 0 {
		strbld.Write([]byte(fmt.Sprintf("limit=%d&", c.limit)))
	}
	if c.before != "" {
		strbld.Write([]byte(fmt.Sprintf("before=%s&", c.before)))
	}
	if c.after != "" {
		strbld.Write([]byte(fmt.Sprintf("after=%s&", c.after)))
	}
	if c.labelIds != "" {
		strbld.Write([]byte(fmt.Sprintf("labelIds=%s&", c.labelIds)))
	}
	return strbld.String()[:len(strbld.String())-1]
}

func (c *clanSearchOptions) hasNoOptions() bool {
	return c.name == "" && c.warFrequency == "" && c.locationID == 0 && c.minMembers == 0 && c.maxMembers == 0 && c.minClanPoints == 0 && c.minClanLevel == 0 && c.limit == 0 && c.before == "" && c.after == "" && c.labelIds == ""
}
