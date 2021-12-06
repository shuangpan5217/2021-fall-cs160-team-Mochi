# 2021-fall-cs160-team-Mochi

> **MochiNote** is a web-based service that allows users to create, share, and collaboratively edit notes using different learning modalities. - _2021 Fall CS 160 Team Mochi_

# Team Members

1. April Chao
2. Feng Zhang
3. Fudong Huang
4. Xiaoshu Xiao
5. Shuang Pan

# Set Up Dev Env

### 1. Fork this repo

### 2. Create a `go/src` dic from your home directory

```
cd ~
mkdir go
cd go
mkdir src
```

### 3. Get forked repo clone url by hitting "Code" and copying the https or ssh url

-   example HTTPS url: `https://github.com/aprilemeraldchao/2021-fall-cs160-team-Mochi.git`

### 4. Clone the forked repo to the `go/src` directory

```
cd ~/go/src
git clone <forked repo clone url>
```

### 5. Set the upstream repo

```
cd ~/go/src/2021-fall-cs160-team-Mochi
git remote add upstream git@github.com:shuangpan5217/2021-fall-cs160-team-Mochi.git
```

### 6. Set up the db server

-   Download PostgreSQL for the db server: `https://www.postgresql.org/download/`
-   Open `PostgreSQL`
-   Create a new server with port `5432`

### 7. Set up the backend server

-   Download Go for the backend server: `https://golang.org/dl/`
-   Go to the api configuration file
    ```
    /2021-fall-cs160-team-Mochi/backend/source/generated/restapi/configure_coreapi.go
    ```
-   Change the `dbname` and `user` on the following line to your username (for your computer)
    ```
    db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=shuangpan user=shuangpan sslmode=disable")
    ```
-   Rebuild the server with this configuration
    ```
    cd 2021-fall-cs160-team-Mochi/backend/source/generated/cmd/coreapi-server
    go build
    ```

### 8. Set up the frontend server

-   Download node and npm for the frontend server: `https://docs.npmjs.com/downloading-and-installing-node-js-and-npm`
-   Install all necessary packages
    ```
    npm install
    ```

# Run the website locally

### 1. Start db server

-   Open `PostgreSQL`
-   Start the server you created with port `5432`

### 2. Start backend server

-   Ensure you have the correct configuration in `configure_coreapi.go` and have run `go build`
-   Start the api server
    ```
    cd /backend/source/generated/cmd/coreapi-server
    ./coreapi-server
    ```
-   _Note: the backend will listen to the `localhost:3001`_

### 3. Start frontend server

-   Ensure you have run `npm install`
-   _Note: If there are any issues with any node modules, try running `npm cache clean --force` and deleting the `node_modules` folder before retrying `npm install`_
-   Start the web server
    ```
    cd /frontend
    npm start
    ```
-   \*Note: You will likely need to enter `y` to start the server on `localhost:3001` instead of the default since the backend server is already running on `localhost:3001 `
-   If the website hasn't opened, go to your browser and goto `localhost:3001`

# Making Changes

All changes made to this directory must go through a PR/review process.

### 1. Create a new branch

```
git switch -c <branch name>
```

-   _Note: each branch should be associated with a single PR, so ensure that the changes planned are small enough to review in a single session_

### 2. Sync with this upstream repo

```
git stash
git pull upstream main
git stash pop
```

-   _Note: whenever creating a new branch or pushing changes to the upstream repo, you should sync with any changes made to the main branch of the upstream repo_

### 3. Make your changes

-   _Note: try to logically space out your commits to make the PR easier to review_

### 4. Push your changes to the forked repo

-   If you are using https and this is your first push, you need to set the upstream branch
    ```
    git push --set-upstream origin <branch name>
    ```
-   If this is not your first push and/or you have an existing PR you'd like to add to, just run
    ```
    git push
    ```
-   _Note: don't push the `configure_coreapi_go` file if you just configured the `db_name` and `user` attributes_

### 5. Create a PR

-   Create a pull request in your forked repo under the `Pull requests` tab
    -   Left branch: `main` branch of the upstream repo
    -   Right branch: your branch from your forked repo
    -   Comment: summarize the changes/features implemented in this pr
-   _Note: always try to assign a reviewer and attach any related `ZenHub` tasks_

### 6. Review a PR

-   Review the file changes for coding best practices:
    -   naming conventions
    -   typos
    -   formatting
    -   overcomplex logic
    -   etc..
-   Checkout a pr locally to test functionality
    ```
    git fetch upstream pull/<pr #>/head:<local branch name>
    git checkout <local branch name>
    ```

### 7. Merge the PR

-   Once approved, merge the PR into the upstream repo
-   _Note: ensure the it has been sync'd with the latest changes first_

# Set Up Test Env

### Front-End test automation, with `Cypress`

1. Start db server
2. Clear any data in the db server (run `TRUNCATE TABLE users CASCADE;` in Postgres)
3. Start backend server
4. Start frontend server
5. Open Cypress from `/frontend`
    ```
    npx cypress open
    ```
6. Click any test file in `Cypress` to test.
   _Note: you may need to clear the db for some tests_

### Back-end test automation, go `testing` and `net/http/httptest` packages

1. Start db server
2. From `2021-fall-cs160-team-Mochi/backend`, run `go test ./source/apis/notes/...`
3. From `2021-fall-cs160-team-Mochi/backend`, run `go test ./source/apis/usermgmt/...`
4. Check results
    ```
     Example of successful test results:
     1. ok  	2021-fall-cs160-team-Mochi/backend/source/apis/notes	1.221s
     2. ok  	2021-fall-cs160-team-Mochi/backend/source/apis/usermgmt	1.108s
    ```
