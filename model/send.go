package model

import "context"

const sendKey = "sendKey"

// Send is send mail or not
type Send struct {
	DatastoreMethod
	Count int `json:"count"`
}

// Get is get from datastore
func (send *Send) Get(ctx context.Context) {
	send.GetDatastore(ctx, sendKey, send)
}

// Reset is reset to 0
func (send *Send) Reset(ctx context.Context) {
	send.Count = 0
	send.PutDatastore(ctx, sendKey, send)
}

// AddAndPut is send.Count++ and put into datastore
func (send *Send) AddAndPut(ctx context.Context) {
	send.Count++
	send.PutDatastore(ctx, sendKey, send)
}
