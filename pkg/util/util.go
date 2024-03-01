package util

func Map[I, O any](inp []I, fn func(I) O) []O {
	res := make([]O, len(inp))
	for i, item := range inp {
		res[i] = fn(item)
	}
	return res
}

func Pointer[I any](inp I) *I {
	return &inp
}

func DerefOrDefault[I any](inp *I) I {
	var res I
	if inp != nil {
		res = *inp
	}
	return res
}
