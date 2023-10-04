package dynamodb

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.resf.org/peridot/base/go/awsutils"
	"go.resf.org/peridot/base/go/kv"
	basepb "go.resf.org/peridot/base/go/pb"
	"google.golang.org/protobuf/proto"
	"strings"
	"time"
)

type DynamoDB struct {
	db        *dynamodb.DynamoDB
	tableName string
}

// New creates a new DynamoDB storage backend.
func New(endpoint string, tableName string) (*DynamoDB, error) {
	awsCfg := &aws.Config{}
	awsutils.FillOutConfig(awsCfg)

	if endpoint != "" {
		awsCfg.Endpoint = aws.String(endpoint)
	}

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(sess)

	// Create the table if it doesn't exist.
	// First check if the table exists.
	_, err = svc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		_, err = svc.CreateTable(&dynamodb.CreateTableInput{
			TableName: aws.String(tableName),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("Key"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("Path"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("Key"),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String("Path"),
					KeyType:       aws.String("RANGE"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		})
		if err != nil {
			return nil, err
		}
	}

	return &DynamoDB{
		db:        svc,
		tableName: tableName,
	}, nil
}

func (d *DynamoDB) Get(ctx context.Context, key string) (*kv.Pair, error) {
	trimmed := strings.TrimPrefix(key, "/")
	parts := strings.Split(trimmed, "/")
	if len(parts) < 2 {
		return nil, kv.ErrNoNamespace
	}
	ns := parts[0]

	result, err := d.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(d.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Key": {
				S: aws.String(ns),
			},
			"Path": {
				S: aws.String(strings.Join(parts[1:], "/")),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, kv.ErrNotFound
	}

	return &kv.Pair{
		Key:   *result.Item["Key"].S,
		Value: result.Item["Value"].B,
	}, nil
}

func (d *DynamoDB) Set(ctx context.Context, key string, value []byte) error {
	trimmed := strings.TrimPrefix(key, "/")
	parts := strings.Split(trimmed, "/")
	if len(parts) < 2 {
		return kv.ErrNoNamespace
	}
	ns := parts[0]

	_, err := d.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(d.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"Key": {
				S: aws.String(ns),
			},
			"Path": {
				S: aws.String(strings.Join(parts[1:], "/")),
			},
			"Value": {
				B: value,
			},
		},
	})
	return err
}

func (d *DynamoDB) Delete(ctx context.Context, key string) error {
	trimmed := strings.TrimPrefix(key, "/")
	parts := strings.Split(trimmed, "/")
	if len(parts) < 2 {
		return kv.ErrNoNamespace
	}
	ns := parts[0]

	_, err := d.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(d.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Key": {
				S: aws.String(ns),
			},
			"Path": {
				S: aws.String(strings.Join(parts[1:], "/")),
			},
		},
	})
	return err
}

func (d *DynamoDB) RangePrefix(ctx context.Context, prefix string, pageSize int32, pageToken string) (*kv.Query, error) {
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}

	// Check if there is a page token.
	var fromKey string
	var fromPath string
	if pageToken != "" {
		// Get the page token from the database.
		res, err := d.Get(ctx, fmt.Sprintf("/_internals/page_tokens/%s", pageToken))
		if err != nil {
			if errors.Is(err, kv.ErrNotFound) {
				return nil, kv.ErrPageTokenNotFound
			}
			return nil, err
		}

		// Parse the page token.
		pt := &basepb.DynamoDbPageToken{}
		err = proto.Unmarshal(res.Value, pt)
		if err != nil {
			return nil, err
		}

		fromKey = pt.LastKey
		fromPath = pt.LastPath
	}

	trimmed := strings.TrimPrefix(prefix, "/")
	parts := strings.Split(trimmed, "/")
	if len(parts) < 2 {
		return nil, kv.ErrNoNamespace
	}
	ns := parts[0]

	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(d.tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"Key": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(ns),
					},
				},
			},
			"Path": {
				ComparisonOperator: aws.String("BEGINS_WITH"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(strings.Join(parts[1:], "/")),
					},
				},
			},
		},
		Limit: aws.Int64(int64(pageSize + 1)),
	}
	if fromKey != "" && fromPath != "" {
		queryInput.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"Key": {
				S: aws.String(fromKey),
			},
			"Path": {
				S: aws.String(fromPath),
			},
		}
	}
	result, err := d.db.Query(queryInput)
	if err != nil {
		return nil, err
	}

	pairs := make([]*kv.Pair, 0, len(result.Items))
	for _, item := range result.Items {
		// Since we fetch pageSize+1, stop if we have enough.
		if len(pairs) >= int(pageSize) {
			break
		}
		pairs = append(pairs, &kv.Pair{
			Key:   *item["Key"].S,
			Value: item["Value"].B,
		})
	}

	var nextToken string
	var lastEvalKey string
	var lastEvalPath string

	// We always fetch pageSize+1, so if we have pageSize+1 results, we need to
	// create a page token. We can't continue from result.LastEvaluatedKey since
	// that will skip over the last result. So we should continue from the last
	// visible result.
	if len(result.Items) > int(pageSize) {
		lastEvalKey = *result.Items[len(pairs)-1]["Key"].S
		lastEvalPath = *result.Items[len(pairs)-1]["Path"].S
	} else if result.LastEvaluatedKey != nil {
		lastEvalKey = *result.LastEvaluatedKey["Key"].S
		lastEvalPath = *result.LastEvaluatedKey["Path"].S
	}

	if lastEvalKey != "" && lastEvalPath != "" {
		// Add page token to the database.
		ptKey, err := generatePageToken()
		if err != nil {
			return nil, err
		}

		ptBytes, err := proto.Marshal(&basepb.DynamoDbPageToken{
			LastKey:  lastEvalKey,
			LastPath: lastEvalPath,
		})
		if err != nil {
			return nil, err
		}

		// Create the page token in the database, and it should expire in 2 days.
		_, err = d.db.PutItem(&dynamodb.PutItemInput{
			TableName: aws.String(d.tableName),
			Item: map[string]*dynamodb.AttributeValue{
				"Key": {
					S: aws.String("_internals"),
				},
				"Path": {
					S: aws.String(fmt.Sprintf("page_tokens/%s", ptKey)),
				},
				"Value": {
					B: ptBytes,
				},
				"ExpiresAt": {
					N: aws.String(fmt.Sprintf("%d", time.Now().Add(48*time.Hour).Unix())),
				},
			},
		})
		if err != nil {
			return nil, err
		}

		nextToken = ptKey
	}

	return &kv.Query{
		Prefix:    prefix,
		Pairs:     pairs,
		NextToken: nextToken,
	}, nil
}

func generatePageToken() (string, error) {
	possible := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	token := make([]byte, 48)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(token); i++ {
		token[i] = possible[token[i]%byte(len(possible))]
	}

	return "v1." + string(token), nil
}
