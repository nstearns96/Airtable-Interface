package airtable

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// Airtable specifies a limit of 5 request per second per base
	ReqsPerSecond = 5

	MaxRecordsCreatedOrDestroyed = 10

	BasesURL   = "https://api.airtable.com/v0/meta/bases"
	TablesURL  = "https://api.airtable.com/v0/meta/bases/%s/tables"
	TableURL   = "https://api.airtable.com/v0/meta/bases/%s/tables/%s"
	FieldsURL  = "https://api.airtable.com/v0/meta/bases/%s/tables/%s/fields"
	FieldURL   = "https://api.airtable.com/v0/meta/bases/%s/tables/%s/fields/%s"
	RecordsURL = "https://api.airtable.com/v0/%s/%s"
	RecordURL  = "https://api.airtable.com/v0/%s/%s/%s"
)

// API responses

type ListBasesResponse struct {
	Offset string `json:"offset,omitempty"`
	Bases  []Base `json:"bases"`
}

type GetBaseSchemaResponse struct {
	Tables []TableModel `json:"tables"`
}

type CreateBaseResponse struct {
	ID     string       `json:"id"`
	Tables []TableModel `json:"tables"`
}

type UpdateTableRepsonse = TableModel

type CreateTableResponse = TableModel

type UpdateFieldResponse = FieldConfig

type CreateFieldResponse = FieldConfig

type ListRecordsResponse struct {
	Offset  string   `json:"offset,omitempty"`
	Records []Record `json:"records"`
}

type SortSpec struct {
	Field     string `json:"field"`
	Direction string `json:"direction,omitempty"`
}

type ListRecordsQueryParams struct {
	TimeZone         string
	UserLocale       string
	PageSize         int
	MaxRecords       int
	Offset           string
	View             string
	Sort             SortSpec
	FiltByFormula    string
	CellFormat       string
	Fields           []string
	ReturnFieldsByID bool
	CommentCount     bool
}

type GetRecordResponse = Record

type UpdateMultiRecordsResponse struct {
	CreatedRecords []string `json:"createdRecords"`
	UpdatedRecords []string `json:"updatedRecords"`
	Records        []Record `json:"records"`
}

type UpdateRecordResponse = Record

type CreateRecordsResponse struct {
	Records     []Record       `json:"records,omitempty"`
	ID          string         `json:"id,omitempty"`
	CreatedTime string         `json:"createdTime,omitempty"`
	Fields      map[string]any `json:"fields,omitempty"`
}

type DeletedRecord struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

type DeleteMultiRecordsResponse struct {
	Records []DeletedRecord `json:"records"`
}

type DeleteRecordResponse = DeletedRecord

// Enums for constriained string fields

type CellType int64

const (
	CellTypeAiText CellType = iota
	CellTypeMultipleAttachments
	CellTypeAutoNumber
	CellTypeBarcode
	CellTypeButton
	CellTypeCheckbox
	CellTypeCollaborator
	CellTypeCount
	CellTypeCreatedBy
	CellTypeCreatedTime
	CellTypeCurrency
	CellTypeDate
	CellTypeDateTime
	CellTypeDuration
	CellTypeEmail
	CellTypeFormula
	CellTypeLastModifiedBy
	CellTypeLastModifiedTime
	CellTypeMultipleRecordLinks
	CellTypeMultilineText
	CellTypeMultipleLookupValues
	CellTypeMultipleCollaborators
	CellTypeMultipleSelects
	CellTypeNumber
	CellTypePercent
	CellTypePhoneNumber
	CellTypeRating
	CellTypeRichText
	CellTypeRollup
	CellTypeSingleLineText
	CellTypeSingleSelect
	CellTypeExternalSyncSource
	CellTypeURL
)

func (ct CellType) String() string {
	switch ct {
	case CellTypeAiText:
		return "aitText"
	case CellTypeMultipleAttachments:
		return "multipleAttachments"
	case CellTypeAutoNumber:
		return "autoNumber"
	case CellTypeBarcode:
		return "barcode"
	case CellTypeButton:
		return "button"
	case CellTypeCheckbox:
		return "checkbox"
	case CellTypeCollaborator:
		return "collaborator"
	case CellTypeCount:
		return "count"
	case CellTypeCreatedBy:
		return "createdBy"
	case CellTypeCreatedTime:
		return "createdTime"
	case CellTypeCurrency:
		return "currency"
	case CellTypeDate:
		return "date"
	case CellTypeDateTime:
		return "dateTime"
	case CellTypeDuration:
		return "duration"
	case CellTypeEmail:
		return "email"
	case CellTypeFormula:
		return "formula"
	case CellTypeLastModifiedBy:
		return "lastModifiedBy"
	case CellTypeLastModifiedTime:
		return "lastModifiedTime"
	case CellTypeMultipleRecordLinks:
		return "multipleRecordLinks"
	case CellTypeMultilineText:
		return "multielineText"
	case CellTypeMultipleLookupValues:
		return "multipleLookupValues"
	case CellTypeMultipleCollaborators:
		return "multipleCollaborators"
	case CellTypeMultipleSelects:
		return "multipleSelects"
	case CellTypeNumber:
		return "number"
	case CellTypePercent:
		return "percent"
	case CellTypePhoneNumber:
		return "phoneNumber"
	case CellTypeRating:
		return "rating"
	case CellTypeRichText:
		return "richText"
	case CellTypeRollup:
		return "rollup"
	case CellTypeSingleLineText:
		return "singleLineText"
	case CellTypeSingleSelect:
		return "singleSelect"
	case CellTypeExternalSyncSource:
		return "externalSyncSource"
	case CellTypeURL:
		return "url"
	}

	panic("Unrecognized cell type")
}

func (ct CellType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.String())
}

type CheckboxColor int64

const (
	CheckboxColorGreen CheckboxColor = iota
	CheckboxColorTeal
	CheckboxColorCyan
	CheckboxColorBlue
	CheckboxColorPurple
	CheckboxColorPink
	CheckboxColorRed
	CheckboxColorOrange
	CheckboxColorYellow
	CheckboxColorGray
)

func (col CheckboxColor) String() string {
	switch col {
	case CheckboxColorGreen:
		return "greenBright"
	case CheckboxColorTeal:
		return "tealBright"
	case CheckboxColorCyan:
		return "cyanBright"
	case CheckboxColorBlue:
		return "blueBright"
	case CheckboxColorPurple:
		return "purpleBright"
	case CheckboxColorPink:
		return "pinkBright"
	case CheckboxColorRed:
		return "redBright"
	case CheckboxColorOrange:
		return "orangeBright"
	case CheckboxColorYellow:
		return "yellowBright"
	case CheckboxColorGray:
		return "grayBright"
	}

	panic("Unrecognized checkbox color")
}

func (col CheckboxColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(col.String())
}

type CheckboxIcon int64

const (
	CheckboxIconCheck CheckboxIcon = iota
	CheckBoxIconXCheckbox
	CheckboxIconStar
	CheckboxIconThumbsUp
	CheckboxIconFlag
	CheckboxIconDot
)

func (ico CheckboxIcon) String() string {
	switch ico {
	case CheckboxIconCheck:
		return "check"
	case CheckBoxIconXCheckbox:
		return "xCheckbox"
	case CheckboxIconStar:
		return "star"
	case CheckboxIconThumbsUp:
		return "thumbsUp"
	case CheckboxIconFlag:
		return "flag"
	case CheckboxIconDot:
		return "dot"
	}

	panic("Unrecognized checkbox icon")
}

func (ico CheckboxIcon) MarshalJSON() ([]byte, error) {
	return json.Marshal(ico.String())
}

type AiTextCellState int64

const (
	EmptyAiTextCellState AiTextCellState = iota
	LoadingAiTextCellState
	GeneratedAiTextCellState
	ErrorAiTextCellState
)

func (state AiTextCellState) String() string {
	switch state {
	case EmptyAiTextCellState:
		return "empty"
	case LoadingAiTextCellState:
		return "loading"
	case GeneratedAiTextCellState:
		return "generated"
	case ErrorAiTextCellState:
		return "error"
	}

	panic("Unrecognized AI text cell state")
}

func (state AiTextCellState) MarshalJSON() ([]byte, error) {
	return json.Marshal(state.String())
}

type PermissionLevel int64

const (
	PermissionNone PermissionLevel = iota
	PermissionRead
	PermissionComment
	PermissionEdit
	PermissionCreate
)

func (perms PermissionLevel) String() string {
	switch perms {
	case PermissionNone:
		return "none"
	case PermissionRead:
		return "read"
	case PermissionComment:
		return "comment"
	case PermissionEdit:
		return "edit"
	case PermissionCreate:
		return "create"
	}

	panic("Unrecognized permission level")
}

func (perms PermissionLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(perms.String())
}

type DateFormatType int64

const (
	DateFormatLocal DateFormatType = iota
	DateFormatFriendly
	DateFormatUS
	DateFormatEuropean
	DateFormatISO
)

func (dft DateFormatType) GetFormat() string {
	switch dft {
	case DateFormatLocal:
		return "l"
	case DateFormatFriendly:
		return "LL"
	case DateFormatUS:
		return "M/D/YYYY"
	case DateFormatEuropean:
		return "D/M/YYYY"
	case DateFormatISO:
		return "YYYY-MM-DD"
	}

	panic("Unrecognized date format type")
}

func (dft DateFormatType) GetName() string {
	switch dft {
	case DateFormatLocal:
		return "local"
	case DateFormatFriendly:
		return "friendly"
	case DateFormatUS:
		return "us"
	case DateFormatEuropean:
		return "european"
	case DateFormatISO:
		return "iso"
	}

	panic("Unrecognized date format type")
}

func (dft DateFormatType) MarshalJSON() ([]byte, error) {
	dftStruct := struct {
		FormatType string `json:"format"`
		Name       string `json:"name"`
	}{
		FormatType: dft.GetFormat(),
		Name:       dft.GetName(),
	}

	return json.Marshal(dftStruct)
}

type TimeFormatType int64

const (
	TimeFormat12Hour TimeFormatType = iota
	TimeFormat24Hour
)

func (dft TimeFormatType) GetFormat() string {
	switch dft {
	case TimeFormat12Hour:
		return "h:mma"
	case TimeFormat24Hour:
		return "HH:mma"
	}

	panic("Unrecognized Time format type")
}

func (dft TimeFormatType) GetName() string {
	switch dft {
	case TimeFormat12Hour:
		return "12hour"
	case TimeFormat24Hour:
		return "24hour"
	}

	panic("Unrecognized Time format type")
}

func (dft TimeFormatType) MarshalJSON() ([]byte, error) {
	dftStruct := struct {
		FormatType string `json:"format"`
		Name       string `json:"name"`
	}{
		FormatType: dft.GetFormat(),
		Name:       dft.GetName(),
	}

	return json.Marshal(dftStruct)
}

type DurationFormatType int64

const (
	DurationFormatMinutes DurationFormatType = iota
	DurationFormatSeconds
	DurationFormatTenths
	DurationFormatHundredths
	DurationFormatThousandths
)

func (dft DurationFormatType) String() string {
	switch dft {
	case DurationFormatMinutes:
		return "h:mm"
	case DurationFormatSeconds:
		return "h:mm:ss"
	case DurationFormatTenths:
		return "h:mm:ss.S"
	case DurationFormatHundredths:
		return "h:mm:ss.SS"
	case DurationFormatThousandths:
		return "h:mm:ss.SSS"
	}

	panic("Unrecognized Time format type")
}

// Field types and cell values

type CellInfo interface {
	GetType() string
	GetOptions() any
}

type CellFormatAiText struct {
	State     string `json:"state"`
	ErrorType string `json:"errorType,omitempty"`
	IsStale   bool   `json:"isStale"`
	Value     string `json:"string,omitempty"`
}

type CellInfoAiText struct {
	Options CellOptionsAiText `json:"options"`
}

type CellOptionsAiText struct {
	Prompt             []any    `json:"prompt,omitempty"`
	ReferencedFieldIDs []string `json:"referencedFieldIds,omitempty"`
}

type AiPromptInterpolatedField struct {
	Field struct {
		FieldID string `json:"fieldId"`
	} `json:"field"`
}

func (ci *CellInfoAiText) GetType() string {
	return CellTypeAiText.String()
}

func (ct *CellInfoAiText) GetOptions() any {
	return ct.Options
}

type ThumbnailData struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type ThumbnailList struct {
	Full  ThumbnailData `json:"full,omitempty"`
	Large ThumbnailData `json:"large,omitempty"`
	Small ThumbnailData `json:"small,omitempty"`
}

type CellFormatAttachmentRead struct {
	ID         string        `json:"id"`
	Type       string        `json:"type"`
	Filename   string        `json:"filename"`
	Height     int           `json:"height"`
	Size       int           `json:"size"`
	URL        string        `json:"url"`
	Width      int           `json:"width"`
	Thumbnails ThumbnailList `json:"thumbnails"`
}

type CellFormatAttachmentWrite struct {
	Attachments []any
}

type AttachmentURL struct {
	URL      string `json:"url"`
	Filename string `json:"filename,omitempty"`
}

type AttachmentID struct {
	ID string `json:"id"`
}

type CellInfoAttachment struct {
	Options CellOptionsAttachment `json:"options"`
}

type CellOptionsAttachment struct {
	IsReversed bool `json:"isReversed"`
}

func (ci *CellInfoAttachment) GetType() string {
	return CellTypeMultipleAttachments.String()
}

func (ct *CellInfoAttachment) GetOptions() any {
	return ct.Options
}

type CellFormatAutoNumber = int

type CellInfoAutoNumber struct {
}

func (ci *CellInfoAutoNumber) GetType() string {
	return CellTypeAutoNumber.String()
}

func (ct *CellInfoAutoNumber) GetOptions() any {
	return nil
}

type CellFormatBarcode struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text"`
}

type CellInfoBarcode struct {
}

func (ci *CellInfoBarcode) GetType() string {
	return CellTypeBarcode.String()
}

func (ct *CellInfoBarcode) GetOptions() any {
	return nil
}

type CellFormatButton struct {
	Label string `json:"label"`
	URL   string `json:"url,omitempty"`
}

type CellInfoButton struct {
}

func (ci *CellInfoButton) GetType() string {
	return CellTypeButton.String()
}

func (ct *CellInfoButton) GetOptions() any {
	return nil
}

type CellFormatCheckbox = bool

type CellInfoCheckbox struct {
	Options CellOptionsCheckbox `json:"options"`
}

type CellOptionsCheckbox struct {
	Color CheckboxColor `json:"color"`
	Icon  CheckboxIcon  `json:"icon"`
}

func (ci *CellInfoCheckbox) GetType() string {
	return CellTypeCheckbox.String()
}

func (ct *CellInfoCheckbox) GetOptions() any {
	return ct.Options
}

type CellFormatCollaboratorRead struct {
	ID            string          `json:"id"`
	Email         string          `json:"email,omitempty"`
	Name          string          `json:"name,omitempty"`
	PermissionLvl PermissionLevel `json:"permissionLevel,omitempty"`
	ProfilePicURL string          `json:"profilePicUrl,omitempty"`
}

type CellFormatCollaboratorWrite struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type CellInfoCollaborator struct {
}

func (ci *CellInfoCollaborator) GetType() string {
	return CellTypeCollaborator.String()
}

func (ct *CellInfoCollaborator) GetOptions() any {
	return nil
}

type CellFormatCount int64

type CellInfoCount struct {
	Options CellOptionsCount `json:"options"`
}

type CellOptionsCount struct {
	IsValid           bool   `json:"isValid"`
	RecordLinkFieldID string `json:"recordLinkFieldId,omitempty"`
}

func (ci *CellInfoCount) GetType() string {
	return CellTypeCount.String()
}

func (ct *CellInfoCount) GetOptions() any {
	return ct.Options
}

type CellFromatCreatedBy struct {
	ID            string          `json:"id"`
	Email         string          `json:"email,omitempty"`
	Name          string          `json:"name,omitempty"`
	PermissionLvl PermissionLevel `json:"permissionLevel,omitempty"`
	ProfilePicURL string          `json:"profilePicUrl,omitempty"`
}

type CellInfoCreatedBy struct {
}

func (ci *CellInfoCreatedBy) GetType() string {
	return CellTypeCreatedBy.String()
}

func (ct *CellInfoCreatedBy) GetOptions() any {
	return nil
}

type CellFormatCreatedTime string

type CellInfoCreatedTime struct {
	Options CellOptionsCreatedTime `json:"options"`
}

type CellOptionsCreatedTime struct {
	Result any `json:"result,omitempty"`
}

func (ci *CellInfoCreatedTime) GetType() string {
	return CellTypeCreatedTime.String()
}

func (ct *CellInfoCreatedTime) GetOptions() any {
	return ct.Options
}

type CellFormatCurrency = float64

type CellInfoCurrency struct {
	Options CellOptionsCurrency `json:"options"`
}

type CellOptionsCurrency struct {
	Precision int64  `json:"precision"`
	Symbol    string `json:"symbol"`
}

func (ci *CellInfoCurrency) GetType() string {
	return CellTypeCurrency.String()
}

func (ct *CellInfoCurrency) GetOptions() any {
	return ct.Options
}

type CellFormatDate string

type CellInfoDate struct {
	Options CellOptionsDate `json:"options"`
}

type CellOptionsDate struct {
	Format DateFormatType `json:"dateFormat"`
}

func (ci *CellInfoDate) GetType() string {
	return CellTypeDate.String()
}

func (ct *CellInfoDate) GetOptions() any {
	return ct.Options
}

type CellFormatDateTime string

type CellInfoDateTime struct {
	Options CellOptionsDateTime `json:"options"`
}

type CellOptionsDateTime struct {
	TimeZone   TimeZone       `json:"timeZone"`
	DateFormat DateFormatType `json:"dateFormat"`
	TimeFormat TimeFormatType `json:"timeFormat"`
}

func (ci *CellInfoDateTime) GetType() string {
	return CellTypeDateTime.String()
}

func (ct *CellInfoDateTime) GetOptions() any {
	return ct.Options
}

type CellFormatDuration float64

type CellInfoDuration struct {
	Options CellOptionsDuration `json:"options"`
}

type CellOptionsDuration struct {
	Format DurationFormatType `json:"durationFormat"`
}

func (ci *CellInfoDuration) GetType() string {
	return CellTypeDuration.String()
}

func (ct *CellInfoDuration) GetOptions() any {
	return ct.Options
}

type CellFormatEmail string

type CellInfoEmail struct {
}

func (ci *CellInfoEmail) GetType() string {
	return CellTypeEmail.String()
}

func (ct *CellInfoEmail) GetOptions() any {
	return nil
}

type CellFormatFormula any

type CellInfoFormula struct {
	Options CellOptionsFormula `json:"options"`
}

type CellOptionsFormula struct {
	Formula            string   `json:"formula"`
	IsValid            bool     `json:"isValid"`
	ReferencedFieldIDs []string `json:"referenceFieldIds"`
	Result             any      `json:"result"`
}

func (ci *CellInfoFormula) GetType() string {
	return CellTypeFormula.String()
}

func (ct *CellInfoFormula) GetOptions() any {
	return ct.Options
}

type CellFormatLastModifiedBy struct {
	ID            string          `json:"id"`
	Email         string          `json:"email,omitempty"`
	Name          string          `json:"name,omitempty"`
	PermissionLvl PermissionLevel `json:"permissionLevel,omitempty"`
	ProfilePicURL string          `json:"profilePicUrl,omitempty"`
}

type CellInfoLastModifiedBy struct {
}

func (ci *CellInfoLastModifiedBy) GetType() string {
	return CellTypeLastModifiedBy.String()
}

func (ct *CellInfoLastModifiedBy) GetOptions() any {
	return nil
}

type CellFormatLastModifiedTime string

type CellInfoLastModifiedTime struct {
	Options CellOptionsLastModifiedTime `json:"options"`
}

type CellOptionsLastModifiedTime struct {
	IsValid            bool     `json:"isValid"`
	ReferencedFieldIDs []string `json:"referencedFieldIds,omitempty"`
	Result             any      `json:"result"`
}

func (ci *CellInfoLastModifiedTime) GetType() string {
	return CellTypeLastModifiedTime.String()
}

func (ct *CellInfoLastModifiedTime) GetOptions() any {
	return ct.Options
}

type CellFormatMultipleRecordLinks []string

type CellFormatMultipleRecordLinksWebhook struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CellInfoMultipleRecordLinksRead struct {
	Options CellOptionsMultipleRecordLinksRead `json:"options"`
}

type CellInfoMultipleRecordLinksWrite struct {
	Options CellOptionsMultipleRecordLinksWrite `json:"options"`
}

type CellOptionsMultipleRecordLinksRead struct {
	IsReversed               bool   `json:"isReversed"`
	LinkedTableID            string `json:"linkedTableId"`
	PrefersSingleRecordLink  bool   `json:"prefersSingleRecordLink"`
	InverseLinkFieldID       string `json:"inverseLinkFieldId,omitempty"`
	ViewIdForRecordSelection string `json:"viewIdForRecordSelection"`
}

type CellOptionsMultipleRecordLinksWrite struct {
	LinkedTableID            string `json:"linkedTableId"`
	ViewIdForRecordSelection string `json:"viewIdForRecordSelection"`
}

func (ci *CellInfoMultipleRecordLinksRead) GetType() string {
	return CellTypeLastModifiedBy.String()
}

func (ct *CellInfoMultipleRecordLinksRead) GetOptions() any {
	return ct.Options
}

func (ci *CellInfoMultipleRecordLinksWrite) GetType() string {
	return CellTypeMultipleRecordLinks.String()
}

func (ct *CellInfoMultipleRecordLinksWrite) GetOptions() any {
	return ct.Options
}

type CellFormatMultilineText string

type CellInfoMultilineText struct {
}

func (ci *CellInfoMultilineText) GetType() string {
	return CellTypeMultilineText.String()
}

func (ct *CellInfoMultilineText) GetOptions() any {
	return nil
}

type CellFormatMultipleLookupValues any

type CellFormatMultipleLookupValuesWebhook struct {
	ValuesByLinkedRecordID map[string][]any `json:"valuesByLinkedRecordId"`
	LinkedRecordIDs        []string         `json:"linkedRecordIds"`
}

type CellInfoMultipleLookupValues struct {
	Options CellOptionsMultipleLookupValues `json:"options"`
}

type CellOptionsMultipleLookupValues struct {
	FieldIdInLinkedTable string `json:"fieldIdInLinkedTable,omitempty"`
	IsValid              bool   `json:"isValid"`
	RecordLinkFieldID    string `json:"recordLinkFieldId,omitempty"`
	Result               any    `json:"result,omitempty"`
}

func (ci *CellInfoMultipleLookupValues) GetType() string {
	return CellTypeMultipleLookupValues.String()
}

func (ct *CellInfoMultipleLookupValues) GetOptions() any {
	return ct.Options
}

// HTTP specific stuff

type HTTPClientError struct {
	StatusCode int
	Err        error
}

func (e *HTTPClientError) Error() string {
	return fmt.Sprintf("status %d, err: %v", e.StatusCode, e.Err)
}

func MakeHTTPClientError(url string, resp *http.Response) error {
	var resError error

	respStatusText := "Unknown status text"
	switch resp.StatusCode {
	case 400:
		respStatusText = "The request encoding is invalid; the request can't be parsed as a valid JSON."
	case 401:
		respStatusText = "Accessing a protected resource without authorization or with invalid credentials."
	case 402:
		respStatusText = "The account associated with the API key making requests hits a quota that can be increased by upgrading the Airtable account plan."
	case 403:
		respStatusText = "Accessing a protected resource with API credentials that don't have access to that resource."
	case 404:
		respStatusText = "Route or resource is not found. This error is returned when the request hits an undefined route, or if the resource doesn't exist (e.g. has been deleted)."
	case 413:
		respStatusText = "Too Large The request exceeded the maximum allowed payload size. You shouldn't encounter this under normal use."
	case 422:
		respStatusText = "The request data is invalid. This includes most of the base-specific validations. You will receive a detailed error message and code pointing to the exact issue."
	case 429:
		respStatusText = "The rate limit has been exceeded. Please wait 30 seconds before subsequent requests can succeed."
	case 500:
		respStatusText = "Error The server encountered an unexpected condition."
	case 502:
		respStatusText = "Airtable's servers are restarting or an unexpected outage is in progress. You should generally not receive this error, and requests are safe to retry."
	case 503:
		respStatusText = "The server could not process your request in time. The server could be temporarily unavailable, or it could have timed out processing your request. You should retry the request with backoffs."
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		resError = fmt.Errorf("HTTP request failure on %s:\n%d %s\n%s\n\nCannot parse body with err: %w",
			url, resp.StatusCode, resp.Status, respStatusText, err)
	} else {
		resError = fmt.Errorf("HTTP request failure on %s:\n%d %s\n%s\n\nBody: %v",
			url, resp.StatusCode, resp.Status, respStatusText, string(body))
	}

	return &HTTPClientError{
		StatusCode: resp.StatusCode,
		Err:        resError,
	}
}
