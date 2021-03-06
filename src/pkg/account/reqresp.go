package account

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	// CreateUserRequest - comment for export
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// CreateUserResponse - comment for export
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
