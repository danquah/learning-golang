package hello

// Direct dependency.
import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}
