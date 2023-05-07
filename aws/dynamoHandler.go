package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func ConsultaDynamo(id string) []byte {

	my_table := "deploy-status"

	// Obtem as configurações de credenciais da AWS
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	dynamoDbClient := dynamodb.NewFromConfig(cfg)

	// Define os query parameters
	queryParameters := &dynamodb.QueryInput{
		TableName:              aws.String(my_table),
		KeyConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{Value: id},
		},
	}

	// Exetuta o envio da query acima
	result, err := dynamoDbClient.Query(context.TODO(), queryParameters)
	if err != nil {
		fmt.Println("Error querying DynamoDB table:", err)
	}

	mapContainingAllTheItemsWithId := result.Items

	// Valida se foram retornados itens
	if len(mapContainingAllTheItemsWithId) == 0 {
		fmt.Println("No items found for the given ID")
	}

	// Define o slice para ontenção e ordenação dos datetimes
	sliceContainingOnlyTheDatetimes := make([]string, len(mapContainingAllTheItemsWithId))

	// Preenche o slice somente com os valores de datetime
	for i, m := range mapContainingAllTheItemsWithId {
		sliceContainingOnlyTheDatetimes[i] = fmt.Sprintf("%v", m["datetime"])
	}

	// Ordena o slice em ordem descrescente de eventos
	// ou seja, do mais recente para o evento inicial (mais antigo)
	sort.Slice(sliceContainingOnlyTheDatetimes, func(i, j int) bool {
		return sliceContainingOnlyTheDatetimes[i] > sliceContainingOnlyTheDatetimes[j]
	})

	var latestEventMap map[string]types.AttributeValue

	// Localiza o map que contém o menor tempo (obtido acima)
	for _, m := range mapContainingAllTheItemsWithId {
		if fmt.Sprintf("%v", m["datetime"]) == sliceContainingOnlyTheDatetimes[0] {
			latestEventMap = m
			break
		}
	}

	latestEventJsonBytes, err := json.Marshal(latestEventMap)
	if err != nil {
		// handle this error for me
	}

	return (latestEventJsonBytes)
}
