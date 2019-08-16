/*
 * Detectors API
 *
 * **Detectors** define rules for identifying conditions of interest to the customer, and the notifications to send when the conditions occur or stop occurring.
 *
 * API version: 2.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package detector

type CreateUpdateDetectorRequest struct {
	AuthorizedWriters AuthorizedWriters `json:"authorizedWriters,omitempty"`
	// User-defined JSON object containing metadata
	CustomProperties string `json:"customProperties,omitempty"`
	// Description of the detector. It appears in the Detector window displayed from the web UI Actions menu
	Description string `json:"description,omitempty"`
	// The number of milliseconds to wait for late datapoints before rejecting them for inclusion in the detector analysis. The default is to detect and apply a sensible value automatically (this option can also be explicitly chosen by setting the property to 0).
	MaxDelay *int32 `json:"maxDelay,omitempty"`
	// The displayed name of the detector in the dashboard
	Name string `json:"name,omitempty"`
	// The SignalFlow program that populates the detector. The program must include one or more detect functions and each detect function must be modified by a publish stream method with a label that's unique across the program. If you wish to support custom notification messages that include input data you must use variables to assign the detect conditions . If more than one line of SignalFlow is included, each line should be separated by either semicolons (\";\") or new line characters (\"\\n\"). See the [Detectors Overview](https://developers.signalfx.com/v2/reference.html#detectors-overview) for more information.
	ProgramText string `json:"programText,omitempty"`
	// An array of *rules* that define aspects of an alert. These include the alert\\'s severity and the specification of how notifications are sent when the alert is triggered. Each rule is connected to a specific detect function in the programText property by the value of the label in its publish method. The connection is to a set of one or more notification directives and an severity indicator. A single condition can be used in multiple rules if different severity indicators are desired for different notification methods. <p> To see the properties for a rule, expand the definition of the rules array. What you see is the \"rules\" object that specifies a single rule.
	Rules []*Rule `json:"rules,omitempty"`
	// A an array of keywords that filters detectors by one of their fields. Use tags to indicate the state of a detector or its data source (for example, you can label a detector with a \"prod\" tag to indicate that it monitors a production environment).
	Tags []string `json:"tags,omitempty"`
	// IDs of teams associated with this detector. Teams associated with a detector can see the detector and its active alerts on the team's landing page in the web UI. The list of teams associated with a detector is independent of notification settings.  Teams specified in this field don\\'\\'t automatically get notified of new alerts, and teams that choose to get alerts do not have to display the detector on their team landing page in the web application.
	Teams []string `json:"teams,omitempty"`
	// Options that control the appearance of a detector in the SignalFx web UI.
	VisualizationOptions *Visualization `json:"visualizationOptions,omitempty"`
}
