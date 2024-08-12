package transaction

import (
	"net/http"

	transactionmodel "github.com/amnestia/xyz-multifinance/internal/domain/model/transaction"
	"github.com/amnestia/xyz-multifinance/internal/domain/service"
	"github.com/amnestia/xyz-multifinance/pkg/json"
	"github.com/amnestia/xyz-multifinance/pkg/logger"
	"github.com/amnestia/xyz-multifinance/pkg/response"
)

// Controller handler for this domain
type Controller struct {
	TransactionSvc service.TransactionServicer
}

func (c *Controller) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response := response.NewResponse(ctx)

	req := &transactionmodel.TransactionRequest{}
	b, _ := json.MarshalIndent(req)
	logger.Logger.Info().Msg(string(b))
	err := json.Decode(r.Body, &req)
	if err != nil {
		response.SetErrorResponse(http.StatusBadRequest, err, "Invalid Request").WriteJSON(w)
		return
	}
	ret := c.TransactionSvc.CreateNewTransaction(ctx, req)
	if ret.Error != nil {
		response.SetErrorResponse(ret.HTTPCode, ret.Error).WriteJSON(w)
		return
	}
}
