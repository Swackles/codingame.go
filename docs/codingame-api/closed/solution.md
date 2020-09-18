# Last Activities

Used to get your solution for the puzzle.

**URL** : `/Puzzle/Solution/findMySolution`

**Method** : `POST`

**Auth required** : YES

**Data constraints** :
When sending data to get puzzles solution, you send json array. First element is userID, second is testSessionQuestionSubmissionId

```json
[ "[userID]", "[testSessionQuestionSubmissionId]" ]
```

## Success Response

**Code** : `200 OK`

**Content example** :

```json
{
    "pseudo": "Swackles",
    "testSessionQuestionSubmissionId": 15691536,
    "programmingLanguageId": "Go",
    "creationTime": 1597864145683,
    "avatar": 46897964584843,
    "codingamerId": 3903812,
    "codingamerHandle": "9683cd12e82596de22c28ce1657cd32f2183093",
    "code": "[Code for the project is here]",
    "votableId": 14743598,
    "commentableId": 14662829,
    "shared": false
}
```
