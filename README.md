# Simple incrementer application

### Init
```go
	i := incrementer.New()
```

### Increment
```go
	i := incrementer.New()
	i.IncrementNumber()
	fmt.Println(i.GetNumber())
```

### SetMaxValue
```go
	i := incrementer.New()
	i.SetMaxValue(3)
	i.IncrementNumber()
	fmt.Println(i.GetNumber())
```