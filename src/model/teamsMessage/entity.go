package teamsMessage

type Facts struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Sections struct {
	Activitytitle    string  `json:"activityTitle"`
	Activitysubtitle string  `json:"activitySubtitle"`
	Activityimage    string  `json:"activityImage"`
	Facts            []Facts `json:"facts"`
	Markdown         bool    `json:"markdown"`
}

type Inputs struct {
	Type        string `json:"@type"`
	ID          string `json:"id"`
	Ismultiline bool   `json:"isMultiline"`
	Title       string `json:"title"`
}

type Actions struct {
	Type   string `json:"@type"`
	Name   string `json:"name"`
	Target string `json:"target"`
}

type Targets struct {
	Os  string `json:"os"`
	URI string `json:"uri"`
}

type Potentialaction struct {
	Type    string    `json:"@type"`
	Name    string    `json:"name"`
	Inputs  []Inputs  `json:"inputs,omitempty"`
	Actions []Actions `json:"actions,omitempty"`
	Targets []Targets `json:"targets,omitempty"`
}

type TeamsMessage struct {
	Type            string            `json:"@type"`
	Context         string            `json:"@context"`
	Themecolor      string            `json:"themeColor"`
	Summary         string            `json:"summary"`
	Sections        []Sections        `json:"sections"`
	Potentialaction []Potentialaction `json:"potentialAction"`
}
