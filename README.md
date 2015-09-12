# Gotains

A module to verify the content of Slices and Arrays

## Use cases

* Verify if one or more items there are in your Slice or Array
* You can verify all types that exists in go

## Usage

* Download `gotains` to your project:
```
go get -u github.com/viniciusfeitosa/gotains
```
* After that use in your code
```
exist, err := gotains.Contains([]string{"test1", "test2", "test3"}, "test1", "test2")
```
* The return is true if all variadic arguments exists in your slice or array
