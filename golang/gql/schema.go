package gql

import (
	_ "embed"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type IRepositories interface {
}

type Resolver interface {
}

func GqlHandler(repos IRepositories, schema string, resolver *Resolver) gin.HandlerFunc {
	opts := []graphql.SchemaOpt{graphql.UseStringDescriptions(), graphql.UseFieldResolvers()}

	return func(c *gin.Context) {
		gqlSchema := graphql.MustParseSchema(schema, resolver, opts...)
		GinSchemaHandler(gqlSchema)(c)
	}
}

func GinSchemaHandler(gqlSchema *graphql.Schema) gin.HandlerFunc {
	handler := &relay.Handler{Schema: gqlSchema}

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
