package middleware

import (
	"context"
	"net/http"
	"strconv"
)

type ContextKey string

const LimitKey ContextKey = "limit"
const PageKey ContextKey = "page"

func Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strPage := r.URL.Query().Get("page")
		strLimit := r.URL.Query().Get("limit")
		if strPage == "" {
			strPage = "1"
		}
		if strLimit == "" {
			strLimit = "-1"
		}

		page, pageErr := strconv.Atoi(strPage)
		if pageErr != nil {
			panic("Page not valid")
		}

		limit, limitErr := strconv.Atoi(strLimit)
		if limitErr != nil {
			panic("Limit not valid")
		}
		ctx := context.WithValue(r.Context(), PageKey, page)
		ctx = context.WithValue(ctx, LimitKey, limit)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}