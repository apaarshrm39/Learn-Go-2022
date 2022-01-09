package models

// TemplateData holds data send from handlers to templates
type TemplateData struct {
	// Whatever type of Data I may need to be sending
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // when not sure what type of data use interface
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
