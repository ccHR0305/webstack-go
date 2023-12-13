package index

import (
	"github.com/ccHR0305/webstack-go/internal/pkg/core"
)

func (h *handler) About() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("about", nil)
	}

}
