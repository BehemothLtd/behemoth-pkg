package contexts

import (
	"behemoth-pkg/golang/constants"
	"behemoth-pkg/golang/exceptions"
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(constants.CtxGin)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func GinCtxTypeAssertFromMixedCtx(ctx context.Context) (*gin.Context, error) {
	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		log.Error().Msg("Failed to assert context to *gin.Context")
		return nil, errors.New("invalid context")
	}

	return ginCtx, nil
}

// getContextValue safely retrieves a value from Gin's context
func GetContextValue[T any](ctx *gin.Context, key string) (T, error) {
	value, exists := ctx.Get(key)
	if !exists {
		var zeroValue T
		return zeroValue, exceptions.NewUnauthorizedError(nil)
	}

	typedValue, ok := value.(T)
	if !ok {
		var zeroValue T
		return zeroValue, exceptions.NewUnauthorizedError(nil)
	}

	return typedValue, nil
}
