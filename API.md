**Article App API**
----
**Get Articles**

Gets a list of articles limited with application parameter "maxrecordperpage" defined in config.yml file.
This api can be used to get an introductory list of articles in inventory. 

* **URL**

  /api/v1/articles

* **Method:** `GET` 

* **Success Response:**
    * **Code:** 200 OK<br />
    * **Content:** 
      * totalCount [int]: Total counts of articles in inventory
      * totalPage [int]: Total count of pages to get the whole inventory. Approximately : (TotalCount/maxrecordperpage)
      * currentPage [int]: Current page of articles in the inventory, default 1
      * perPage [int]: Maximum number of articles listed in a query response. (application.maxrecordperpage parameter)
      * keyword [string]: Search text used in filtering. Default empty 
      * data [list of articles] : Article list in the current page
        * article._id [string] : objectId of article in db inventory
        * article.title [string] : Article Title
        * article.intro [string] : Article Introduction Part
        * article.body [string] : Article Body Text
        * article.postDate [string] : Date string in yyyy-mm-dd hh:mi:ss (2022-01-16 21:43:31) format

    > {
       "totalCount": 3,
      "totalPage": 1,
      "currentPage": 1,
      "perPage": 12,
      "keyword": "",
      "data": [{
      "_id": "61e402e964d20180845ea17e",
      "title": "Best Practices in Go Programming",
      "intro": "10 Best Practices tips for go programming language",
      "body": "1. Naming Conventions 2. Modularity",
      "postDate": "2022-01-16 21:43:31"
      },
      {
      "_id": "61e402e964d20180845ea17f",
      "title": "Thread Management in Go",
      "intro": "Learn how to use goroutines",
      "body": "Creating an OS Thread or switching from one to another can be costly for your programs in terms of memory and performance. Go aims to get advantages as much as possible from the cores. It has been designed with concurrency in mind from the beginning.",
      "postDate": "2022-01-16 21:43:31"
      },
      {
      "_id": "61e402e964d20180845ea180",
      "title": "Restful API with GO",
      "intro": "Creating RESTFUL API using Golang and MongoDB",
      "body": "Hello, Gophers welcome back, In our previous tutorial we integrated Postgres to our Go app. In this tutorial, we are going to integrate MongoDB using the Go Mongo DB driver.",
      "postDate": "2022-01-16 21:43:31"
      }]
      }


* **Error Response:**
  * **Code:** 500 Internal Server Error <br />
  * **Content:** 
    * `{ "error" : "Could not get articles for page. Error: Count query error."}`
    * `{ "error" : "Could not get articles for page. Error: Query error."}`
    * `{ "error" : "Could not get articles for page. Error: Query result mapping error."}`
    * `{ "error" : "Could not get articles for page. Error: Unexpected error."}`


* **Sample Call:**

CURL

    >curl -v http://localhost:8080/api/v1/articles

----
**Get Articles For Page**

Gets a list of articles from inventory with the page number given in the request's path variable.  
This api can be used to get so many articles in inventory with repetitive requests with incrementing page id.  
Article count in response is limited with application parameter "maxrecordperpage" defined in config.yml file.

* **URL**

  /api/v1/articles/page/{page}

* **Method:** `GET`

* **Success Response:**
    * **Code:** 200 OK<br />
    * **Content:**
      * totalCount [int]: Total counts of articles in inventory
      * totalPage [int]: Total count of pages to get the whole inventory. Approximately : (TotalCount/maxrecordperpage)
      * currentPage [int]: Current page of articles in the inventory, default 1
      * perPage [int]: Maximum number of articles listed in a query response. (application.maxrecordperpage parameter)
      * keyword [string]: Search text used in filtering. Default empty
      * data [list of articles] : Article list in the current page
          * article._id [string] : objectId of article in db inventory
          * article.title [string] : Article Title
          * article.intro [string] : Article Introduction Part
          * article.body [string] : Article Body Text
          * article.postDate [string] : Date string in yyyy-mm-dd hh:mi:ss (2022-01-16 21:43:31) format

  > {
  "totalCount": 30,
  "totalPage": 3,
  "currentPage": 3,
  "perPage": 12,
  "keyword": "",
  "data": [
  {
  "_id": "61e08dfd656ff63a47b0b788",
  "title": "16.Restful API with GO",
  "intro": "Creating RESTFUL API using Golang and MongoDB",
  "body": "Hello, Gophers welcome back, In our previous tutorial we integrated Postgres to our Go app. In this tutorial, we are going to integrate MongoDB using the Go Mongo DB driver.",
  "postDate": "2022-01-16 21:43:31"
  },
  {
  "_id": "61e33d5f0899430cbd9dc7f6",
  "title": "16.Restful API with GO",
  "intro": "Creating RESTFUL API using Golang and MongoDB",
  "body": "Hello, Gophers welcome back, In our previous tutorial we integrated Postgres to our Go app. In this tutorial, we are going to integrate MongoDB using the Go Mongo DB driver.",
  "postDate": "2022-01-16 21:43:31"
  },
  {
  "_id": "61e35b0119961a798c4c3dad",
  "title": "16.Restful API with GO",
  "intro": "Creating RESTFUL API using Golang and MongoDB",
  "body": "Hello, Gophers welcome back, In our previous tutorial we integrated Postgres to our Go app. In this tutorial, we are going to integrate MongoDB using the Go Mongo DB driver.",
  "postDate": "2022-01-16 21:43:31"
  }
  ]
  }


* **Error Response:**
    * **Code:** 500 Internal Server Error <br />
    * **Content:**
        * `{ "error" : "Could not get articles for page. Error: Count query error."}`
        * `{ "error" : "Could not get articles for page. Error: Query error."}`
        * `{ "error" : "Could not get articles for page. Error: Query result mapping error."}`
        * `{ "error" : "Could not get articles for page. Error: Unexpected error."}`


* **Sample Call:**

CURL

    >curl -v http://localhost:8080/api/v1/articles/page/1

 ----
**Get Article By Id**

Gets an article with a given id if exists in the articles inventory. 
This api can be used to get a specific article which is already known.  

* **URL**

  /api/v1/articles/{id}

* **Method:** `GET`

* **Success Response:**
    * **Code:** 200 OK<br />
    * **Content:**
        * article._id [string] : objectId of article in db inventory
        * article.title [string] : Article Title
        * article.intro [string] : Article Introduction Part
        * article.body [string] : Article Body Text
        * article.postDate [string] : Date string in yyyy-mm-dd hh:mi:ss (2022-01-16 21:43:31) format

  > {
  "_id": "61e33d5f0899430cbd9dc7f6",
  "title": "16.Restful API with GO",
  "intro": "Creating RESTFUL API using Golang and MongoDB",
  "body": "Hello, Gophers welcome back, In our previous tutorial we integrated Postgres to our Go app. In this tutorial, we are going to integrate MongoDB using the Go Mongo DB driver.",
  "postDate": "2022-01-16 21:43:31"
  }


* **Error Response:**
    * **Code:** 406 Not Acceptable <br />
    * **Content:**
        * `{ "error": "the provided hex string is not a valid ObjectID"}`

    * **Code:** 404 Not Found <br />
    * **Content:**
      * `{ "error" : "No article found with id:{id}"}`


* **Sample Call:**

CURL

    >curl -v http://localhost:8080/api/v1/articles/61e33d5f0899430cbd9dc7f6

 ----
**Publish Article**

Publish an article with a given article information in the request body.
This api can be used to insert an article to the article inventory.

* **URL**

  /api/v1/articles

* **Method:** `POST`

* **Content Type:** `Content-Type: application/json`

* **Request Body:**
  * article.title [string] : Article Title
  * article.intro [string] : Article Introduction Part
  * article.body [string] : Article Body Text
    > `{   
    "title": "16.Restful API with GO",
    "intro": "Creating RESTFUL API using Golang and MongoDB",
    "body": "Hello, Gophers welcome back, In our previous tutorial we integrated Postgres to our Go app. In this tutorial, we are going to integrate MongoDB using the Go Mongo DB driver."
    }`

Note : postdate in article is assigned during insertion and not given by client 

* **Success Response:**
    * **Code:** 200 <br />
    * **Content:**
        * id: Object id of article created in the inventory
  
  > `{
  "id": "61e416b24687182076fd6773"
  }`


* **Error Response:**
    * **Code:** 400 Bad Request <br />
    * **Content:**
        * `{ "error": "Article input is not valid. Error:id can not be empty}`
        * `{ "error": "Article input is not valid. Error:title can not be empty}`
        * `{ "error": "Article input is not valid. Error:intro can not be empty}`
        * `{ "error": "Article input is not valid. Error:body can not be empty}`

* **Error Response:**
    * **Code:** 500 Internal Server Error <br />
    * **Content:**
        * `{ "error": "Could not publish the article. Error:{errorMessage}}"`


* **Sample Call:**

CURL

    >curl -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/articles -d "{\"title\": \"16.Restful API with GO\", \"intro\": \"Creating RESTFUL API using Golang and MongoDB\", \"body\": \"Hello, Gophers welcome back, In our previous tutorial we integrated Postgres to our Go app. In this tutorial, we are going to integrate MongoDB using the Go Mongo DB driver.\" }"

 ----
**Search Article By Text**

Searches an article with a given text in the query parameter.
This api can be used to query articles according to their title, intro and body texts.
Mongodb text index is created and equal priority weight is given to each field in search operation.  
Search text needs to be URL encoded string.

* **URL**

  /api/v1/articles/search?searchText={searchText}&page={page}

* **Method:** `GET`

* **Query Parameters:**
    * searchText: Text to search 
    * page: Filter articles by page id in case of too many returns

* **Success Response:**
    * **Code:** 200 <br />
    * **Content:**
      * totalCount [int]: Total counts of articles in inventory
      * totalPage [int]: Total count of pages to get the whole inventory. Approximately : (TotalCount/maxrecordperpage)
      * currentPage [int]: Current page of articles in the inventory, default 1
      * perPage [int]: Maximum number of articles listed in a query response. (application.maxrecordperpage parameter)
      * keyword [string]: Search text used in filtering. Default empty
      * data [list of articles] : Article list in the current page
          * article._id [string] : objectId of article in db inventory
          * article.title [string] : Article Title
          * article.intro [string] : Article Introduction Part
          * article.body [string] : Article Body Text
          * article.postDate [string] : Date string in yyyy-mm-dd hh:mi:ss (2022-01-16 21:43:31) format

  > `{
  "totalCount": 1,
  "totalPage": 1,
  "currentPage": 1,
  "perPage": 12,
  "keyword": "URL encode/decode+parseQuery",
  "data": [
  {
  "_id": "61db01c1a44382e1f4161b55",
  "title": "1.Best Practices in Go Programming",
  "intro": "10 Best Practices tips for go programming language",
  "body": "URL encode/decode+parseQuery 1. Naming Conventions 2. Modularity",
  "postDate": "2022-01-16 21:43:31"
  }
  ]
  }`

* **Sample Call:**

CURL

    >curl -v http://localhost:8080/api/v1/articles/search?searchText=URL%20encode%2Fdecode%2BparseQuery&page=1

