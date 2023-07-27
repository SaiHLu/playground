# Testing

1. append "\_test" to file which you want to run tests.

```
E.g => main_test.go
```

2. init project module with "go mod init module_name"

```
- Run all tests in project/module
	1. without verbose
	=> go test ./...

	2. with verbose
	=> go test -v ./...

- Run a specific test function
	1. without verbose
	=> go test ./ -run TestFunctionName

	2. With verbose
	=> go test ./ -v -run TestFunctionName

- Run without cache
    1. without verbose
	=> go test ./ -run TestFunctionName -count=1

	2. With verbose
	=> go test ./ -v -run TestFunctionName -count=1

```

# Channels

1. There are two types of channel

   - Buffered (fix size)

   ```
   # buffer size is one
   # it will be fulled when there is one value is given to channel # then it will be blocked
   bufferedCh := make(chan string, 1)
   ```

   - Unbuffered

   ```
   # will block in initial state
   unbufferedCh := make(chan string)
   ```

- Channels in go will be blocked if it is fulled

Ref: https://github.dev/percoguru/tutorial-notes-api-go-fiber
