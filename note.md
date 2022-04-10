# Catatan Belajar

## Function

- go menyalin nilai ketika di lewatkan kedalam fungsi/method

## Nil

- Pointer bisa bernilai nil
- Jika fungsi yang mengembalikan pointer, pastikan untuk membuat pengecekan `nil` atau `runtime exception`

## Type

- go bisa membuat `type` baru dari tipe yang sudah ada
  - membrikan makna yang tepat kepata tipe tersebut
  - bisa mengimplementasikan `interface`

## Map

- A map value is a pointer to a runtime.hmap structure.
- So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.
- initialize empty map like this

```go
var dictionary = map[string]string{} //right
var dictionary = make(map[string]string) //wrong
var dictionary map[string]string //wrong
```
