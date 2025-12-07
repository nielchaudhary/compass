package scheduler

// will be used for pinging the /health APIs to check if the service is working well or not
import (
	core "github.com/nielchaudhary/compass/internal/core/services"
	"github.com/nielchaudhary/compass/pkg/logger"
)

type UptimeSchedulerResp struct {
	Health string `bson:"health"`
	Error  string `bson:"error"`
}

func UptimeScheduler() ([]UptimeSchedulerResp, error) {
	log := logger.GetLogger()
	defer log.Sync()

	endpoints, fetchEndpointFromDBError := core.FetchEndpointsDataFromDB("endpoints")
	if fetchEndpointFromDBError != nil {
		log.Error("Could not fetch endpoints from DB due to: ", fetchEndpointFromDBError)
		return nil, fetchEndpointFromDBError
	}

	var uptimeResults []UptimeSchedulerResp

	for i := 0; i < len(endpoints); i++ {
		pingServiceResp, pingServiceError := PingEndpoint(endpoints[i].Endpoint, endpoints[i].Method)

		if pingServiceError != nil {
			uptimeResults = append(uptimeResults, UptimeSchedulerResp{
				Health: "",
				Error:  pingServiceError.Error(),
			})
			continue
		}

		uptimeResults = append(uptimeResults, UptimeSchedulerResp{
			Health: pingServiceResp.ServiceStatus,
			Error:  "",
		})
	}

	return uptimeResults, nil
}
