package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	// Success will give message Success
	Success = "Success"
	// ErrorInternalServer will throw if any the Internal Server Error happen
	ErrorInternalServer = errors.New("internal server error")
	// ErrorNotFound will throw if the requested item is not exists
	ErrorNotFound = errors.New("your requested item is not found")
	// ErrorDuplicateData will throw if the current action already exists
	ErrorDuplicateData = errors.New("your item already exist")
	// ErrorBadRequest will throw if the given request-body or params is not valid
	ErrorBadRequest = errors.New("given param is not valid")
	// ErrorUnauthorized will throw if not authorized
	ErrorUnauthorized = errors.New("unauthorized")
	// ErrorInternalServerThirdParty will throw if any Internal Server Error from third party
	ErrorInternalServerThirdParty = errors.New("third party internal server error")
	// ErrorResultNotFound will throw if endpoint returns an empty list
	ErrorResultNotFound = errors.New("Result Not Found")

	ErrorInvalidBearerToken  = errors.New("unauthorized: Invalid Bearer token")
	ErrorBearer              = errors.New("unauthorized: Bearer token is missing or empty")
	ErrorState               = errors.New("unauthorized: State is invalid")
	ErrorRefreshTokenRevoked = errors.New("refresh token revoked")
	// 2xx

	// ErrorNoContent will throw if resource is not found but query is correct
	ErrorNoContent           = errors.New("data not found")
	ErrorAccessTokenExpired  = errors.New("access token is expired")
	ErrorRefreshTokenExpired = errors.New("refresh token is expired")
	ErrorInvalidToken        = errors.New("token is invalid")
)

const (
	SUCCESS               = "success"
	INTERNAL_SERVER_ERROR = "internal_server_error"
	DATA_NOT_EXIST        = "data_not_exist"
	RESULT_NOT_FOUND      = "result_not_found"
	DUPLICATE_DATA        = "duplicate_data"
	BAD_REQUEST           = "bad_request"
	UNAUTHORIZE           = "unauthorized"
	REFRESH_TOKEN_REVOKED = "refresh_token_revoked"
	NO_CONTENT            = "no_content"
	ACCESS_TOKEN_EXPIRED  = "access_token_expired"
	REFRESH_TOKEN_EXPIRED = "refresh_token_expired"
)

// GetStatusCode for handle status error
func GetStatusCode(err error) (int, string) {
	if err == nil {
		return http.StatusOK, SUCCESS
	}
	switch err {
	case ErrorResultNotFound:
		return http.StatusOK, RESULT_NOT_FOUND
	case ErrorBadRequest, ErrorInvalidToken:
		return http.StatusBadRequest, BAD_REQUEST
	case ErrorNotFound:
		return http.StatusNotFound, DATA_NOT_EXIST
	case ErrorDuplicateData:
		return http.StatusConflict, DUPLICATE_DATA
	case ErrorUnauthorized, ErrorState, ErrorBearer, ErrorInvalidBearerToken:
		return http.StatusUnauthorized, UNAUTHORIZE
	case ErrorRefreshTokenRevoked:
		return http.StatusUnauthorized, REFRESH_TOKEN_REVOKED
	case ErrorNoContent:
		return http.StatusNoContent, NO_CONTENT
	case ErrorAccessTokenExpired:
		return http.StatusUnauthorized, ACCESS_TOKEN_EXPIRED
	default:
		return http.StatusInternalServerError, INTERNAL_SERVER_ERROR
	}
}

func HashToStr(authCode string) string {
	hashed := sha256.Sum256([]byte(authCode))
	return fmt.Sprintf("%x", hashed[:])
}

func HashSHA1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// Check if an interface is nil or there is some value in it
func IsInterfaceNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

// Hash sha1
func HashSessionID(orderId string) string {
	return HashSHA1(orderId)
}

func GetCurrentFunctionName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Func.Name()
}

// PtrToValue converts pointer to value, if pointer == nil the value would be the zero value of the type
// else the value stored on the pointer would be return.
func PtrToValue[T any](valuePtr *T) T {
	var v T
	if valuePtr != nil {
		v = *valuePtr
	}
	return v
}

func GenerateUniqueUsernameFromEmail(email string) string {
	username := strings.Split(email, "@")[0]
	randInt, _ := rand.Int(rand.Reader, big.NewInt(999))
	strInt := strconv.Itoa(int(randInt.Int64()))
	return fmt.Sprintf("%s%s", username, strInt)
}

func EqualAny[T comparable](value T, targets ...T) bool {
	for _, t := range targets {
		if value == t {
			return true
		}
	}
	return false
}

// GenerateStateValue generates a random "state" value of the specified length.
// The generated value helps prevent cross-site request forgery (CSRF) attacks.
//
// Parameters:
//   - length: The desired length of the state value (default: 15).
//
// Returns:
//   - string: The random "state" value.
func GenerateStateValue(length int) string {
	// set default to 15 if length is not greater than 0
	if length < 0 {
		length = 15
	}
	b := make([]byte, length)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}

// Define a function to check if a Unix timestamp (int) is expired.
func IsExpired(unixTimestamp int64) bool {
	// Convert the integer time to a time.Time object.
	expirationTime := time.Unix(unixTimestamp, 0)

	// Get the current time.
	currentTime := time.Now()

	// Compare the two times.
	return currentTime.After(expirationTime)
}

func StringContains(mainStr string, subStr string) bool {
	// Check if the mainString contains the subString
	return strings.Contains(mainStr, subStr)
}

// convert seconds in times (int64) to time.Time
// example: 86400 seconds -> current time + 1days
func ConvertSecondsInt64ToTime(seconds int64) time.Time {
	current := time.Now()
	return current.Add(time.Second * time.Duration(seconds))
}

// ConvertStrToInt converts a string to an integer, using a default value (n) if the conversion fails.
// If the conversion is successful, the result is returned; otherwise, the default value is returned.
func ConvertStrToInt(str string, n int) int {
	size := n
	s, err := strconv.Atoi(str)
	if err == nil {
		size = s
	}

	return size
}

// GetInt64Value retrieves an int64 value from the map
func GetInt64Value(data map[string]interface{}, key string) int64 {
	if value, ok := data[key]; ok {
		switch v := value.(type) {
		case float64:
			return int64(v)
		case int64:
			return v
		default:
			panic(fmt.Sprintf("Unexpected type for key '%s': %T", key, value))
		}
	}
	return 0
}

// GetBoolValue retrives a bool value from a map
func GetBoolValue(data map[string]interface{}, key string) bool {
	if value, ok := data[key]; ok {
		if str, ok := value.(bool); ok {
			return str
		}
	}
	return false
}

// GetStringValue retrieves a string value from the map
func GetStringValue(data map[string]interface{}, key string) string {
	if value, ok := data[key]; ok {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

// GetTimeValue retrieves a time.Time value from the map
func GetTimeValue(data map[string]interface{}, key string) time.Time {
	if value, ok := data[key]; ok {
		switch v := value.(type) {
		case string:
			// Assuming the time is represented as a string in RFC3339Nano format
			t, err := time.Parse(time.RFC3339Nano, v)
			if err != nil {
				panic(fmt.Sprintf("Error parsing time for key '%s': %v", key, err))
			}
			return t
		default:
			panic(fmt.Sprintf("Unexpected type for key '%s': %T", key, value))
		}
	}
	return time.Time{}
}

// GetFloat64Value retrieves a float64 value from the map
func GetFloat64Value(data map[string]interface{}, key string) float64 {
	if value, ok := data[key]; ok {
		switch v := value.(type) {
		case float64:
			return v
		case int64:
			return float64(v)
		default:
			panic(fmt.Sprintf("Unexpected type for key '%s': %T", key, value))
		}
	}
	return 0
}

func GetArrayOfInt64(data map[string]interface{}, key string) []int64 {
	value, ok := data[key]
	if !ok {
		return nil
	}
	arr, ok := value.([]interface{})
	if !ok {
		fmt.Printf("%+v\n", value)
		return nil
	}
	res := ConvertSliceToInt64Slice(arr)
	return res
}

func ConvertSliceToInt64Slice(data []interface{}) []int64 {
	result := make([]int64, 0, len(data))
	for _, value := range data {
		switch v := value.(type) {
		case float64:
			result = append(result, int64(v))
		case int64:
			result = append(result, v)
		default:
			panic(fmt.Sprintf("Cannot convert value of Type '%T'to %T", v, result[0]))
		}
	}
	return result
}

// isValidSlug checks if the slug is in a valid format (alphanumeric and dashes)
func IsValidSlug(slug string) bool {
	validSlugPattern := `^/?[a-zA-Z0-9-_]+(/?[a-zA-Z0-9-_]+)*$`
	matched, _ := regexp.MatchString(validSlugPattern, slug)
	return matched
}

func GetLastSlug(paramSlug string) string {
	// Check if the paramSlug is empty
	if paramSlug == "" {
		return ""
	}

	// Split the string by "/"
	slugParts := strings.Split(paramSlug, "/")

	// Check if there are no parts after splitting
	if len(slugParts) == 0 {
		return ""
	}

	// Return the last part of the slice
	return slugParts[len(slugParts)-1]
}
