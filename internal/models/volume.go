package models

// VolumeResponse API response struct.
type VolumeResponse struct {
	SpecificVolumeLiquid float64 `json:"specific_volume_liquid"`
	SpecificVolumeVapor  float64 `json:"specific_volume_vapor"`
}

// NewVolumeResponse returns an instance of VolumeResponse.
func NewVolumeResponse(vLiquid, vVapor float64) *VolumeResponse {
	return &VolumeResponse{
		SpecificVolumeLiquid: vLiquid,
		SpecificVolumeVapor:  vVapor,
	}
}
