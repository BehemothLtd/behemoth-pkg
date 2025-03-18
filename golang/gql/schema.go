package gql

import (
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

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

func CreateRootSchema(schemaDir, baseSchemaPath string, sharedSchemaPath string, outputSchemaPath string, outputFile string) error {

	schema, err := FetchSchema(filepath.Join(baseSchemaPath, schemaDir), sharedSchemaPath)
	if err != nil || schema == "" {
		return fmt.Errorf("failed to get schema for %s: %w", schemaDir, err)
	}

	outputPath := filepath.Join(outputSchemaPath, outputFile)
	if err := os.WriteFile(outputPath, []byte(schema), 0644); err != nil {
		return fmt.Errorf("failed to write schema to %s: %w", outputPath, err)
	}

	return nil
}

func FetchSchema(schemaPath string, sharedSchemaPath string) (string, error) {
	var schemaContent []byte

	if err := ReadSchemaFiles(schemaPath, &schemaContent); err != nil {
		return "", err
	}

	if err := ReadSchemaFiles(sharedSchemaPath, &schemaContent); err != nil {
		return "", err
	}

	return string(schemaContent), nil
}

func ReadSchemaFiles(directory string, schemaContent *[]byte) error {
	return filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %s: %w", path, err)
		}
		if d.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			log.Printf("⚠️ Failed to read file %s: %v", path, err)
			return nil
		}

		*schemaContent = append(*schemaContent, content...)
		*schemaContent = append(*schemaContent, '\n')

		return nil
	})
}
