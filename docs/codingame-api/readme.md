# CodinGame API

This is a non-official CodinGame REST API description.

All request use `https://www.codingame.com/services/`

You can see all error responses [here](./error_response.md)

<br><br>

## Table of contents

- [CodinGame .](#codingame-api)
  - [Table of contents](#table-of-contents)
  - [Open Endpoints](#open-endpoints)
  - [Closed Endpoints](#closed-endpoints)
    - [Activities routes `/LastActivities`](#activities-routes-lastactivities)
    - [Puzzle routes `/Puzzle`](#puzzle-routes-puzzle)

<br><br>

## Open Endpoints

Open endpoint does not require authentication

* [Login](./login.md) : `POST /codingamerRemoteServiceloginSiteV2`

<br>

## Closed Endpoints

Closed endpoints need authentication to access

### Activities routes `/LastActivities`

* [Last Activities](./closed/lastActivities.md) : `POST /getLastActivities`

### Puzzle routes `/Puzzle`

* [Programming languages](./closed/lastActivities.md) : `POST /findAvailableProgrammingLanguages`
* [Progress](./closed/progress.md) : `POST /findProgressByPrettyId`
* [Solution](./closed/solution.md) : `POST /Solution/findMySolution`
* [Solutions](./closed/solutions.md) : `POST /Solution/findMySolutions`
