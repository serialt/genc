# genc

1) to generate code, it will create the query dir.
```shell
go run cmd/gen/generate.go 
```
2) to run main
```shell
go run cmd/sugar/main.go 
&{1 student1 1}
&{1 teacher1 []}
&{1 teacher1 [{1 student1 1} {2 student2_new 1}]}
student
teacher
```