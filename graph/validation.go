package graph

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go_gql/validator"
)

func validation(ctx context.Context, v validator.Validation) bool {
	isValid, errs := v.Validate()

	if !isValid {
		for k, e := range errs {
			graphql.AddError(ctx, &gqlerror.Error{
				Message: e,
				Extensions: map[string]interface{}{
					"field": k,
				},
			})
		}
	}

	return isValid
}
