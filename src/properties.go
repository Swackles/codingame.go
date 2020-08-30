package main

type Properties struct {
	TestSessionConfig          TestSessionConfig      `json:"testSessionConfig"`
	XpConfig                   XpConfig               `json:"xpConfig"`
	SeenChallengeIds           []int                  `json:"seenChallengeIds"`
	PrivacySettingsCodingame   PrivacySettings        `json:"privacySettings-codingame"`
	ChatSettings               ChatSettings           `json:"chatSettings"`
	ContributionsListLastVisit int64                  `json:"contributionsListLastVisit"`
	Abtesting                  Abtesting              `json:"abtesting"`
	IsDicordDisclaimerMinified bool                   `json:"isDicordDisclaimerMinified"`
	IdeHelperStatus            IdeHelperStatus        `json:"ideHelperStatus"`
	CookiesBannerCodingame     CookiesBannerCodingame `json:"cookiesBanner-codingame"`
}

type TestSessionConfig struct {
	PredilectionLanguage string `json:"predilectionLanguage"`
}

type XpConfig struct {
	LastTotalXp int `json:"lastTotalXp"`
}

type PrivacySettings struct {
	Facebook    bool `json:"facebook"`
	Analytics   bool `json:"analytics"`
	Advertising bool `json:"advertising"`
}

type ChatSettings struct {
	Open bool `json:"open"`
}

type Abtesting struct {
	QuickAB1 string `json:"Quick AB 1"`
}

type IdeHelperStatus struct {
	MultiPlayViews   int `json:"multiPlayViews"`
	SoloSubmitViews  int `json:"soloSubmitViews"`
	RuleChangeViews  int `json:"ruleChangeViews"`
	MultiSubmitViews int `json:"multiSubmitViews"`
}

type CookiesBannerCodingame struct {
	Seen bool `json:"seen"`
}
