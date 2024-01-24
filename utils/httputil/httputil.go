package httputil

import (
	"encoding/json"
	"io"
	"net/http"
)

// RequestBodyToStruct destination using &struct
func RequestBodyToStruct(w http.ResponseWriter, body io.ReadCloser, destination interface{}) error {
	// Read body
	b, err := io.ReadAll(body)
	defer body.Close()
	if err != nil {
		return err
	}

	// Unmarshal
	err = json.Unmarshal(b, destination)
	if err != nil {
		return err
	}
	return nil
}

// func SendRequest(appContext model.AppContext, method string, url string, formData url.Values, listHeader map[string]string) (result string, restyResp *resty.Response, err error) {
// 	// Inititate Client
// 	restyClient := NewRestyClientWithJaeger(appContext)

// 	// set attempt
// 	attempt := 0

// 	// Set Retry
// 	restyClient.SetRetryCount(constant.RETRY_COUNT)
// 	restyClient.SetRetryWaitTime(time.Duration(constant.RETRY_WAIT_TIME) * time.Second)
// 	restyClient.SetRetryMaxWaitTime(time.Duration(constant.RETRY_MAX_WAIT_TIME) * time.Second)
// 	restyClient = restyClient.AddRetryCondition(utils.RestyRetryCondition(&attempt, url))

// 	// Initate Request
// 	restyReq := restyClient.R()

// 	restyReq = restyReq.
// 		SetFormDataFromValues(formData).
// 		SetHeaders(listHeader)
// 	resp := &resty.Response{}

// 	switch method {
// 	case http.MethodGet:
// 		resp, err = restyReq.Get(url)
// 	case http.MethodPost:
// 		resp, err = restyReq.Post(url)
// 	}

// 	if err != nil {
// 		log.Error().Msgf("error resty param %v,%v erorr %v", url, formData, err)
// 		return result, resp, err
// 	}

// 	result = string(resp.Body())
// 	return result, resp, err
// }

// // Init Resty
// func NewRestyClientWithJaeger(appContext model.AppContext) *resty.Client {
// 	var spanResty opentracing.Span
// 	var ctx context.Context

// 	// Inititate Client
// 	restyClient := resty.New()

// 	// Set Before Request
// 	restyClient = restyClient.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
// 		spanResty, ctx = jaegerutil.StartSpan(appContext.Context, utility.GetCurrentFunctionName())

// 		appContext = model.AppContext{
// 			Context:          ctx,
// 			MandatoryRequest: appContext.MandatoryRequest,
// 		}

// 		// Mapping Request
// 		reqByte, err := json.Marshal(structs.Map(r))
// 		if err != nil {
// 			return nil
// 		}
// 		spanResty.LogKV("request", string(reqByte))

// 		return nil
// 	})

// 	// Set After Response
// 	restyClient = restyClient.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
// 		defer spanResty.Finish()
// 		return nil
// 	})

// 	return restyClient
// }
