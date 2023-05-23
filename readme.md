# Golang Unit Test

# Aturan

Tambahkan \_test pada akhir nama file

```go
hello_world_test.go
```

Awali Test pada nama function-nya

Tambahkan parameter t \*Testing.T dan tidak mengembalikan return value

```go
func TestHelloWorld(t *Testing.T) {
	//
}
```

---

Jalankan perintah go mod init `nama-project`

Untuk menjalankan unit test, project harus memiliki struktur seperti berikut

```
- go-unit
    |- helper
    |   |- hello_world.go
    |   |- hello_world_test.go
    |
    |- go.mod
```

Jalankan unit test dengan perintah `go test` di dalam folder helper

Jika ingin menjalankan satu function testing saja bisa menggunakan `go test -v -run TestHelloWorld`

# Assertion

```shell
go get github.com/stretchr/testify
```
