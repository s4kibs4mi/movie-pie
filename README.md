# Movie Pie

#### Requirements
- Docker
- Docker Compose
- MySQL
- Consul

#### Instructions
- Clone this project

- Run Dependencies
```
docker-compose -f docker-compose-dependencies.yml up
```

- Run App
```
docker-compose up
```

Note : Before running app, change environment variables
related to consul & configs in consul. Example config given in config.example.yml

- Browse at `http://localhost:5005`

### Endpoints

- ###### `POST /v1/register`
```json
{
  "email": "who@example.com",
  "name": "Sakib Sami",
  "password": "111111"
}
```

- ##### `POST /v1/login`
```json
{
  "email": "who@example.com",
  "password": "111111"
}
```

- ##### `GET /v1/movies/search?title={movie_title}&page={page_no}`

response,
```json
{
    "data": [
        {
            "Poster": "https://m.media-amazon.com/images/M/MV5BMTg0NTM3MTI1MF5BMl5BanBnXkFtZTgwMTAzNTAzNzE@._V1_SX300.jpg",
            "Title": "Hello, My Name Is Doris",
            "Type": "movie",
            "Year": "2015",
            "imdbID": "tt3766394"
        },
        {
            "Poster": "https://m.media-amazon.com/images/M/MV5BODJmZmFiNzQtMDJiYS00ZTgzLTljZGMtNjEzNzM4NmEyYjNiXkEyXkFqcGdeQXVyNjE5MjUyOTM@._V1_SX300.jpg",
            "Title": "Hello, Dolly!",
            "Type": "movie",
            "Year": "1969",
            "imdbID": "tt0064418"
        },
        {
            "Poster": "https://m.media-amazon.com/images/M/MV5BNjYxMjI3MzY3NF5BMl5BanBnXkFtZTgwMTgyNzg3MDE@._V1_SX300.jpg",
            "Title": "Hello Ladies",
            "Type": "series",
            "Year": "2013–2014",
            "imdbID": "tt2378794"
        },
        {
            "Poster": "https://m.media-amazon.com/images/M/MV5BMzkzMDc0Nzg5OF5BMl5BanBnXkFtZTcwMDU0MzAyOA@@._V1_SX300.jpg",
            "Title": "Hello I Must Be Going",
            "Type": "movie",
            "Year": "2012",
            "imdbID": "tt2063666"
        },
        {
            "Poster": "http://ia.media-imdb.com/images/M/MV5BMTQ5MjYxMjkwOV5BMl5BanBnXkFtZTgwODE3MjY0MzE@._V1_SX300.jpg",
            "Title": "Hello Ladies: The Movie",
            "Type": "movie",
            "Year": "2014",
            "imdbID": "tt3762944"
        },
        {
            "Poster": "https://m.media-amazon.com/images/M/MV5BMjk1MDczMGQtY2RkNS00OGVhLWJhNzYtNWMwMzFhNTcyNjczXkEyXkFqcGdeQXVyODE5NzE3OTE@._V1_SX300.jpg",
            "Title": "Hello Brother",
            "Type": "movie",
            "Year": "1999",
            "imdbID": "tt0233856"
        },
        {
            "Poster": "https://images-na.ssl-images-amazon.com/images/M/MV5BNDAyOTY2MzE4N15BMl5BanBnXkFtZTgwMjY0OTI5NDE@._V1_SX300.jpg",
            "Title": "Hello Ghost",
            "Type": "movie",
            "Year": "2010",
            "imdbID": "tt1848926"
        },
        {
            "Poster": "https://m.media-amazon.com/images/M/MV5BMTI1MDU1OTM2N15BMl5BanBnXkFtZTcwNjYzMjUyMQ@@._V1_SX300.jpg",
            "Title": "Hello Again",
            "Type": "movie",
            "Year": "1987",
            "imdbID": "tt0093175"
        },
        {
            "Poster": "https://m.media-amazon.com/images/M/MV5BYTE0Yzk1ZDEtN2E2Mi00Y2I0LTkzN2QtZjU4ODlhYTgxODgyXkEyXkFqcGdeQXVyNzI1NzMxNzM@._V1_SX300.jpg",
            "Title": "Hello Stranger",
            "Type": "movie",
            "Year": "2010",
            "imdbID": "tt1725995"
        },
        {
            "Poster": "https://m.media-amazon.com/images/M/MV5BZmQ3YmM0NGMtYmRmNi00ZWY4LTk5MGYtYzUyODA4ODBlODE3XkEyXkFqcGdeQXVyMjQzNzk2ODk@._V1_SX300.jpg",
            "Title": "Oh, Hello on Broadway",
            "Type": "movie",
            "Year": "2017",
            "imdbID": "tt6987652"
        }
    ]
}
```

### Note
ssh into docker container. then run `movie-pie migration auto`, this will populate the database tables.
