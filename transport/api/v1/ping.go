package v1

import (
	"log/slog"
	"net/http"
	"newsletter_backend_api/transport/util"
)

func (h *Handler) Ping(writer http.ResponseWriter, request *http.Request) {
	slog.Info("request received", request)
	slog.Info("ping triggered")
	slog.Error("ping triggered")
	util.WriteResponse(writer, http.StatusOK, "pong")

}
