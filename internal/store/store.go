package store

import "context"

type Store = context.Context 

type contextKey string
const mapKey contextKey = "storageContext"
