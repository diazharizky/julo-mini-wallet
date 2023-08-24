package app

type err string

func (e err) Error() string {
	return string(e)
}
