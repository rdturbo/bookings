package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}  // For sending structs
	CSRFToken string // security token for forms
	Flash string // Success message
	Warning string
	Error string
}
