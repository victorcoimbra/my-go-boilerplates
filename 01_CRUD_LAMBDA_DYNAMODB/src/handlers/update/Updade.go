package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID               string  `json:"id"`
	Date             string  `json:"Date"`
	Ticket           string  `json:"Ticket"`
	Type             string  `json:"Type"`
	SellPrice        float64 `json:"SellPrice"`
	Qt               int     `json:"Qt"`
	TaxLiquidacao    float64 `json:"TaxLiquidacao"`
	TaxRegistro      float64 `json:"TaxRegistro"`
	TaxTermo         float64 `json:"TaxTermo"`
	TaxANA           float64 `json:"TaxANA"`
	TaxEmolumentos   float64 `json:"TaxEmolumentos"`
	TaxOperacional   float64 `json:"TaxOperacional"`
	TaxExecucao      float64 `json:"TaxExecucao"`
	TaxCustodia      float64 `json:"TaxCustodia"`
	TaxImpostos      float64 `json:"TaxImpostos"`
	TaxIRRF          float64 `json:"TaxIRRF"`
	TaxOutros        float64 `json:"TaxOutros"`
	TotalLiquidoNota float64 `json:"TotalLiquidoNota"`
	TaxTotal         float64 `json:"TaxTotal"`
	TotalCalculo     float64 `json:"TotalCalculo"`
	AvgPriceNota     float64 `json:"AvgPriceNota"`
	AvgPriceCalculo  float64 `json:"AvgPriceCalculo"`
	DiffNota         float64 `json:"DiffNota"`
	DiffPrice        float64 `json:"DiffPrice"`
	CreatedAt        string  `json:"created_at"`
}

// type Todo struct {
// 	ID          string  `json:"id"`
// 	Description string 	`json:"description"`
// 	Done        bool   	`json:"done"`
// 	CreatedAt   string 	`json:"created_at"`
// }

var ddb *dynamodb.DynamoDB

func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) // Create DynamoDB client
	}
}

func AddTodo(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("AddTodo")

	var (
		id        = uuid.Must(uuid.NewV4(), nil).String()
		tableName = aws.String(os.Getenv("TODOS_TABLE_NAME"))
	)

	// Initialize todo
	todo := &Todo{
		ID:        id,
		CreatedAt: time.Now().String(),
	}

	// Parse request body
	json.Unmarshal([]byte(request.Body), todo)

	// Write to DynamoDB
	item, _ := dynamodbattribute.MarshalMap(todo)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tableName,
	}
	if _, err := ddb.PutItem(input); err != nil {
		return events.APIGatewayProxyResponse{ // Error HTTP response
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	} else {
		body, _ := json.Marshal(todo)
		return events.APIGatewayProxyResponse{ // Success HTTP response
			Body:       string(body),
			StatusCode: 200,
		}, nil
	}
}

func main() {
	lambda.Start(AddTodo)
}
