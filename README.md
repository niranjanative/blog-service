# Micro-Blog

Micro-Blog is a simple microservice application which is written in GoLang

**Note**: this is not a production ready application and we do not recommend its in production

## Background details
* This application uses MySQL database to store all the relevant data
* It uses **two tables - articles & comments**
* **articles** table is used to store all the posts & related data
* **comments** table is used to store all comments & related data

## API details
* Application exposes 5 APIs -
  * GET  - `/articles` -> lists all the articles posted on Blog till date without the article content
  * GET  - `/articles/:id` -> gets an article with provided ID with content if present
  * POST - `/articles` -> creates a new article with provided metadata & content
    * Required Body while making this API call: 
    ```
    {
        "nickname": "goat",
        "content": "Ronaldo moves from Juventus to ManU",
        "title": "2021 transfer window"
    }
    ```
  * GET - `/comments/:parent-id` -> parent-id is an integer ID which refers to either an article 
    to which comment(s) belong OR a particular comments ID whose
    replies are to be fetched
    * An additional `parent-type` query param is required for this call which 
    can take 2 values (i.e. `article` OR `comment`) depending on whose sub-comments 
    are to be fetched respectively
  * POST - `/comments` -> Used for creating a comment either for an article OR for an
  existing comment 
    * Required Body while making this API call:
    - If replying to a comment use:
    ```
    {
        "parent_id": 10,
        "parent_type": "comment",
        "nickname": "messi",
        "content": "Barcelona to PSG"
    }
    ```
    - If replying to an article/post use:
    ```
    {
        "parent_id": 7,
        "parent_type": "article",
        "nickname": "ronaldo",
        "content": "Juventus to Manchester United"
    }
    ```
      
## Prerequisites:
* GoLang 1.16+ required
* MySQL DB already installed

## Try it yourself
* create a database in mysql-db
* create a user in mysql-db and provide it permissions to the above created DB
* Create a build by running `go build -o blog` from project root
* Export below env variables by running (change below values as per your DB settings):
  ```
  export DATABASE_NAME=mydbname
  export USER_NAME=myusername
  export PASSWORD=mypassword
  export HOST=localhost
  export PORT=3306
  ```
* Start the application by running - `./blog` (give appropriate name of your build here 
in my case it is `blog`)
* Have fun!!! Start accessing above mentioned APIs

## Future scope:
* Dockerize this microservice
* Create helm-chart for easy deployment on Kubernetes along with other MySQL container which will be used as database
* Add more features
* Create Contributing guidelines
