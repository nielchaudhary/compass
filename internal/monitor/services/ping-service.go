package monitor

import (
	"fmt"
	"net/http"

	"github.com/nielchaudhary/compass/pkg/logger"
)

type pingEndpointResp struct {
	Message       string `json:"message"`
	ServiceStatus string `json:"serviceStatus"`
	RespStatus    int    `json:"respStatus"` // Fixed missing closing quote
}

func PingEndpoint(endpoint string, method string) (pingEndpointResp, error) {
	log := logger.GetLogger("services/ping-service.go")
	log.Info("Pinging the following endpoint: ", endpoint, " with Method: ", method)

	request, newRequestError := http.NewRequest(method, endpoint, nil)
	if newRequestError != nil {
		log.Error("Error creating http request: ", newRequestError)
		errorResp := pingEndpointResp{
			Message:       "Error creating http request: " + newRequestError.Error(),
			ServiceStatus: "Error",
			RespStatus:    http.StatusBadRequest,
		}
		return errorResp, newRequestError
	}
	request.Header.Set("Content-Type", "application/json")

	resp, sendRequestError := http.DefaultClient.Do(request)
	if sendRequestError != nil {
		log.Error("Error while sending the request through default client: ", sendRequestError)
		errorResp := pingEndpointResp{
			Message:       "Error sending request: " + sendRequestError.Error(),
			ServiceStatus: "Error",
			RespStatus:    http.StatusBadGateway,
		}
		return errorResp, sendRequestError
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		clientResp := pingEndpointResp{
			Message:       "Ping Successful, Received 200",
			ServiceStatus: "Healthy",
			RespStatus:    http.StatusOK,
		}
		return clientResp, nil
	}

	clientErrorResp := pingEndpointResp{
		Message:       fmt.Sprintf("Pinged, Received status: %d", resp.StatusCode),
		ServiceStatus: "Unhealthy",
		RespStatus:    resp.StatusCode,
	}
	return clientErrorResp, nil
}
