package model

import (
	"context"
	"reflect"
	"strings"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// DatastoreMethod define some public datastore methods
type DatastoreMethod struct {
}

// PutDatastore is store to Datastore
func (dm *DatastoreMethod) PutDatastore(ctx context.Context, key string, value interface{}) {
	datastoreKey := datastore.NewKey(ctx, getType(value), key, 0, nil)

	log.Infof(ctx, "PUT into Datastore with Key: %v", key)
	if _, err := datastore.Put(ctx, datastoreKey, value); err != nil {
		log.Errorf(ctx, "PUT into Datastore failed %v", err)
	}
	return
}

// GetDatastore is get from Datastore
func (dm *DatastoreMethod) GetDatastore(ctx context.Context, key string, value interface{}) {
	datastoreKey := datastore.NewKey(ctx, getType(value), key, 0, nil)

	log.Infof(ctx, "GET Datastore with Key: %v", key)
	if err := datastore.Get(ctx, datastoreKey, value); err != nil {
		if !strings.HasPrefix(err.Error(), `datastore: cannot load field`) {
			log.Errorf(ctx, "GET Datastore failed %v", err)
		}
	}
	return
}

func getType(myvar interface{}) string {
	t := reflect.TypeOf(myvar)
	if t.Kind() == reflect.Ptr {
		return `gae#DMS#` + t.Elem().Name()
	}
	return `gae#DMS#` + t.Name()
}
