package mock

type Aaaaa struct {
	Contents string
}

func (r *Aaaaa) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "OK"
}

func (r *Aaaaa) Get(url string) string {
	return r.Contents
}
