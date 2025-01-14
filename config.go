package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gocv.io/x/gocv"

	"github.com/pidgy/unitehud/team"
)

type filter struct {
	*team.Team
	file  string
	value int
}

type template struct {
	filter
	gocv.Mat
	category    string
	subcategory string
}

var (
	filenames map[string]map[string][]filter
	templates map[string]map[string][]template
)

func load() {
	switch game {
	case "switch":
		loadSwitch()
	case "ios":
		loadIOS()
	}

	for category := range filenames {
		for subcategory, filters := range filenames[category] {
			for _, filter := range filters {
				templates[category][filter.Team.Name] = append(templates[category][filter.Team.Name],
					template{
						filter,
						gocv.IMRead(filter.file, gocv.IMReadColor),
						category,
						subcategory,
					},
				)
			}
		}
	}

	for category := range templates {
		for _, templates := range templates[category] {
			for _, t := range templates {
				if t.Empty() {
					kill(fmt.Errorf("invalid scored template: %s", t.file))
				}

				log.Debug().Object("template", t).Msg("score template loaded")
			}
		}
	}
}

func loadSwitch() {
	filenames = map[string]map[string][]filter{
		"game": {
			"vs": {
				filter{team.None, "img/switch/game/vs.png", -0},
			},
			"end": {
				filter{team.None, "img/switch/game/end.png", -0},
			},
		},
		"scored": {
			team.Purple.Name: {
				filter{team.Purple, "img/switch/purple/score/score.png", -0},
				filter{team.Purple, "img/switch/purple/score/score_alt.png", -0},
			},
			team.Orange.Name: {
				filter{team.Orange, "img/switch/orange/score/score.png", -0},
				filter{team.Orange, "img/switch/orange/score/score_alt.png", -0},
			},
			team.Self.Name: {
				//filter{team.Self, "img/switch/self/score/score.png", -0},
				filter{team.Self, "img/switch/self/score/score_alt.png", -0},
				/*
					filter{team.Self, "img/switch/self/score/score_alt_alt.png", -0},
					filter{team.Self, "img/switch/self/score/score_alt_alt_alt.png", -0},
					filter{team.Self, "img/switch/self/score/score_alt_alt_alt_alt.png", -0},
					filter{team.Self, "img/switch/self/score/score_alt_alt.png", -0},
					filter{team.Self, "img/switch/self/score/score_big_alt.png", -0},
				*/
			},
		},
		"points": {
			team.Purple.Name: {
				filter{team.Purple, "img/switch/purple/points/point_0.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_alt.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_alt_alt.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_alt_alt_alt.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_alt_alt_alt_alt.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_alt_alt_alt_alt_alt.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_alt_alt_alt_alt_alt_alt.png", 0},

				filter{team.Purple, "img/switch/purple/points/point_0_big.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_big_alt.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_big_alt_alt.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_big_alt_alt_alt.png", 0},
				filter{team.Purple, "img/switch/purple/points/point_0_big_alt_alt_alt_alt.png", 0},

				filter{team.Purple, "img/switch/purple/points/point_1.png", 1},
				filter{team.Purple, "img/switch/purple/points/point_1_alt.png", 1},
				filter{team.Purple, "img/switch/purple/points/point_1_alt_alt.png", 1},
				filter{team.Purple, "img/switch/purple/points/point_1_big.png", 1},
				filter{team.Purple, "img/switch/purple/points/point_1_big_alt.png", 1},
				filter{team.Purple, "img/switch/purple/points/point_1_big_alt_alt.png", 1},

				filter{team.Purple, "img/switch/purple/points/point_2.png", 2},
				filter{team.Purple, "img/switch/purple/points/point_2_alt.png", 2},
				filter{team.Purple, "img/switch/purple/points/point_2_alt_alt.png", 2},
				filter{team.Purple, "img/switch/purple/points/point_2_alt_alt_alt.png", 2},
				filter{team.Purple, "img/switch/purple/points/point_2_big_alt.png", 2},

				filter{team.Purple, "img/switch/purple/points/point_3.png", 3},
				filter{team.Purple, "img/switch/purple/points/point_3_alt.png", 3},

				filter{team.Purple, "img/switch/purple/points/point_4.png", 4},
				filter{team.Purple, "img/switch/purple/points/point_4_alt.png", 4},
				filter{team.Purple, "img/switch/purple/points/point_4_alt_alt.png", 4},
				filter{team.Purple, "img/switch/purple/points/point_4_big.png", 4},
				filter{team.Purple, "img/switch/purple/points/point_4_big_alt.png", 4},
				filter{team.Purple, "img/switch/purple/points/point_4_big_alt_alt.png", 4},
				filter{team.Purple, "img/switch/purple/points/point_4_big_alt_alt_alt.png", 4},

				filter{team.Purple, "img/switch/purple/points/point_5_alt.png", 5},
				filter{team.Purple, "img/switch/purple/points/point_5_big.png", 5},

				filter{team.Purple, "img/switch/purple/points/point_6.png", 6},
				filter{team.Purple, "img/switch/purple/points/point_6_alt.png", 6},
				filter{team.Purple, "img/switch/purple/points/point_6_big.png", 6},
				filter{team.Purple, "img/switch/purple/points/point_6_big_alt.png", 6},

				filter{team.Purple, "img/switch/purple/points/point_7.png", 7},
				filter{team.Purple, "img/switch/purple/points/point_7_big.png", 7},

				filter{team.Purple, "img/switch/purple/points/point_8.png", 8},
				filter{team.Purple, "img/switch/purple/points/point_8_big.png", 8},
				filter{team.Purple, "img/switch/purple/points/point_8_big_alt.png", 8},
				filter{team.Purple, "img/switch/purple/points/point_8_big_alt_alt.png", 8},

				filter{team.Purple, "img/switch/purple/points/point_9.png", 9},
				filter{team.Purple, "img/switch/purple/points/point_9_alt.png", 9},
				filter{team.Purple, "img/switch/purple/points/point_9_big.png", 9},
			},
			team.Orange.Name: {
				filter{team.Orange, "img/switch/orange/points/point_0.png", 0},
				filter{team.Orange, "img/switch/orange/points/point_0_alt.png", 0},
				filter{team.Orange, "img/switch/orange/points/point_0_big.png", 0},
				filter{team.Orange, "img/switch/orange/points/point_0_big_alt.png", 0},
				filter{team.Orange, "img/switch/orange/points/point_0_big_alt_alt.png", 0},
				filter{team.Orange, "img/switch/orange/points/point_0_big_alt_alt_alt.png", 0},
				filter{team.Orange, "img/switch/orange/points/point_0_big_alt_alt_alt_alt.png", 0},

				filter{team.Orange, "img/switch/orange/points/point_1.png", 1},
				filter{team.Orange, "img/switch/orange/points/point_1_alt.png", 1},
				filter{team.Orange, "img/switch/orange/points/point_1_big.png", 1},
				filter{team.Orange, "img/switch/orange/points/point_1_big_alt.png", 1},

				filter{team.Orange, "img/switch/orange/points/point_2.png", 2},
				filter{team.Orange, "img/switch/orange/points/point_2_alt.png", 2},
				filter{team.Orange, "img/switch/orange/points/point_2_big_alt.png", 2},

				filter{team.Orange, "img/switch/orange/points/point_3.png", 3},
				filter{team.Orange, "img/switch/orange/points/point_3_alt.png", 3},

				filter{team.Orange, "img/switch/orange/points/point_4.png", 4},
				filter{team.Orange, "img/switch/orange/points/point_4_alt.png", 4},
				filter{team.Orange, "img/switch/orange/points/point_4_alt_alt.png", 4},
				filter{team.Orange, "img/switch/orange/points/point_4_alt_alt_alt.png", 4},
				filter{team.Orange, "img/switch/orange/points/point_4_big_alt.png", 4},

				filter{team.Orange, "img/switch/orange/points/point_5.png", 5},
				filter{team.Orange, "img/switch/orange/points/point_5_alt.png", 5},

				filter{team.Orange, "img/switch/orange/points/point_6.png", 6},
				filter{team.Orange, "img/switch/orange/points/point_6_alt.png", 6},
				filter{team.Orange, "img/switch/orange/points/point_6_alt_alt.png", 6},
				filter{team.Orange, "img/switch/orange/points/point_6_big_alt.png", 6},
				filter{team.Orange, "img/switch/orange/points/point_6_big_alt_alt.png", 6},

				filter{team.Orange, "img/switch/orange/points/point_7.png", 7},
				filter{team.Orange, "img/switch/orange/points/point_7_big.png", 7},

				filter{team.Orange, "img/switch/orange/points/point_8.png", 8},
				filter{team.Orange, "img/switch/orange/points/point_8_alt.png", 8},
				filter{team.Orange, "img/switch/orange/points/point_8_alt_alt.png", 8},
				filter{team.Orange, "img/switch/orange/points/point_8_big_alt.png", 8},

				filter{team.Orange, "img/switch/orange/points/point_9.png", 9},
				filter{team.Orange, "img/switch/orange/points/point_9_alt.png", 9},
				filter{team.Orange, "img/switch/orange/points/point_9_big.png", 9},
			},
			team.Self.Name: {
				filter{team.Self, "img/switch/self/points/point_0.png", 0},
				filter{team.Self, "img/switch/self/points/point_0_alt.png", 0},
				filter{team.Self, "img/switch/self/points/point_0_alt_alt.png", 0},
				filter{team.Self, "img/switch/self/points/point_0_alt_alt_alt.png", 0},
				filter{team.Self, "img/switch/self/points/point_1.png", 1},
				filter{team.Self, "img/switch/self/points/point_1_alt.png", 1},
				filter{team.Self, "img/switch/self/points/point_2.png", 2},
				filter{team.Self, "img/switch/self/points/point_2_alt.png", 2},
				filter{team.Self, "img/switch/self/points/point_5.png", 5},
				filter{team.Self, "img/switch/self/points/point_5_alt.png", 5},
				filter{team.Self, "img/switch/self/points/point_5_alt_alt.png", 5},
				filter{team.Self, "img/switch/self/points/point_5_alt_alt_alt.png", 5},
				filter{team.Self, "img/switch/self/points/point_5_alt_alt_alt_alt.png", 5},
				filter{team.Self, "img/switch/self/points/point_6.png", 6},
				filter{team.Self, "img/switch/self/points/point_6_alt.png", 6},
				filter{team.Self, "img/switch/self/points/point_7.png", 7},
				filter{team.Self, "img/switch/self/points/point_7_alt.png", 7},
				filter{team.Self, "img/switch/self/points/point_7_alt_alt.png", 7},
				filter{team.Self, "img/switch/self/points/point_8_alt.png", 8},
			},
		},
		"time": {
			team.Time.Name: {
				filter{team.Time, "img/switch/time/points/point_0.png", 0},
				filter{team.Time, "img/switch/time/points/point_1.png", 1},
				filter{team.Time, "img/switch/time/points/point_2.png", 2},
				filter{team.Time, "img/switch/time/points/point_3.png", 3},
				filter{team.Time, "img/switch/time/points/point_4.png", 4},
				filter{team.Time, "img/switch/time/points/point_5.png", 5},
				filter{team.Time, "img/switch/time/points/point_6.png", 6},
				filter{team.Time, "img/switch/time/points/point_7.png", 7},
				filter{team.Time, "img/switch/time/points/point_8.png", 8},
				filter{team.Time, "img/switch/time/points/point_9.png", 9},
			},
		},
	}

	templates = map[string]map[string][]template{
		"game": {
			team.None.Name: {},
		},
		"scored": {
			team.Orange.Name: {},
			team.Purple.Name: {},
			team.Self.Name:   {},
		},
		"points": {
			team.Orange.Name: {},
			team.Purple.Name: {},
			team.Self.Name:   {},
		},
		"time": {
			team.Time.Name: {},
		},
	}
}

func loadIOS() {
	filenames = map[string]map[string][]filter{
		"game": {
			"vs":  {},
			"end": {},
		},
		"scored": {
			team.Purple.Name: {},
			team.Orange.Name: {},
			team.Self.Name:   {},
		},
		"points": {
			team.Purple.Name: {},
			team.Orange.Name: {},
			team.Self.Name:   {},
		},
		"time": {
			team.Time.Name: {
				filter{team.Time, "img/ios/time/points/point_0.png", 0},
				filter{team.Time, "img/ios/time/points/point_0_alt.png", 0},

				filter{team.Time, "img/ios/time/points/point_1.png", 1},
				filter{team.Time, "img/ios/time/points/point_1_alt.png", 1},

				filter{team.Time, "img/ios/time/points/point_2.png", 2},
				filter{team.Time, "img/ios/time/points/point_2_alt.png", 2},

				filter{team.Time, "img/ios/time/points/point_3.png", 3},
				filter{team.Time, "img/ios/time/points/point_3_alt.png", 3},

				filter{team.Time, "img/ios/time/points/point_4.png", 4},
				filter{team.Time, "img/ios/time/points/point_4_alt.png", 4},

				filter{team.Time, "img/ios/time/points/point_5.png", 5},
				filter{team.Time, "img/ios/time/points/point_5_alt.png", 5},

				filter{team.Time, "img/ios/time/points/point_6.png", 6},
				filter{team.Time, "img/ios/time/points/point_6_alt.png", 6},

				filter{team.Time, "img/ios/time/points/point_7.png", 7},
				filter{team.Time, "img/ios/time/points/point_7_alt.png", 7},

				filter{team.Time, "img/ios/time/points/point_8.png", 8},
				filter{team.Time, "img/ios/time/points/point_8_alt.png", 8},

				filter{team.Time, "img/ios/time/points/point_9.png", 9},
				filter{team.Time, "img/ios/time/points/point_9_alt.png", 9},
			},
		},
	}

	templates = map[string]map[string][]template{
		"game": {
			team.None.Name: {},
		},
		"scored": {
			team.Orange.Name: {},
			team.Purple.Name: {},
			team.Self.Name:   {},
		},
		"points": {
			team.Orange.Name: {},
			team.Purple.Name: {},
			team.Self.Name:   {},
		},
		"time": {
			team.Time.Name: {},
		},
	}
}
