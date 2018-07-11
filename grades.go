package moira

var eventStateWeight = map[string]int{
	"OK":     0,
	"WARN":   1,
	"ERROR":  100,
	"NODATA": 10000,
}

const (
	// EventHighDegradation is a grade that describes High Degradation
	EventHighDegradation = "HIGH DEGRADATION"
	// EventDegradation is a grade that describes Degradation
	EventDegradation = "DEGRADATION"
	// EventProgress is a grade that describes Progress
	EventProgress = "PROGRESS"
)

// GetEventGrades returns grades based on trigger state
func (eventData *NotificationEvent) GetEventGrades() []string {
	grades := make([]string, 0)
	if oldStateWeight, ok := eventStateWeight[eventData.OldState]; ok {
		if newStateWeight, ok := eventStateWeight[eventData.State]; ok {
			if newStateWeight > oldStateWeight {
				if newStateWeight-oldStateWeight >= 100 {
					grades = append(grades, EventHighDegradation)
				}
				grades = append(grades, EventDegradation)
			}
			if newStateWeight < oldStateWeight {
				grades = append(grades, EventProgress)
			}
		}
	}
	return grades
}
