package tag

type Tags []Tag

type Tag struct {
	ID  int64
	Tag string
}

func New(tag string) *Tag {
	t := Tag{
		Tag: tag,
	}
	return &t
}
