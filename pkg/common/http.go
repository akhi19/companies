package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akhi19/companies/configs"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/golang-jwt/jwt"
)

type Code int

const (
	serverErrorCode  Code = 500
	paramMissingCode Code = 400
	BadRequestCode   Code = 404
	Unauthorized     Code = 401
)

const (
	InvalidRequestMsg = "Invalid Request"
)

type Entity struct {
	HttpStatusCode int
	Message        string
	Code           Code
}

func (ue *Entity) Error() string {
	return ue.Message
}

func ValidationError(msg string) error {
	return &Entity{
		HttpStatusCode: http.StatusBadRequest,
		Message:        msg,
		Code:           paramMissingCode,
	}
}

func InternalServerError() error {
	return &Entity{
		HttpStatusCode: http.StatusInternalServerError,
		Message:        "Something Went Wrong",
		Code:           serverErrorCode,
	}
}

func BadRequest(code Code, msg string) error {
	return &Entity{
		HttpStatusCode: http.StatusBadRequest,
		Message:        msg,
		Code:           code,
	}
}

func SendHttpError(ctx context.Context, w http.ResponseWriter, err error) {
	httpStatusCode := http.StatusInternalServerError
	code := serverErrorCode
	message := err.Error()

	if errVal, ok := err.(*Entity); ok {
		httpStatusCode = errVal.HttpStatusCode
		message = errVal.Message
		code = errVal.Code
	}

	resp := map[string]interface{}{
		"message": message,
		"code":    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(resp)
}

func SendHttpResponse(
	responseWriter http.ResponseWriter,
	statusCode int,
	response interface{},
) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	json.NewEncoder(responseWriter).Encode(response)
}

type emptyHttpResponse struct{}

func SendEmptyHttpResponse(
	responseWriter http.ResponseWriter,
	statusCode int,
) {
	responseData := emptyHttpResponse{}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	json.NewEncoder(responseWriter).Encode(responseData)
}

func RecoverHandler(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Panic(err)
				serverErr := InternalServerError()
				SendHttpError(r.Context(), w, serverErr)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

func AuthenticateHandler(
	next http.HandlerFunc,
	role domain.Role,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if role == domain.RoleAny {
			next.ServeHTTP(w, r)
			return
		}
		if r.Header["Token"] == nil {
			SendHttpResponse(
				w,
				int(Unauthorized),
				"No Token Found",
			)
			return
		}
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return configs.GetConfig().JWTSecret, nil
		})
		if err != nil {
			SendHttpResponse(
				w,
				int(Unauthorized),
				"Invalid Token",
			)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if role == domain.RoleAdmin && claims["role"] != string(role) {
				SendHttpResponse(
					w,
					int(Unauthorized),
					"Role not authorized for operation",
				)
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}

func HttpRequestHandler(
	next http.HandlerFunc,
	role domain.Role,
) http.HandlerFunc {
	next = AuthenticateHandler(next, role)
	next = RecoverHandler(next)
	return next
}
