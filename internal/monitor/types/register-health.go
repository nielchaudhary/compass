package types

type RegisterHealthRequestBody struct {
	ID          string `json:"id"`
	Endpoint    string `json:"endpoint"`
	Method      string `json:"method"` //POST, GET, PUT, PATCH
	ServiceName string `json:"serviceName"`
}
