# Last Activities

Used to get persons last activies. You can only get last activites of the user you are logged in as.

**URL** : `/LastActivities/getLastActivities`

**Method** : `POST`

**Auth required** : YES

**Data constraints** :
When sending data to get users last activities, you send json array. First element in that json array is userID and second element is amount of items you want to get from your activities.

```json
[ "[userID]", "[count]" ]
```

## Success Response

**Code** : `200 OK`

**Content example** : Returns array of fallowing objects

```json
{
  "type": "PUZZLE",
  "puzzle": {
    "id": 41,
    "level": "medium",
    "rank": 0,
    "thumbnailBinaryId": 1511310761406,
    "previewBinaryId": 1511321167592,
    "coverBinaryId": 5691367634830,
    "logoBinaryId": 4768403842892,
    "title": "Shadows of the Knight - Episode 1",
    "description": "<i>NA NA NA NA NA BATMAN</i><br>\nWe love Batman's adventures : The Joker, bombs, hostages, and a hero. But this time, it is you who are the hero. Your job is to find the bombs before they explode! Don't worry, Alfred's got you covered, he's handed you a heat detector set to recognize the thermal signature of the Joker's bombs. Easy? Let's find out.\n<br>\n<br>\n<u>Topic</u> : Convergence.<br>\n<br>\n<i>This medium puzzle is the first of the two exercises proposed for the challenge «&nbsp;Shadows of the Knight&nbsp;». Once you have completed it, why not try to solve the next one «&nbsp;Triangulation&nbsp;», on the same theme but with an increased difficulty.</i>",
    "validatorScore": 0,
    "achievementCount": 2,
    "doneAchievementCount": 0,
    "lastActivity": 1598041949060,
    "forumLink": "shadows-of-the-knight-episode-1-puzzle-discussion/264",
    "chatRoom": "shadows_of_knight_episode_1",
    "solvedCount": 46834,
    "attemptCount": 106682,
    "xpPoints": 100,
    "feedback": {
        "feedbackId": 161,
        "feedbacks": [
            102,
            68,
            359,
            1539,
            4984
        ]
    },
    "topics": [
        {
            "handle": "algorithms",
            "value": "Algorithms",
            "children": [
                {
                    "handle": "binary-search",
                    "value": "Binary search",
                    "children": []
                }
            ]
        },
        {
            "handle": "maths-physics",
            "value": "Maths / Physics",
            "children": [
                {
                    "handle": "intervals",
                    "value": "Intervals",
                    "children": []
                }
            ]
        }
    ],
    "creationTime": 1407700652000,
    "type": "SOLO",
    "prettyId": "shadows-of-the-knight-episode-1",
    "detailsPageUrl": "/training/medium/shadows-of-the-knight-episode-1",
    "replayIds": [ 
    ],
    "testSessionHandle": "31049118afb43f71b27564ef14602ab52a44b33c",
    "communityCreation": false
  }
}
```
