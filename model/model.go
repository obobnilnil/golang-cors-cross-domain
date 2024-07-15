package model

// type LoginRequest struct {
// 	Validate bool `json:"validate"`
// }

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MalopResolutionTracking struct {
	Timestamp         int64 `json:"timestamp"`
	ClosedMalopsCount int   `json:"closed_malops_count"`
	TotalMalopsCount  int   `json:"total_malops_count"`
}

type MalopResolutionRequest struct {
	MalopResolution []MalopResolutionTracking `json:"malop_resolution_tracking"`
}
