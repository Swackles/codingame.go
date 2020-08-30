# Last Activities

Used to get your solutions for the puzzle. You can only get last activites of the user you are logged in as.

**URL** : `/Puzzle/Solution/findMySolutions`

**Method** : `POST`

**Auth required** : YES

**Data constraints** :
When sending data to get puzzles solution, you send json array. First element is userID, second is puzzle id and Third is langugaeID (received from [languages](langugages.md))

```json
[ "[userID]", "[puzzleID]", "[language]" ]
```

## Success Response

**Code** : `200 OK`

**Content example** :

```json
[
    {
        "pseudo": "Swackles",
        "testSessionQuestionSubmissionId": 15691536,
        "programmingLanguageId": "Go",
        "creationTime": 1597864145683,
        "avatar": 46897964584843,
        "codingamerId": 3903812,
        "commentableId": 14662829,
        "commentCount": 0,
        "upVotes": 0,
        "downVotes": 0,
        "shared": false
    }
]
```
