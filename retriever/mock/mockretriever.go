package mock

type Aaaaa struct {
	Contents string
}

func (r Aaaaa) Post(url string, form map[string]string) string {
	panic("implement me")
}

func (r Aaaaa) Get(url string) string {
	return r.Contents
}
