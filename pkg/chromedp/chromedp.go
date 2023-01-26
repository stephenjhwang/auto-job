// Command text is a chromedp example demonstrating how to extract text from a
// specific element.
package chromedp

import (
	"context"

	"github.com/chromedp/chromedp"
)

func Initialize() (context.Context, func()) {
	ctx, cancel := chromedp.NewContext(context.Background())
	return ctx, cancel
}
