// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AutomaticRepliesMailTips undocumented
type AutomaticRepliesMailTips struct {
	// Object is the base model of AutomaticRepliesMailTips
	Object
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// MessageLanguage undocumented
	MessageLanguage *LocaleInfo `json:"messageLanguage,omitempty"`
	// ScheduledStartTime undocumented
	ScheduledStartTime *DateTimeTimeZone `json:"scheduledStartTime,omitempty"`
	// ScheduledEndTime undocumented
	ScheduledEndTime *DateTimeTimeZone `json:"scheduledEndTime,omitempty"`
}