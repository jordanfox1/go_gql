package graph

import (
	"context"
	"github.com/go-pg/pg/v10"
	"go_gql/graph/model"
	"net/http"
	"time"
)

const userloaderKey = "userloader"

func DataloaderMiddleware(db *pg.DB, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([][]*model.User, []error) {
				var users []*model.User

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()

				if err != nil {
					return nil, []error{err}
				}

				// Create a slice of slices containing a single user for each key.
				result := make([][]*model.User, len(ids))
				for i, id := range ids {
					for _, user := range users {
						if user.ID == id {
							result[i] = []*model.User{user}
							break
						}
					}
				}

				return result, nil
			},
		}

		ctx := context.WithValue(r.Context(), userloaderKey, &userloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userloaderKey).(*UserLoader)
}
