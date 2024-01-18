package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Games struct {
	Name   string  // "Super Mario Bros", "Legend of Zelda", etc.
	Year   int     // 1998, 2001, etc.
	Rating string  // E, T, M, etc.
	Price  float64 // 59.99, 49.99, etc.
}

func getTable() string {
	return os.Getenv("TABLE_NAME")
}

func newSess() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
}

func AddItem(game Games) error {
	sess, err := newSess()
	if err != nil {
		return err
	}

	db := dynamo.New(sess)
	table := db.Table(getTable())

	newGame := Games{
		Name:   game.Name,
		Year:   game.Year,
		Rating: game.Rating,
		Price:  game.Price,
	}

	perr := table.Put(&newGame).Run()

	if perr != nil {
		return perr
	}

	return nil
}
