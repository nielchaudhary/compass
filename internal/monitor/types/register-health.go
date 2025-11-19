package types

type RequestMethod string

const (
	POST   RequestMethod = "POST"
	GET    RequestMethod = "GET"
	PUT    RequestMethod = "PUT"
	PATCH  RequestMethod = "PATCH"
	DELETE RequestMethod = "DELETE"
)

type RegisterHealthRequestBody struct {
	ID          string        `json:"id"`
	Endpoint    string        `json:"endpoint"`
	Method      RequestMethod `json:"method"` //POST, GET, PUT, PATCH, DELETE
	ServiceName string        `json:"serviceName"`
}
