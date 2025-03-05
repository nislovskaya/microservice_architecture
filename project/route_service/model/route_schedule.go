package model

type RouteWithSchedule struct {
	Route    Route      `json:"route"`
	Schedule []Schedule `json:"schedule"`
}
