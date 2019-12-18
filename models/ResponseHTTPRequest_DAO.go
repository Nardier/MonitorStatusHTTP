package models

import "gopkg.in/mgo.v2/bson"

type ResponseDataBody_OBJ struct {

	ID                  bson.ObjectId `bson:"_id" json:"id"`
	Response_Body       string        `bson:"response_Body" json:"response_body"`
	Response_Time       float64       `bson:"response_Time" json:"response_time"`
	Response_StatusCode int           `bson:"response_StatusCode" json:"response_status_code"`
	Response_Status     string        `bson:"response_Status" json:"response_status"`
	TEOM                string        `bson:"tEOM" json:"teom"`
	Request_Url         string        `bson:"request_Url" json:"request_url"`
}
