# CodinGame API

This is a non-official CodinGame REST API description.

All request use `https://www.codingame.com/services/`

You can see all error responses [here](api/error_response.md)

<br><br>

## Table of contents

- [CodinGame API](#codingame-api)
  - [Table of contents](#table-of-contents)
  - [Open Endpoints](#open-endpoints)
  - [Closed Endpoints](#closed-endpoints)
    - [Activities routes `/LastActivities`](#activities-routes-lastactivities)
    - [Puzzle routes `/Puzzle`](#puzzle-routes-puzzle)

<br><br>

## Open Endpoints

Open endpoint does not require authentication

* [Login](api/login.md) : `POST /codingamerRemoteServiceloginSiteV2`

<br>

## Closed Endpoints

Closed endpoints need authentication to access

### Activities routes `/LastActivities`

* [Last Activities](api/closed/lastActivities.md) : `POST /getLastActivities`

### Puzzle routes `/Puzzle`

* [Programming languages](api/closed/lastActivities.md) : `POST /findAvailableProgrammingLanguages`
* [Progress](api/closed/progress.md) : `POST /findProgressByPrettyId`
* [Solution](api/closed/solution.md) : `POST /Solution/findMySolution`
* [Solutions](api/closed/solutions.md) : `POST /Solution/findMySolutions`
