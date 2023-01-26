package autofill

import (
	"context"
)

type Autofiller interface {
	Autofill(context.Context)
}
