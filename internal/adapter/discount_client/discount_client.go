package discount

type DiscountClient struct {
	URL string
}

func New(url string) *DiscountClient {
	return &DiscountClient{
		URL: url,
	}
}
