# go-generic
Fun with go generic

## Generic function

```
func AssertEqual[T comparable](want, actual T)
```

Check more at: `assert.go`

[Playground](https://go2goplay.golang.org/p/PPSgLIYLvLI)

## Generic type

To declare a generic type:
```
type Stack[T any] struct {
	data []T
}
```

Function uses a generic type as parameters
```
func (s *Stack[T]) Push(e T) {
	s.data = append(s.data, e)
}
```

Function return a generic type
```
func (s *Stack[T]) Pop() (T, bool) {
```

Init struct with concrete type
```
stack := new(Stack[int])
```

Check more at: `stack.go`

[Playgound](https://go2goplay.golang.org/p/EiRBzQqsVSN)
