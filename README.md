# phonebook
For practicing golang

Instruction

1. Clone project.
git clone https://github.com/pbunluesin/phonebook_api.git

2. Install lib and dependencies. (using ECHO API web framework)

go get -u github.com/labstack/echo/
go get -u github.com/globalsign/mgo
go get -u github.com/labstack/echo/middleware

3. Run MongoDB
docker run -p 27017:27017 -d  --name test mongo:latest

4. Run API
go run server.go

5. Test call api

  5.1 List all guests in phonebook 
    http://localhost:1324/api/list

  5.2 Add guest to phonebook
    
    localhost:1324/api/insert
      {

          "firstname": “Phatthara”,
          "lastname": “Bunluesin”,
          "telephone": "888-888-8888",
          "address": "Bangkok"
      }


  5.3 Search guest by specific name

    http://localhost:1324/api/search/Phatthara

  5.4 Update guest by specific name

  PUT /api/update/Phatthara HTTP/1.1
  Host: localhost:1324
  Content-Type: application/json
  cache-control: no-cache
  Postman-Token: 604d365e-f954-48de-abff-019f5b681a6b
      {

           "firstname": "Phatthara",
          "lastname": "Bunluesin",
          "telephone": "099-999-9999",
          "address": "Bangkok"
      }------WebKitFormBoundary7MA4YWxkTrZu0gW--

  5.5Delete guest by specific name
  
    localhost:1324/api/delete/Phatthara

