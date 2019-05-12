
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "io/ioutil"
    "fmt"
)

// Yelps is a string of yelp structs
type Yelps struct {
    Businesses []Yelp `json:"businesses"`
}

// Reviews format
type Reviews struct {
    Reviews []Review `json:"reviews"`
}

// Yelp fusion api format
type Yelp struct {
    ID           string       `json:"id"`
    Alias        string       `json:"alias"`
    Name         string       `json:"name"`
    ImageURL     string       `json:"image_url"`
    IsClosed     bool         `json:"is_closed"`
    URL          string       `json:"url"`
    ReviewCount  int          `json:"review_count"`
    Category     []Categories `json:"categories"`
    Rating       float32      `json:"rating"`
    Coordinates  Coords       `json:"coordinates"`
    Transactions []string     `json:"transactions"`
    Price        string       `json:"price"`
    Location     Location     `json:"location"`
    Phone        string       `json:"phone"`
    DisplayPhone string       `json:"display_phone"`
    Distance     float32      `json:"distance"`
}

// Review for yelp restaurant
type Review struct {
    ID          string   `json:"id"`
    URL         string   `json:"url"`
    Text        string   `json:"text"`
    Rating      float32  `json:"rating"`
    TimeCreated string   `json:"time_created"`
    User        UserInfo `json:"user"`
}

// Categories format
type Categories struct {
    Alias string `json:"alias"`
    Title string `json:"title"`

}

// Location format in yelp fusion api
type Location struct {
    Address1       string   `json:"address1"`
    Address2       string   `json:"address2"`
    Address3       string   `json:"address3"`
    City           string   `json:"city"`
    Zip            string   `json:"zip_code"`
    Country        string   `json:"country"`
    State          string   `json:"state"`
    DisplayAddress []string `json:"display_address"`
}

// Coords format in yelp for longitude latitude
type Coords struct {
    Latitude  float32 `json:"latitude"`
    Longitude float32 `json:"longitude"`
}
// UserInfo info for user reviews
type UserInfo struct {
    ImageURL string `json:"image_url"`
    Name     string `json:"name"`
}


func main(){
//    var token string = "1cK7Lxy3bj1aRdQn7wJseUp5pNZOCs48U9dJfdhmO5f20Ko0iltUO0Asu4YjVXeKfY00gewXW_0IG2pXWPRMi7VMCkQSJDDrKxYJs0WBMxXRvOBmmCt7ZsDjhszQXHYx"

    //var urlpath = "https://api.yelp.com/v3/businesses/search?location=Natick"
    var urlpath = "https://api.yelp.com/v3/businesses/search?location=Somerville&limit=50"


    var bearer = "Bearer " + "1cK7Lxy3bj1aRdQn7wJseUp5pNZOCs48U9dJfdhmO5f20Ko0iltUO0Asu4YjVXeKfY00gewXW_0IG2pXWPRMi7VMCkQSJDDrKxYJs0WBMxXRvOBmmCt7ZsDjhszQXHYx"
    var data Yelps

    req, err := http.NewRequest("GET",urlpath,nil)
    req.Header.Add("Authorization",bearer)
    client := &http.Client{}
    resp,err := client.Do(req)

    if err != nil {
        log.Fatalln(err)

    }



    body, _ := ioutil.ReadAll(resp.Body)


    err = json.Unmarshal(body,&data)

    if err != nil {
        panic(err)
    }

    
    for i := 0; i < len(data.Businesses); i++ {
        fmt.Println("Name:" , data.Businesses[i].Name," = ", data.Businesses[i].Rating)

    }



}

