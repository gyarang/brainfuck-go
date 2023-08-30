# brainfuck-go

## About
brainfuck-go 는 Go 로 작성 된 Brainfuck 언어의 인터프리터입니다. 걍 해봄.
아직 개발 중이며 [Brainfuck test](https://github.com/rdebath/Brainfuck) 통과를 목표로 합니다.

## Usage
### Compile
```
go build ./cli/bf.go
```
### Execute
```
./bf filename
```

## Example
```
./bf ./hello.bf
Hello, Brainfuck-Go
```

## TODO
- [ ] 중첩 반복문 처리