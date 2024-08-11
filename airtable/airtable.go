package airtable

type Base struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Permissions string `json:"permissionLevel"`
}

type FieldConfig struct {
	ID      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	Name    string `json:"name"`
	Desc    string `json:"description,omitempty"`
	Options any    `json:"options,omitempty"`
}

type ViewType string

type ViewConfig struct {
	ID              string   `json:"id"`
	Type            ViewType `json:"type"`
	Name            string   `json:"name"`
	VisibleFieldIDs []string `json:"visibleFieldIds,omitempty"`
}

type TableModel struct {
	ID             string        `json:"id"`
	PrimaryFieldID string        `json:"primaryFieldID"`
	Name           string        `json:"name"`
	Desc           string        `json:"description,omitempty"`
	Fields         []FieldConfig `json:"fields"`
	Views          []ViewConfig  `json:"views"`
}

type TableConfig struct {
	Name   string        `json:"name"`
	Desc   string        `json:"description,omitempty"`
	Fields []FieldConfig `json:"fields"`
}

type Record struct {
	ID           string         `json:"id,omitempty"`
	CreatedTime  string         `json:"createdTime,omitempty"`
	Fields       map[string]any `json:"fields,omitempty"`
	CommentCount int            `json:"commentCount,omitempty"`
}
