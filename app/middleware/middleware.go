package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/erwinwahyura/go-boilerplate/app/model"
	"github.com/erwinwahyura/go-boilerplate/app/model/constant"
	"github.com/erwinwahyura/go-boilerplate/utils"
	"github.com/erwinwahyura/go-boilerplate/utils/jwt"
	"github.com/justinas/nosurf"
	"github.com/spf13/viper"
)

type contextKey string

const (
	isAuthenticatedContextKey = contextKey("isAuthenticated")
	uid                       = contextKey("uid")
	token                     = contextKey("token")
	PaginationPage            = contextKey("page")
	PaginationSize            = contextKey("size")
	userIdKey                 = contextKey("userId")
	platformKey               = contextKey("platform")
)

type (

	// LogHTTP to show fields of log
	LogHTTP struct {
		RequestID string `json:"request_id,omitempty"`
		Method    string `json:"method,omitempty"`
		URI       string `json:"uri,omitempty"`
		IP        string `json:"ip,omitempty"`
		RemoteIP  string `json:"remote_ip,omitempty"`
		Host      string `json:"host,omitempty"`
		Status    int    `json:"status,omitempty"`
		Size      int64  `json:"size,omitempty"`
		UserAgent string `json:"user_agent,omitempty"`
		Header    string `json:"header,omitempty"`
		Body      string `json:"body,omitempty"`
	}

	// GoMiddleware struct of middleware
	GoMiddleware struct {
		Config model.Config
		// add dependecies to logging the log, perhaps store into mongodb?
	}
)

// InitMiddleware will initialize the middleware handler
func InitMiddleware(config model.Config) *GoMiddleware {
	return &GoMiddleware{
		Config: config,
	}
}

func (m *GoMiddleware) SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// default-src 'self' -> defines the sources from which content can be loaded. self indicates that content
		// can only be loaded from the same origin as the page.
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func IsAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)
	if !ok {
		return false
	}

	return isAuthenticated
}

func (m *GoMiddleware) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				ServerError(w, fmt.Errorf("%s", err), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (m *GoMiddleware) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthenticated(r) {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")

		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return csrfHandler
}

func ServerError(w http.ResponseWriter, err error, code int) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.Output(2, trace)
	w.Header().Set("Content-Type", "application/json")

	http.Error(w, http.StatusText(code), code)
}

func (m *GoMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.Contains(header, "Bearer") {
			model.MapBaseResponse(w, r, utils.ErrorBearer.Error(), nil, nil, utils.ErrorUnauthorized)
			return
		}

		// TODO: add redis to check is token revoked or not for better performance

		tokenStr := strings.Replace(header, "Bearer ", "", -1)
		jwt := jwt.NewJWT()
		claims, err := jwt.ValidateToken(tokenStr, viper.GetString("SECRETKEY"))
		if err != nil {
			if err.Error() == errors.New("token has invalid claims: token is expired").Error() {
				model.MapBaseResponse(w, r, utils.ACCESS_TOKEN_EXPIRED, nil, nil, utils.ErrorAccessTokenExpired)
				return
			}
			model.MapBaseResponse(w, r, utils.ErrorUnauthorized.Error(), nil, nil, utils.ErrorUnauthorized)
			return
		}
		if claims["id"] != "" {
			// Map App Context
			id, ok := claims["id"].(string)
			if !ok {
				id = ""
			}
			issuer, ok := claims["iss"].(string)
			if !ok {
				issuer = ""
			}
			appContext := model.AppContext{
				Context: r.Context(),
				UID:     id,
				Token:   tokenStr,
				Issuer:  issuer,
			}

			// Set App Context
			r = r.WithContext(appContext)
		}

		next.ServeHTTP(w, r)
	})
}

func (m *GoMiddleware) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestLog := MapLogRequest(w, r)
		fmt.Println(requestLog)

		next.ServeHTTP(w, r)
	})
}

// MapLogRequest for map log request
func MapLogRequest(w http.ResponseWriter, r *http.Request) string {
	headerByte, _ := json.Marshal(r.Header)

	// Read the content
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}
	// Use the content
	var req interface{}
	json.Unmarshal(bodyBytes, &req)
	bodyBytes, _ = json.Marshal(req)

	// Restore the io.ReadCloser to its original state
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return fmt.Sprint("[IN_REQUEST: ", r.URL, "] REQUEST_ID: ", r.Header.Get(constant.RequestID), " HEADER:", string(headerByte))
}

func (m *GoMiddleware) Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		page := r.URL.Query().Get("page")
		size := r.URL.Query().Get("size")

		ctx := context.WithValue(r.Context(), PaginationPage, page)
		ctx = context.WithValue(ctx, PaginationSize, size)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

// userId (numerical)
type UserId int64

func UserIdFromContext(ctx context.Context) (UserId, bool) {
	userId, ok := ctx.Value(userIdKey).(UserId)
	return userId, ok
}

type Platform string

func PlatformFromContext(ctx context.Context) (Platform, bool) {
	platform, ok := ctx.Value(platformKey).(Platform)
	return platform, ok
}
