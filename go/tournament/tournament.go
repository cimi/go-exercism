package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	WIN  string = "win"
	LOSS string = "loss"
	DRAW string = "draw"
)

type TeamStats struct {
	wins   int
	losses int
	draws  int
}

func (t *TeamStats) GamesPlayed() int {
	return t.wins + t.losses + t.draws
}

func (t *TeamStats) Points() int {
	return t.wins*3 + t.draws
}

type Team struct {
	name  string
	stats *TeamStats
}

const FORMAT_STRING = "%s | %2d | %2d | %2d | %2d | %2d"

func (t *Team) String() string {
	return fmt.Sprintf(FORMAT_STRING,
		t.name+strings.Repeat(" ", 30-len(t.name)),
		t.stats.GamesPlayed(),
		t.stats.wins,
		t.stats.draws,
		t.stats.losses,
		t.stats.Points())
}

type Table map[string]*TeamStats

func (t Table) Ranking() []Team {
	teams := make([]Team, 0)
	for k, v := range t {
		teams = append(teams, Team{name: k, stats: v})
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].stats.Points() == teams[j].stats.Points() {
			return teams[i].name < teams[j].name
		}
		return teams[i].stats.Points() > teams[j].stats.Points()
	})
	return teams
}

func (t Table) addTeam(name string) {
	if _, ok := t[name]; !ok {
		t[name] = &TeamStats{}
	}
}

func (t Table) addTeams(info []string) {
	t.addTeam(info[0])
	t.addTeam(info[1])
}

func (t Table) update(info []string) {
	t.addTeams(info)
	switch {
	case info[2] == WIN:
		t[info[0]].wins += 1
		t[info[1]].losses += 1
	case info[2] == LOSS:
		t[info[1]].wins += 1
		t[info[0]].losses += 1
	case info[2] == DRAW:
		t[info[0]].draws += 1
		t[info[1]].draws += 1
	}
}

type Input struct {
	ignored bool
	info    []string
	err     error
}

func isIgnored(l string) bool {
	return l == "" || strings.Index(l, "#") == 0
}

func isValidOutcome(outcome string) bool {
	switch outcome {
	case
		WIN,
		LOSS,
		DRAW:
		return true
	}
	return false
}

func ParseLine(l string) Input {
	if isIgnored(l) {
		return Input{ignored: true}
	}
	input := Input{info: strings.Split(l, ";")}
	if len(input.info) != 3 {
		input.err = fmt.Errorf("Invalid input: %s", l)
	} else if !isValidOutcome(input.info[2]) {
		input.err = fmt.Errorf("Invalid match outcome: %s", input.info[2])
	}
	return input
}

func Tally(reader io.Reader, writer io.Writer) error {
	table := make(Table)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		input := ParseLine(strings.TrimSpace(scanner.Text()))
		if input.ignored {
			continue
		}
		if input.err != nil {
			return input.err
		}
		table.update(input.info)
	}

	w := bufio.NewWriter(writer)
	fmt.Fprintln(w, "Team                           | MP |  W |  D |  L |  P")
	for _, team := range table.Ranking() {
		fmt.Fprintln(w, team.String())
	}
	w.Flush()
	return nil
}
