/*
 * Runalyze Personal API
 *
 * Go client for the [Runalyze Personal API](https://runalyze.com/help/article/personal-api).
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package runalyze
// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	Sleep []ApiV1MetricsSleep `json:"sleep,omitempty"`
	BodyComposition []ApiV1MetricsBodyComposition `json:"bodyComposition,omitempty"`
	BloodPressure []ApiV1MetricsBodyComposition `json:"bloodPressure,omitempty"`
	HeartRateRest []ApiV1MetricsBodyComposition `json:"heartRateRest,omitempty"`
	HeartRateMax []ApiV1MetricsBodyComposition `json:"heartRateMax,omitempty"`
	Mental []ApiV1MetricsBodyComposition `json:"mental,omitempty"`
}
