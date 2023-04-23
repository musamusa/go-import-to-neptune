package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"thiapp/neptune_loader/ent"
	"thiapp/neptune_loader/lib"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalln("no .env file to load")
	}
	downloadSrcUrl := os.Getenv("MITRE_BUNDLE_URL")
	stripDollar := os.Getenv("STRIP_DOLLAR") == "true"
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" || downloadSrcUrl == "" {
		logrus.Fatalln("DB_URL or MITRE_BUNDLE_URL not set, kindly set both continue")
	}

	db, err := ent.Open("gremlin", dbURL)

	if err != nil {
		logrus.Fatalln(err)
	}

	dataBundle, err := lib.GetBundleFromUrl(downloadSrcUrl)

	if err != nil {
		logrus.Fatalln(err)
	}

	ctx := context.Background()
	_, _ = db.AttackPattern.Delete().Exec(ctx) // reset previous data
	lib.PushAttackPatternsToDB(ctx, db, dataBundle.Objects, stripDollar)
}
