package label

type Labels []Label

type Label struct {
	Key   string
	Value string
}

func New(key, value string) *Label {
	return &Label{
		Key:   key,
		Value: value,
	}
}
