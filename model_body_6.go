/*
 * Runalyze Personal API
 *
 * Go client for the [Runalyze Personal API](https://runalyze.com/help/article/personal-api).
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package runalyze

type Body6 struct {
	// Only date will be stored
	DateTime string `json:"date_time"`
	// Heart Rate in bpm
	HeartRate int32 `json:"heart_rate"`
}
