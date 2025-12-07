package scheduler

// will be used for pinging the /health APIs to check if the service is working well or not
import (
	core "github.com/nielchaudhary/compass/internal/core/services"
	"github.com/nielchaudhary/compass/pkg/logger"
)

func UptimeScheduler() ([]core.Endpoints, error) {

	log := logger.GetLogger()
	defer log.Sync()

	endpoints, fetchEndpointFromDBError := core.FetchEndpointsDataFromDB("endpoints")

	if fetchEndpointFromDBError != nil {
		log.Error("Could not fetch endpoints from DB due to : ", fetchEndpointFromDBError)
		return nil, fetchEndpointFromDBError
	}

}
