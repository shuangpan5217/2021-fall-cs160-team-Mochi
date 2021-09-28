# 2021-fall-cs160-team-Mochi
2021 Fall CS160 Team Mochi

##### Team Members
1. April Chao
2. Feng Zhang
3. Fudong Huang
4. Xiaoshu Xiao
5. Shuang Pan

### Run the program 
####1. Start db server
Download PostgreSQL
```
https://www.postgresql.org/download/
```
Start the `psql` server and listen to port `5432`

###2. Start backend Server
Download Go env
```
https://golang.org/dl/
```
Create a `go/src` dic from your home directory, and clone the repo to the `src` repo.
```
git clone https://github.com/shuangpan5217/2021-fall-cs160-team-Mochi.git
```
Open `/2021-fall-cs160-team-Mochi/backend/source/generated/restapi/configure_coreapi.go`
For the following line,
```
db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=shuangpan user=shuangpan sslmode=disable")
```
change `dbname` and `user` as required, add `password=` if needed

```
cd 2021-fall-cs160-team-Mochi/backend/source/generated/cmd/coreapi-server
```
Run
```
1. go build
2. ./coreapi-server
```
The backend will listen to the `localhost:3000`

###Start frontend server



