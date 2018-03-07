package mock

type Aaaaa struct {
	Contents string
}

func (r Aaaaa) Get(url string) string {
	return r.Contents
}
