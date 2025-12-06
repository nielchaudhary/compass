package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Endpoints struct {
	ID          primitive.ObjectID `bson:"_id"`
	ServiceID   string             `bson:"serviceId"`
	ServiceName string             `bson:"serviceName"`
	Method      string             `bson:"method"`
	Endpoint    string             `bson:"endpoint"`
}
