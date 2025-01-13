package main

import (
	"context"

	"github.com/gunktp20/digital-hubx-be/pkg/bucket"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
	"github.com/gunktp20/digital-hubx-be/pkg/database"
	"github.com/gunktp20/digital-hubx-be/server"
)

// @title digital-hubx
// @version 1.0
// @description digital hubx api
// @contact.name API Support
// @contact.email support@example.com
// @host localhost:3000
// @BasePath /api
func main() {
	ctx := context.Background()
	_ = ctx

	conf := config.GetConfig()

	db := database.NewGormPostgresDatabase(ctx, conf)
	bucketClient := bucket.NewBucketClient(ctx, conf)

	server.NewFiberServer(conf, db, bucketClient).Start()
}
