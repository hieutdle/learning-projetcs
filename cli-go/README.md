# control
My CLI written in Go

## Count

Count the number of lines in a file

```go
control -count
```

Count the number of words in a file

```go
control -count -words
```
test time zone
Count the number of bytes in a file

```go
control -count -bytes
```

## To-do list

Add tasks the list

```go
control -task -add "task 1"
```

Print the list

```go
control -task -list
```

Delete a task

```go
control -task -delete 1
```

Mark the task as done

```go
control -task -done 1
```
