package model

import "context"

const ignoreKey = "ignoreKey"

// IgnoreChars is stroe the ignore chars into firestore
type IgnoreChars struct {
	DatastoreMethod
	Chars []string `json:"chars"`
}

// Get is get from datastore
func (ignoreChars *IgnoreChars) Get(ctx context.Context) {
	ignoreChars.GetDatastore(ctx, ignoreKey, ignoreChars)
}

// AddAndPut is put into datastore
func (ignoreChars *IgnoreChars) AddAndPut(ctx context.Context) {
	ignoreChars.PutDatastore(ctx, ignoreKey, ignoreChars)
}
