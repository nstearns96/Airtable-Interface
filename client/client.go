package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	at "github.com/nstearns96/Airtable-Interface/airtable"

	"golang.org/x/time/rate"
)

type Client struct {
	apiClient  *http.Client
	apiLimiter *rate.Limiter
	token      string
}

const (
	timeBetweenReqs = time.Second / at.ReqsPerSecond
)

func NewClient(tok string) *Client {
	return &Client{
		apiClient:  http.DefaultClient,
		apiLimiter: rate.NewLimiter(rate.Every(timeBetweenReqs), 1),
		token:      tok,
	}
}

func (cl *Client) ListBases(offset string) (at.ListBasesResponse, error) {
	vals := url.Values{}
	if offset != "" {
		vals.Add("offset", offset)
	}

	ctx := context.Background()
	baseRequest, err := cl.prepareRequest(ctx, "GET", at.BasesURL, vals, nil)
	if err != nil {
		return at.ListBasesResponse{}, err
	}

	resp, err := cl.doRequest(ctx, baseRequest)
	if err != nil {
		return at.ListBasesResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.ListBasesResponse](baseRequest.URL.RawPath, resp)
}

func (cl *Client) GetBaseSchema(baseID string, visibleFieldIDs bool) (at.GetBaseSchemaResponse, error) {
	vals := url.Values{}
	if visibleFieldIDs {
		vals.Add("include", "visibleFieldIds")
	}

	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.TablesURL, baseID)
	baseSchemaReq, err := cl.prepareRequest(ctx, "GET", reqUrl, vals, nil)
	if err != nil {
		return at.GetBaseSchemaResponse{}, err
	}

	resp, err := cl.doRequest(ctx, baseSchemaReq)
	if err != nil {
		return at.GetBaseSchemaResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.GetBaseSchemaResponse](baseSchemaReq.URL.RawPath, resp)
}

func (cl *Client) CreateBase(name string, workspace string, tables []at.TableConfig) (at.CreateBaseResponse, error) {
	jsonData := struct {
		Name      string           `json:"name"`
		Workspace string           `json:"workspaceId"`
		Tables    []at.TableConfig `json:"tables"`
	}{
		name,
		workspace,
		tables,
	}

	bodyBytes, err := json.Marshal(jsonData)
	if err != nil {
		return at.CreateBaseResponse{}, err
	}

	body := bytes.NewReader(bodyBytes)
	ctx := context.Background()
	createBaseReq, err := cl.prepareRequest(ctx, "POST", at.BasesURL, nil, body)
	if err != nil {
		return at.CreateBaseResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.CreateBaseResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.CreateBaseResponse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) UpdateTable(baseID string, tableIdOrName string, name string, description string) (at.UpdateTableRepsonse, error) {
	jsonData := struct {
		Name string `json:"name,omitempty"`
		Desc string `json:"description,omitempty"`
	}{
		name,
		description,
	}

	bodyBytes, err := json.Marshal(jsonData)
	if err != nil {
		return at.UpdateTableRepsonse{}, err
	}

	body := bytes.NewReader(bodyBytes)
	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.TableURL, baseID, tableIdOrName)
	createBaseReq, err := cl.prepareRequest(ctx, "PATCH", reqUrl, nil, body)
	if err != nil {
		return at.UpdateTableRepsonse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.UpdateTableRepsonse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.UpdateTableRepsonse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) CreateTable(baseID string, name string, description string, fields []at.FieldConfig) (at.CreateTableResponse, error) {
	jsonData := struct {
		Name   string           `json:"name"`
		Desc   string           `json:"description,omitempty"`
		Fields []at.FieldConfig `json:"fields"`
	}{
		name,
		description,
		fields,
	}

	bodyBytes, err := json.Marshal(jsonData)
	if err != nil {
		return at.CreateTableResponse{}, err
	}

	body := bytes.NewReader(bodyBytes)
	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.TablesURL, baseID)
	createBaseReq, err := cl.prepareRequest(ctx, "POST", reqUrl, nil, body)
	if err != nil {
		return at.CreateTableResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.CreateTableResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.CreateTableResponse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) UpdateField(baseID string, tableID string, fieldID string, name string, description string) (at.UpdateFieldResponse, error) {
	jsonData := struct {
		Name string `json:"name,omitempty"`
		Desc string `json:"description,omitempty"`
	}{
		name,
		description,
	}

	bodyBytes, err := json.Marshal(jsonData)
	if err != nil {
		return at.UpdateFieldResponse{}, err
	}

	body := bytes.NewReader(bodyBytes)
	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.FieldURL, baseID, tableID, fieldID)
	createBaseReq, err := cl.prepareRequest(ctx, "PATCH", reqUrl, nil, body)
	if err != nil {
		return at.UpdateFieldResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.UpdateFieldResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.UpdateFieldResponse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) CreateField(baseID string, tableID string, field at.FieldConfig) (at.CreateFieldResponse, error) {
	bodyBytes, err := json.Marshal(field)
	if err != nil {
		return at.UpdateFieldResponse{}, err
	}

	body := bytes.NewReader(bodyBytes)
	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.FieldsURL, baseID, tableID)
	createBaseReq, err := cl.prepareRequest(ctx, "POST", reqUrl, nil, body)
	if err != nil {
		return at.UpdateFieldResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.UpdateFieldResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.UpdateFieldResponse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) ListRecords(baseID string, tableIdOrName string, queryParams at.ListRecordsQueryParams) (at.ListRecordsResponse, error) {
	vals := url.Values{}
	if queryParams.TimeZone != "" {
		vals.Add("timeZone", queryParams.TimeZone)
	}
	if queryParams.UserLocale != "" {
		vals.Add("userLocale", queryParams.UserLocale)
	}
	if queryParams.PageSize > 0 {
		vals.Add("pageSize", fmt.Sprintf("%d", queryParams.PageSize))
	}
	if queryParams.MaxRecords > 0 {
		vals.Add("maxRecords", fmt.Sprintf("%d", queryParams.MaxRecords))
	}
	if queryParams.Offset != "" {
		vals.Add("offset", queryParams.Offset)
	}
	if queryParams.View != "" {
		vals.Add("view", queryParams.View)
	}
	if queryParams.Sort.Field != "" {
		sortJson, err := json.Marshal(queryParams.Sort)
		if err != nil {
			return at.ListRecordsResponse{}, err
		}

		vals.Add("view", string(sortJson))
	}
	if queryParams.FiltByFormula != "" {
		vals.Add("filterByFormula", queryParams.FiltByFormula)
	}
	if queryParams.CellFormat != "" {
		vals.Add("cellFormat", queryParams.CellFormat)
	}
	for _, field := range queryParams.Fields {
		vals.Add("fields", field)
	}
	if queryParams.ReturnFieldsByID {
		vals.Add("returnFieldsByFieldsId", fmt.Sprintf("%t", queryParams.ReturnFieldsByID))
	}
	if queryParams.CommentCount {
		vals.Add("recordMetadata", "commentCount")
	}

	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.RecordsURL, baseID, tableIdOrName)
	createBaseReq, err := cl.prepareRequest(ctx, "GET", reqUrl, vals, nil)
	if err != nil {
		return at.ListRecordsResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.ListRecordsResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.ListRecordsResponse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) GetRecord(baseID string, tableIdOrName string, recordID string, cellFormat string, returnFieldsByID bool) (at.GetRecordResponse, error) {
	vals := url.Values{}
	if cellFormat != "" {
		vals.Add("cellFormat", cellFormat)
	}
	if returnFieldsByID {
		vals.Add("returnFieldsByFieldId", "true")
	}

	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.RecordURL, baseID, tableIdOrName, recordID)
	createBaseReq, err := cl.prepareRequest(ctx, "GET", reqUrl, vals, nil)
	if err != nil {
		return at.GetRecordResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.GetRecordResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.GetRecordResponse](resp.Request.URL.RawPath, resp)
}

// WARNING: Airtable will throttle upsert requests differently than the standard rate limit throttling policy
func (cl *Client) UpdateMultipleRecords(baseID string, tableIdOrName string, performUpsert bool, fieldsToMergeOn []string, returnFieldsById bool, typecast bool, records []at.Record, destructive bool) (at.UpdateMultiRecordsResponse, error) {
	type UpsertData struct {
		MergeFields []string `json:"fieldsToMergeOn,omitempty"`
	}

	jsonData := struct {
		Upsert           *UpsertData `json:"performUpsert,omitempty"`
		ReturnFieldsByID bool        `json:"returnFieldsByFieldId,omitempty"`
		Typecast         bool        `json:"typecast,omitempty"`
		Records          []at.Record `json:"records"`
	}{
		ReturnFieldsByID: returnFieldsById,
		Typecast:         typecast,
		Records:          records,
		Upsert:           nil,
	}

	if performUpsert {
		jsonData.Upsert = &UpsertData{fieldsToMergeOn}
	}

	bodyBytes, err := json.Marshal(jsonData)
	if err != nil {
		return at.UpdateMultiRecordsResponse{}, err
	}

	body := bytes.NewReader(bodyBytes)
	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.RecordsURL, baseID, tableIdOrName)

	reqMethod := "PATCH"
	if destructive {
		reqMethod = "PUT"
	}

	createBaseReq, err := cl.prepareRequest(ctx, reqMethod, reqUrl, nil, body)
	if err != nil {
		return at.UpdateMultiRecordsResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.UpdateMultiRecordsResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.UpdateMultiRecordsResponse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) UpdateRecord(baseID string, tableIdOrName string, recordID string, fields map[string]any, returnFieldsById bool, typecast bool, destructive bool) (at.UpdateRecordResponse, error) {
	jsonData := struct {
		ReturnFieldsByID bool           `json:"returnFieldsByFieldId,omitempty"`
		Typecast         bool           `json:"typecase,omitempty"`
		Fields           map[string]any `json:"fields"`
	}{
		ReturnFieldsByID: returnFieldsById,
		Typecast:         typecast,
		Fields:           fields,
	}
	bodyBytes, err := json.Marshal(jsonData)
	if err != nil {
		return at.UpdateRecordResponse{}, err
	}

	body := bytes.NewReader(bodyBytes)
	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.RecordURL, baseID, tableIdOrName, recordID)

	reqMethod := "PATCH"
	if destructive {
		reqMethod = "PUT"
	}

	createBaseReq, err := cl.prepareRequest(ctx, reqMethod, reqUrl, nil, body)
	if err != nil {
		return at.UpdateRecordResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.UpdateRecordResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.UpdateRecordResponse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) CreateRecords(baseID string, tableIdOrName string, multiple bool, fields map[string]any, records []at.Record, typecast bool, returnFieldsById bool) (at.CreateRecordsResponse, error) {
	if multiple && len(records) > at.MaxRecordsCreatedOrDestroyed {
		return at.CreateRecordsResponse{}, fmt.Errorf("cannot create more than %d records at a time", at.MaxRecordsCreatedOrDestroyed)
	}

	jsonData := struct {
		ReturnFieldsByID bool           `json:"returnFieldsByFieldId,omitempty"`
		Typecast         bool           `json:"typecase,omitempty"`
		Fields           map[string]any `json:"fields,omitempty"`
		Records          []at.Record    `json:"records,omitempty"`
	}{
		ReturnFieldsByID: returnFieldsById,
		Typecast:         typecast,
		Fields:           fields,
		Records:          records,
	}
	bodyBytes, err := json.Marshal(jsonData)
	if err != nil {
		return at.CreateRecordsResponse{}, err
	}

	body := bytes.NewReader(bodyBytes)
	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.RecordsURL, baseID, tableIdOrName)
	createBaseReq, err := cl.prepareRequest(ctx, "POST", reqUrl, nil, body)
	if err != nil {
		return at.CreateRecordsResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.CreateRecordsResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.CreateRecordsResponse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) DeleteMultiRecords(baseID string, tableIdOrName string, recordIDs []string) (at.DeleteMultiRecordsResponse, error) {
	if len(recordIDs) > at.MaxRecordsCreatedOrDestroyed {
		return at.DeleteMultiRecordsResponse{}, fmt.Errorf("cannot delete more than %d records at a time", at.MaxRecordsCreatedOrDestroyed)
	}

	vals := url.Values{}
	for _, recordID := range recordIDs {
		vals.Add("records", recordID)
	}

	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.RecordsURL, baseID, tableIdOrName)
	createBaseReq, err := cl.prepareRequest(ctx, "DELETE", reqUrl, vals, nil)
	if err != nil {
		return at.DeleteMultiRecordsResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.DeleteMultiRecordsResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.DeleteMultiRecordsResponse](resp.Request.URL.RawPath, resp)
}

func (cl *Client) DeleteRecord(baseID string, tableIdOrName string, recordID string) (at.DeleteRecordResponse, error) {
	ctx := context.Background()
	reqUrl := fmt.Sprintf(at.RecordURL, baseID, tableIdOrName, recordID)
	createBaseReq, err := cl.prepareRequest(ctx, "DELETE", reqUrl, nil, nil)
	if err != nil {
		return at.DeleteRecordResponse{}, err
	}

	resp, err := cl.doRequest(ctx, createBaseReq)
	if err != nil {
		return at.DeleteRecordResponse{}, err
	}

	defer resp.Body.Close()

	return parseResponse[at.DeleteRecordResponse](resp.Request.URL.RawPath, resp)
}

func parseResponse[T any](reqUrl string, resp *http.Response) (T, error) {
	var baseResp T

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return baseResp, at.MakeHTTPClientError(reqUrl, resp)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return baseResp, err
	}

	err = json.Unmarshal(b, &baseResp)
	if err != nil {
		return baseResp, err
	}

	return baseResp, nil
}

func (cl *Client) prepareRequest(ctx context.Context, method string, reqUrl string, params url.Values, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, reqUrl, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cl.token))

	if params != nil {
		request.URL.RawQuery = params.Encode()
	}

	return request, nil
}

func (cl *Client) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	cl.apiLimiter.Wait(ctx)

	return cl.apiClient.Do(req)
}
