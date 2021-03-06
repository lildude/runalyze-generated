/*
 * Runalyze Personal API
 *
 * Go client for the [Runalyze Personal API](https://runalyze.com/help/article/personal-api).
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package runalyze

type Body3 struct {
	// Date/Time
	DateTime string `json:"date_time"`
	// Weight in kg
	Weight float32 `json:"weight"`
	// Fat Percentage
	FatPercentage float32 `json:"fat_percentage,omitempty"`
	// Water Percentage
	WaterPercentage float32 `json:"water_percentage,omitempty"`
	// Muscle Percentage
	MusclePercentage float32 `json:"muscle_percentage,omitempty"`
	// Bone Percentage
	BonePercentage float32 `json:"bone_percentage,omitempty"`
}
