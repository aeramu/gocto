# gocto
Golang Clean Architecture Code Generator with unit testing template with testify mock to encourage Test Driven Development

## Getting Started
```bash
go get -u github.com/aeramu/gocto
go get -u github.com/vektra/mockery
```

## Usage
### Init project
```bash
gocto init github.com/foo/bar
```
### Edit ```schema.json```
```json
{
  "module_path": "github.com/foo/bar",
  "entity": [
    "Book",
    "Foo"
  ],
  "service": [
    "CreateBook",
    "GetFoo"
	],
  "adapter": [
    "BookRepository",
    "FooClient"
  ]
}
```
- ```entity```: list of entity (business domain) in this service
- ```service```: list of service method (usecase domain) in this service
- ```adapter```: list of interface that needed by service (act as adapter that to be implemented)
### Generate
```bash
make generate
```

## Development (Recommended Practice)
### Define the ```entity```
It's a good practice to define the entity first. What properties and behavior that entity has? Remember, this entity doesn't represent database model. It's business domain.
```go
type Book struct {
  ID string
  Title string
  Author string
  Rating float32
  Price float32
}
```
### Define service method api at ```service/api```
Each service method has their own api. Add properties that needed by that usecase and what it's response.
```go
type CreateBookReq struct {
  Title string
  Author string
  Price float32
}
type CreateBookRes struct {
  Messaage string
}
```
### Start from ```service_test.go```
It's a good practice to begin with test first (Test Driven Development). This already has a template table driven test, so you just need to fill the table with every case.
```go
{
  name:    "should success",
  prepare: func(){
    mockFooRepository.On("test", mock.Anything, mock.MatchedBy(func(b entity.Book) {
      assert.Equal(t, expTitle, b.Title)
      assert.Equal(t, expAuthor, b.Author)
      assert.Equal(t, expPrice, b.Price)
      assert.Equal(t, expRating, b.Rating)
    }))
  },
  args:   args{
    ctx: ctx,
    req: api.CreateBookReq{
      Title: title,
      Author: author,
      Price: price,
    },
  },
  want:    &api.CreateBookRes{
    Message: "Success",
  },
  wantErr: false,
},
```
I personally prefer using testify mocking because I don't need to define the method of adapter interface first, so I can make test first. 
And testify mock has ```mock.MatchedBy()``` that very flexible to detailed whitebox testing.
### Green the test
#### Add needed adapter interface method
```go
type BookRepository interface {
  CreateBook(ctx context.Context, book entity.Book) error
}
```
#### Don't forget to re-generate mock
```bash
make mock
```
#### Do test
```bash
make test
```
Green the test and Do it over again
### Make implementation package
And after service layer done, you can get to work with the implementation layer. Implementation just need to implement the interface that needed by service layer.

## Benefit
The benefit of this approach, if you want to change the infrastructure (ex: changing db from sql to nosql), you don't need to touch the service layer at all. 
You just need to make new implementation that implement the adapter.
And for the handler, if you want to change it (ex: REST to gRPC), you don't need to touch the service layer too, because you just need to adjust the handler
to api package in the service layer.

## Future
I want to make this not just a "first time build service" generator. I want this can be used in the middle of the development to (ex: ```gocto add deleteBook```)<br/>
But, I need to learn more about reading the existing go files, implemented test, etc. It's because generator has a tendency to overwrite the existing file. Should be careful.


