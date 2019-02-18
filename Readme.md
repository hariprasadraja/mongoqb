# mongoqb 
## (Stopped Developement due to lack of support & personal time) 

mongoqb is a versatile query builder for mongodb from http request built in golang.

## How It works

| Request Query            | Monogo DB Query        | Operators |
| ----------------------- | ---------------------- | --------- |
| filterBy=name&name=hari | bson.M{"name": "hari"} | None      |
| filterBy=name(and),emailId(or),age&name=hari,rahul,krishnan|bson.M{"name":bson.M{"$and": []string{"hari", "rahul", "krishnan"}}}| conditional operators|

## Features

- Build mongodb query from http request.
- Conditional operators in mongodb are allowed.
- Custom functions to do operations before making mongodb query.
- Only exported keys are allowed to transform into mongodb query.

### Dependencies

mongoqb uses only one source projects to work properly:

- [mgo] - a powerfull golang mongodb driver.

### Installation

mongoqb requires [mgo](https://labix.org/mgo) v2 to run.

To install dependencies, run the below

```sh
go get gopkg.in/mgo.v2
```

Install mongoqb

```sh
go get github.com/hariprasadraja/mongoqb

```

### TODO's

1. Need to add more examples.
2. extend functionality to all operators
3. paras query from request body. if the request is made via post call. 
4. parse query from the form request. 
### Development

Want to contribute? Great!
contact me @ hariprasadcsmails@gmail.com

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [mgo]: <http://labix.org/mgo>
