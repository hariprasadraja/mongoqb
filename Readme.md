# mongoqb

mongoqb is a versatile query builder for mongodb from http request built in golang.  

  - do a Request like http://localhost:9000?filterBy=name(and),emailId(or),age&name=hari,rahul,krishnan&emailId=hariprasadcsmails@gmail.com&age=24,

  - get a mgo query as bson.M
  bson.M{
				"name":            bson.M{"$and": []string{"hari", "rahul", "krishnan"}},
				"contact.emailId": bson.M{"$or": []string{"hariprasadcsmails@gmail.com"}},
				"personal.age":    "24",
			},
  - Magic

# Features!
  - Build mongodb query from http request. 
  - Conditional operators in mongodb are allowed. 
  - Custom functions to do operations before making mongodb query.
  - Only exported keys are allowed to transform into mongodb query. 

### Dependencies
mongoqb uses only one source projects to work properly:
* [mgo] - a powerfull golang mongodb driver. 


And of course mongodb itself is open source with a [public repository][dill]
 on GitHub.

### Installation
mongoqb requires [mgo](https://labix.org/mgo) v2 to run.

Install the dependencies by the following

To install dependencies, run the below 

```sh
$ go get gopkg.in/mgo.v2
```

### Development

Want to contribute? Great!
contact me @ hariprasadcsmails@gmail.com

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [mgo]: <http://labix.org/mgo>
