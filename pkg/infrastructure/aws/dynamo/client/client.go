package client

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Client struct {
	DB        *dynamodb.DynamoDB
	tableName string
}

// SetTable は使用するテーブルを設定する
func (c *Client) SetTable(table interface{}) *dynamodb.DynamoDB {
	var prefix string
	switch reflect.TypeOf(table).Kind() {
	case reflect.Ptr:
		prefix = "*entity."
	case reflect.Struct:
		prefix = "entity."

	}
	c.tableName = strings.TrimPrefix(reflect.TypeOf(table).String(), prefix)
	fmt.Println(c.tableName)
	return c.DB
}

// CreateItems は挿入するitemを用意する
func (c *Client) CreateItems(i interface{}) (*dynamodb.AttributeValue, error) {
	return &dynamodb.AttributeValue{}, nil
}
