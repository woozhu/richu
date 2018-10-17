package WXTime

import (
	"errors"
)

// WXTimeInterval represents an interval between two WXTimes.
type WXTimeInterval struct {
	Start *WXTime
	End *WXTime
}

// NewWXTimeInterval returns a pointer to a new WXTimeInterval instance
func NewWXTimeInterval(start, end *WXTime) (*WXTimeInterval, error) {
	if start.Gte(end) {
		return nil, errors.New("The end date must be greater than the start date.")
	}

	return &WXTimeInterval{
		Start: start,
		End: end,
	}, nil
}

// DiffInHours return the difference in hours between start and end date
func (ci *WXTimeInterval) DiffInHours() int64 {
	return ci.End.DiffInHours(ci.Start, true)
}
