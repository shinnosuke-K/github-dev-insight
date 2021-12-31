package client

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	BatchSize = 25
)

type Client struct {
	DB *dynamodb.DynamoDB
}

// SetTable は使用するテーブルを設定する
func (c *Client) table(table interface{}) string {
	var prefix string
	switch reflect.TypeOf(table).Kind() {
	case reflect.Ptr:
		prefix = "*entity."
	case reflect.Struct:
		prefix = "entity."
	}
	return strings.TrimPrefix(reflect.TypeOf(table).String(), prefix)
}

// CreateBatchItems は BatchCreate で使う item を用意する
func (c *Client) CreateBatchItems(ents interface{}, table interface{}) ([]*dynamodb.TransactWriteItem, error) {
	v := reflect.Indirect(reflect.ValueOf(ents))
	if v.Type().Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Struct {
		return nil, errors.New("ents must be struct slice")
	}
	tableName := c.table(table)
	items := make([]*dynamodb.TransactWriteItem, v.Len())
	for i := 0; i < v.Len(); i++ {
		item, err := c.createItems(v.Index(i).Interface())
		if err != nil {
			return nil, fmt.Errorf("failed to create item. %w", err)
		}
		items[i] = &dynamodb.TransactWriteItem{
			Put: &dynamodb.Put{
				TableName: aws.String(tableName),
				Item:      item,
			},
		}
	}
	return items, nil
}

// CreateItems は挿入するitemを用意する
func (c *Client) createItems(i interface{}) (map[string]*dynamodb.AttributeValue, error) {
	av, err := dynamodbattribute.MarshalMap(i)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare item. %w", err)
	}
	return av, nil
}

// BatchCreate は item をまとめて挿入する
// NOTE: dynamoDB の仕様上、一度の挿入は25個まで
func (c *Client) BatchCreate(ctx context.Context, items []*dynamodb.TransactWriteItem) error {
	var (
		size = len(items)
	)
	for n := 0; n < size; n += BatchSize {
		start := n
		end := n + BatchSize
		if end > size {
			end = size
		}
		input := &dynamodb.TransactWriteItemsInput{
			TransactItems: items[start:end],
		}

		if _, err := c.DB.TransactWriteItemsWithContext(ctx, input); err != nil {
			return fmt.Errorf("failed to create ")
		}
	}
	return nil
}
