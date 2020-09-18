# Last Activities

Used to get persons progress in a puzzle. You can only get last activites of the user you are logged in as.

**URL** : `/Puzzle/findProgressByPrettyId`

**Method** : `POST`

**Auth required** : YES

**Data constraints** :
When sending data to get users last activities, you send json array. First element is prettyID of the puzzle and second element in that json array is userID.
If you don't set valid prettyID, the response will be nil.

```json
[ "[prettyId]", "[userID]" ]
```

## Success Response

**Code** : `200 OK`

**Content example** :

```json
{
    "id": 148,
    "level": "multi",
    "rank": -1,
    "thumbnailBinaryId": 4989863231600,
    "previewBinaryId": 3825971694298,
    "coverBinaryId": 4769442240713,
    "logoBinaryId": 4768637804814,
    "title": "Coders Strike Back",
    "titleMap": {
        "1": "Coders Strike Back",
        "2": "Coders Strike Back"
    },
    "statement": "<div class=\"statement_back\" id=\"statement_back\" style=\"display:none\">&nbsp;</div>\n\n<div class=\"statement-body\"><!-- GOAL -->\n<div class=\"statement-section statement-goal\">\n<h1><span class=\"icon icon-goal\">&nbsp;</span> <span>The Goal</span></h1>\n\n<div class=\"statement-goal-content\">Win the race.</div>\n</div>\n<!-- RULES -->\n\n<div class=\"statement-section statement-rules\">\n<h1><span class=\"icon icon-rules\">&nbsp;</span> <span>Rules</span></h1>\n\n<div>\n<div class=\"statement-rules-content\">The circuit of the race is made up of checkpoints. To complete one lap, your vehicle (pod) must pass through each checkpoints <b>in order</b> and back through the start. The first player to reach the starting checkpoint on the final lap wins.<br />\n<br />\nThe pods work as follows:\n<ul>\n\t<li>To pass a checkpoint, the center of a pod must be inside the radius of the checkpoint.</li>\n\t<li>To move a pod, you must print a target <b>destination point</b> followed by a thrust value. Details of the protocol can be found further down.</li>\n\t<li>The thrust value of a pod is its acceleration and must be between <const>0</const> and <const>100</const>.</li>\n</ul>\n</div>\n<!-- Victory conditions -->\n\n<div class=\"statement-victory-conditions\">\n<div class=\"icon victory\">&nbsp;</div>\n\n<div class=\"blk\">\n<div class=\"title\">Victory Conditions</div>\n\n<div class=\"text\">Be the first to complete all the laps of the circuit with your pod.</div>\n</div>\n</div>\n<!-- Lose conditions -->\n\n<div class=\"statement-lose-conditions\">\n<div class=\"icon lose\">&nbsp;</div>\n\n<div class=\"blk\">\n<div class=\"title\">Lose Conditions</div>\n\n<div class=\"text\">\n<ul style=\"padding-bottom: 0; margin-bottom: 0;\">\n\t<li>Your program provides incorrect output.</li>\n\t<li>Your program times out.</li>\n\t<li>Your pod does not reach the next checkpoint in time.</li>\n\t<li>Somebody else wins.</li>\n</ul>\n</div>\n</div>\n</div>\n</div>\n</div>\n<!-- EXPERT RULES -->\n\n<div class=\"statement-section statement-expertrules\">\n<h1><span class=\"icon icon-expertrules\">&nbsp;</span> <span>Expert Rules</span></h1>\n\n<div class=\"statement-expert-rules-content\">The game is played on a map <const>16000</const> units wide and <const>9000</const> units high. The coordinate <const>X=0, Y=0</const> is the top left pixel.<br />\n<br />\nThe checkpoints work as follows:\n<ul>\n\t<li>The checkpoints are circular, with a radius of <const>600</const> units.</li>\n\t<li>The disposition of the checkpoints is selected randomly for each race.</li>\n</ul>\nThe pods work as follows:\n\n<ul>\n\t<li>If your pod does not reach the next checkpoint under <const>100</const> turns, you are <b>eliminated</b> and lose the game.</li>\n\t<li>The pods may move normally outside the game area.</li>\n</ul>\n<br />\nNote: You may activate debug mode in the settings panel (<img height=\"18\" src=\"https://www.codingame.com/servlet/fileservlet?id=3463235186409\" style=\"opacity: 0.8; background: #20252a;\" width=\"18\" />) to view additional game data.</div>\n</div>\n<!-- WARNING -->\n\n<div class=\"statement-section statement-warning\">\n<h1><span class=\"icon icon-warning\">&nbsp;</span> <span>Note</span></h1>\n\n<div class=\"statement-warning-content\">The program must,<strong> within an infinite loop</strong>, read the contextual data from the standard input and provide to the standard output the desired instructions.</div>\n</div>\n<!-- PROTOCOL -->\n\n<div class=\"statement-section statement-protocol\">\n<h1><span class=\"icon icon-protocol\">&nbsp;</span> <span>Game Input</span></h1>\n<!-- Protocol block -->\n\n<div class=\"blk\">\n<div class=\"title\">Input for one game turn</div>\n\n<div class=\"text\"><span class=\"statement-lineno\">One line</span>: 4 integers <var>x</var> &amp; <var>y</var> for the position. <var>nextCheckPointX</var> &amp; <var>nextCheckPointY</var> for the number of the next checkpoint the pod must go through.</div>\n</div>\n<!-- Protocol block -->\n\n<div class=\"blk\">\n<div class=\"title\">Output for one game turn</div>\n\n<div class=\"text\"><span class=\"statement-lineno\">One line</span>: 2 integers for the target coordinates of your pod followed by <var>thrust</var>, the thrust to give your pod.</div>\n</div>\n<!-- Protocol block -->\n\n<div class=\"blk\">\n<div class=\"title\">Constraints</div>\n\n<div class=\"text\"><const>0</const> &le; <var>thrust</var> &le; 1<const>00</const><br />\nResponse time first turn &le; 1000ms<br />\nResponse time per turn &le; 75ms</div>\n</div>\n</div>\n</div>\n",
    "validatorScore": 0,
    "achievementCount": 0,
    "doneAchievementCount": 0,
    "linkedAchievements": [],
    "lastActivity": 1598567645627,
    "forumLink": "coders-strike-back-puzzle-discussion/1833",
    "chatRoom": "csb",
    "score": 2.38,
    "position": 22854,
    "total": 23510,
    "globalTotal": 104571,
    "solvedCount": 0,
    "attemptCount": 147119,
    "feedback": {
        "feedbackId": 47,
        "codingamerFeedback": 5,
        "feedbacks": [
            36,
            20,
            48,
            264,
            2313
        ]
    },
    "optimCriteriaId": "rank",
    "topics": [
        {
            "handle": "algorithms",
            "value": "Algorithms",
            "children": [
                {
                    "handle": "optimization",
                    "value": "Optimization",
                    "children": []
                }
            ]
        },
        {
            "handle": "artificial-intelligence",
            "value": "Artificial Intelligence",
            "children": [
                {
                    "handle": "multi-agent",
                    "value": "Multi-agent",
                    "children": []
                },
                {
                    "handle": "neural-network",
                    "value": "Neural network",
                    "children": []
                }
            ]
        },
        {
            "handle": "maths-physics",
            "value": "Maths / Physics",
            "children": [
                {
                    "handle": "distance",
                    "value": "Distances",
                    "children": []
                },
                {
                    "handle": "trigonometry",
                    "value": "Trigonometry",
                    "children": []
                }
            ]
        }
    ],
    "creationTime": 1460655472759,
    "league": {
        "divisionIndex": 3,
        "divisionCount": 7,
        "openingLeaguesCount": 0,
        "divisionOffset": 0
    },
    "type": "ARENA",
    "prettyId": "coders-strike-back",
    "detailsPageUrl": "/multiplayer/bot-programming/coders-strike-back",
    "replayIds": [
        124686535,
        124686828,
        124811607,
        127192976
    ],
    "testSessionHandle": "30621524d5f00457620f93dbb0f42ac8c8486893",
    "puzzleLeaderboardId": "coders-strike-back",
    "contentDetails": {
        "description": "This puzzle is divided into two distinct parts:\n\nFrom Wood to Silver League: A tutorial containing several missions for newcomers to the multiplayer mode with a simple goal: your program must win the race. You'll unlock new rules at every league.\n\nFrom Gold league upwards: You will be given a large number of parameters to manage (list of waypoints, vector speed, remaining boosts ...) in order to improve your AI.",
        "learnDescription": "This puzzle game starts with a step by step tutorial that will help you get familiar with CodinGame’s multiplayer games. It provides an easy introduction to bot programming through a starship race. \n\nThe aim of the game is of course to win the race against other players! To succeed in this challenge, you will be able to use different mathematical concepts such as trajectory calculation, collisions, speed vector, or inertia.\n\nThe game is very simple to start. Rules are easy to understand and it only requires a few lines of code to move your ship around.\n\nHowever, it has near-infinite possibilities of evolution as you can improve your artificial intelligence step by step, while sharpening your coding skills.\n\n<h2>LEARN ALGORITHMS ASSOCIATED WITH THIS PUZZLE</h2>\n<a target=\"blank\" href=\"https://tech.io/playgrounds/334/genetic-algorithms?utm_source=codingame&utm_medium=details-page&utm_campaign=puzzle-to-playground&utm_content=coders-strike-back\">Genetic Algorithms</a> by <a href=\"/profile/301b82e88dc8149b604e41696584ddac406848\">Sablier</a>\n<br/><br/>\n<a target=\"blank\" href=\"https://www.codingame.com/playgrounds/36476/smitsimax\">Smitsimax</a> par <a href=\"/profile/04d6badfff034762c87be88072d7d6840902252\">MSmits</a>",
        "story": "",
        "externalResources": [
            {
                "publicId": "pid-control",
                "url": "https://en.wikipedia.org/wiki/PID_controller",
                "name": "PID Controller"
            },
            {
                "publicId": "multi-agent",
                "url": "https://en.wikipedia.org/wiki/Multi-agent_system",
                "name": "Multi-agent system"
            },
            {
                "publicId": "steering-behaviors",
                "url": "https://gamedevelopment.tutsplus.com/series/understanding-steering-behaviors--gamedev-12732",
                "name": "Steering Behaviors (Seek)"
            },
            {
                "publicId": "neural-network-forum",
                "url": "https://www.codingame.com/forum/t/neural-network-ressources/1667",
                "name": "Neural Network Resources"
            },
            {
                "publicId": "csb-magus",
                "url": "http://files.magusgeek.com/csb/csb_en.html",
                "name": "Genetic Algorithm post-mortem by Magus"
            },
            {
                "publicId": "csb-jeff06",
                "url": "https://www.codingame.com/blog/genetic-algorithms-coders-strike-back-game/?utm_source=codingame&utm_medium=details-page&utm_campaign=cg-blog&utm_content=csb",
                "name": "Genetic Algorithm post-mortem by Jeff06"
            },
            {
                "publicId": "csb-pb4608",
                "url": "https://www.codingame.com/blog/evolutionary-trajectory-optimization/?utm_source=codingame&utm_medium=details-page&utm_campaign=cg-blog&utm_content=csb",
                "name": "Genetic Algorithm post-mortem by pb4608"
            }
        ]
    },
    "communityCreation": false
}
```
