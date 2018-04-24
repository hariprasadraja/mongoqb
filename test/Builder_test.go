// This file provides utility tools to parse the query params from the request.
// It provides a standardised approach to read the data from request query.
// version: 0.1

package test

import (
        "net/http"
        "reflect"
        "testing"

        "gopkg.in/mgo.v2/bson"
        "mongoqb/qb"
)

func TestBuildQuery(t *testing.T) {
        type args struct {
                request     *http.Request
                allowedkeys map[string]qb.Options
        }
        tests := []struct {
                name    string
                args    args
                want    bson.M
                wantErr bool
                url     string
        }{
                {
                        name: "examine single filter params",
                        args: args{
                                allowedkeys: map[string]qb.Options{
                                        "name": {
                                                DBPath: "name",
                                        },
                                },
                        },
                        want: bson.M{
                                "name": "hari",
                        },
                        wantErr: false,
                        url:     "http://localhost:9000?filterBy=name&name=hari",
                },
                {
                        name: "examine multiple filter params",
                        args: args{
                                allowedkeys: map[string]qb.Options{
                                        "name": {
                                                DBPath: "name",
                                        },
                                        "emailId": {
                                                DBPath: "contact.emailId",
                                        },
                                        "age": {
                                                DBPath: "personal.age",
                                        },
                                },
                        },

                        want: bson.M{
                                "name":            "hari",
                                "contact.emailId": "hariprasadcsmails@gmail.com",
                                "personal.age":    "24",
                        },
                        wantErr: false,
                        url:     "http://localhost:9000?filterBy=name,emailId,age&name=hari&emailId=hariprasadcsmails@gmail.com&age=24",
                },
                {
                        name: "examine logical operators in filter params",
                        args: args{
                                allowedkeys: map[string]qb.Options{
                                        "name": {
                                                DBPath: "name",
                                        },
                                        "emailId": {
                                                DBPath: "contact.emailId",
                                        },
                                        "age": {
                                                DBPath: "personal.age",
                                        },
                                },
                        },
                        want: bson.M{
                                "name":            bson.M{"$and": []string{"hari", "rahul", "krishnan"}},
                                "contact.emailId": bson.M{"$or": []string{"hariprasadcsmails@gmail.com"}},
                                "personal.age":    "24",
                        },
                        wantErr: false,
                        url:     "http://localhost:9000?filterBy=name(and),emailId(or),age&name=hari,rahul,krishnan&emailId=hariprasadcsmails@gmail.com&age=24",
                },
                {
                        name: "examine conditional with logical operators in filter params",
                        args: args{
                                allowedkeys: map[string]qb.Options{
                                        "name": {
                                                DBPath: "name",
                                        },
                                        "emailId": {
                                                DBPath: "contact.emailId",
                                        },
                                        "age": {
                                                DBPath: "personal.age",
                                        },
                                },
                        },
                        want: bson.M{
                                "name":            bson.M{"$and": []string{"hari", "rahul", "krishnan"}},
                                "contact.emailId": bson.M{"$or": []string{"hariprasadcsmails@gmail.com"}},
                                "personal.age":    bson.M{"$lte": "24"},
                        },
                        wantErr: false,
                        url:     "http://localhost:9000?filterBy=name(and),emailId(or),age(lte)&name=hari,rahul,krishnan&emailId=hariprasadcsmails@gmail.com&age=24",
                },
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {

                        tt.args.request, _ = http.NewRequest(http.MethodGet, tt.url, nil)
                        got, err := qb.BuildQuery(tt.args.request, tt.args.allowedkeys)
                        if (err != nil) != tt.wantErr {
                                t.Errorf("BuildQuery() error = %v, wantErr %v", err, tt.wantErr)
                                return
                        }
                        if !reflect.DeepEqual(got, tt.want) {
                                t.Errorf("BuildQuery() = %v, want %v", got, tt.want)
                        }
                })
        }
}