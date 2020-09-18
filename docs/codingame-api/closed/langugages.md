# Last Activities

Used to get available languages for a puzzle and if it has been solved in that languages. You can only get last activites of the user you are logged in as.

**URL** : `/Puzzle/findAvailableProgrammingLanguages`

**Method** : `POST`

**Auth required** : YES

**Data constraints** :
When sending data to get users last activities, you send json array. First element in that json array is puzzleID and second is your userID.

```json
[ "[puzzleID]", "[userID]" ]
```

## Success Response

**Code** : `200 OK`

**Content example** :

```json
[
    {
        "id": "Bash",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "C",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "C#",
        "solved": false,
        "last": false,
        "onboarding": true
    },
    {
        "id": "C++",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Clojure",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "D",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Dart",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "F#",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Go",
        "solved": true,
        "last": true,
        "onboarding": false
    },
    {
        "id": "Groovy",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Haskell",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Java",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Javascript",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Kotlin",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Lua",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "ObjectiveC",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "OCaml",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Pascal",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Perl",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "PHP",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Python3",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Ruby",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Rust",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Scala",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "Swift",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "TypeScript",
        "solved": false,
        "last": false,
        "onboarding": false
    },
    {
        "id": "VB.NET",
        "solved": false,
        "last": false,
        "onboarding": false
    }
]
```
