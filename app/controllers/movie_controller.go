package controllers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/revel/revel"
)

// Movie struct

type Movie struct {
	ID     string  `json:"id"`
	Year   int     `json:"year"`
	Title  string  `json:"title"`
	Plot   string  `json:"plot"`
	Rating float64 `json:"rating"`
}

// MovieController handles movie API

type MovieController struct {
	revel.Controller
}

// ListMovies returns all movies

// ...existing code...

// filepath: [movie_controller.go](http://_vscodecontentref_/0)
// ...existing code...

// ...existing code...

func (c MovieController) ListMovies() revel.Result {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	input := &dynamodb.ScanInput{
		TableName: aws.String("Movies"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		return c.RenderJSON(map[string]string{"error": err.Error()})
	}

	var movies []Movie
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &movies)
	if err != nil {
		return c.RenderJSON(map[string]string{"error": err.Error()})
	}

	return c.RenderJSON(map[string]interface{}{"movies": movies})
}
