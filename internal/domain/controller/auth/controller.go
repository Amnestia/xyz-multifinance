package auth

import (
	"fmt"
	"net/http"
	"strings"

	authmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/auth"
	"github.com/amnestia/xyz-multifinance/internal/domain/service"
	"github.com/amnestia/xyz-multifinance/pkg/json"
	"github.com/amnestia/xyz-multifinance/pkg/response"
)

// Controller handler for this domain
type Controller struct {
	AuthSvc service.AuthServicer
}

// Auth authentication login handler
func (c *Controller) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response := response.NewResponse(ctx)

	req := &authmodel.ConsumerAuthRequest{}
	err := json.Decode(r.Body, &req)
	if err != nil {
		response.SetErrorResponse(http.StatusBadRequest, err, "Invalid Request").WriteJSON(w)
		return
	}
	if strings.TrimSpace(req.NIK) == "" {
		response.SetErrorResponse(http.StatusBadRequest, fmt.Errorf("NIK is required")).WriteJSON(w)
		return
	}
	if strings.TrimSpace(req.Password) == "" {
		response.SetErrorResponse(http.StatusNotFound, fmt.Errorf("Password is required")).WriteJSON(w)
		return
	}
	ret := c.AuthSvc.Auth(ctx, req)
	if ret.Error != nil {
		response.SetErrorResponse(ret.HTTPCode, ret.Error).WriteJSON(w)
		return
	}
	response.SetResponse(ret.HTTPCode, ret.TokenData, "Successfully logged in").WriteJSON(w)
}

// Register register new account
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response := response.NewResponse(ctx)

	req := &authmodel.RegisterRequest{}
	err := json.Decode(r.Body, &req)
	if err != nil {
		response.SetErrorResponse(http.StatusBadRequest, err).WriteJSON(w)
		return
	}
	err = validateRegister(req)
	if err != nil {
		response.SetErrorResponse(http.StatusBadRequest, err).WriteJSON(w)
		return
	}
	ret := c.AuthSvc.Register(ctx, req)
	if ret.Error != nil {
		response.SetErrorResponse(ret.HTTPCode, ret.Error).WriteJSON(w)
		return
	}
	response.SetResponse(ret.HTTPCode, nil, "Successfully registered").WriteJSON(w)
}

// RegisterNewPartner register new account
func (c *Controller) RegisterNewPartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response := response.NewResponse(ctx)

	req := &authmodel.Partner{}
	err := json.Decode(r.Body, &req)
	if err != nil {
		response.SetErrorResponse(http.StatusBadRequest, err).WriteJSON(w)
		return
	}
	ret := c.AuthSvc.CreateNewPartner(ctx, req)
	if ret.Error != nil {
		response.SetErrorResponse(ret.HTTPCode, ret.Error).WriteJSON(w)
		return
	}
	response.SetResponse(ret.HTTPCode, map[string]string{"api_key": ret.APIKey, "client_id": ret.ClientID}, "Successfully created").WriteJSON(w)
}
