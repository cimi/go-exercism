package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
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

func heading() string {
	return "Team                           | MP |  W |  D |  L |  P"
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

func (t Table) update(info []string) error {
	t.addTeams(info)
	switch info[2] {
	case "win":
		t[info[0]].wins += 1
		t[info[1]].losses += 1
	case "loss":
		t[info[1]].wins += 1
		t[info[0]].losses += 1
	case "draw":
		t[info[0]].draws += 1
		t[info[1]].draws += 1
	default:
		return fmt.Errorf("Invalid match outcome: %s", info[2])
	}
	return nil
}

func isComment(line string) bool {
	return strings.Index(line, "#") == 0
}

func (t Table) parse(line string) error {
	if isComment(line) {
		return nil
	}
	info := strings.Split(line, ";")
	if len(info) != 3 {
		return fmt.Errorf("Invalid input: %s", line)
	}
	return t.update(info)
}

func Tally(reader io.Reader, writer io.Writer) error {
	table := make(Table)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if err := table.parse(line); line != "" && err != nil {
			return err
		}
	}

	w := bufio.NewWriter(writer)
	fmt.Fprintln(w, heading())
	for _, team := range table.Ranking() {
		fmt.Fprintln(w, team.String())
	}
	w.Flush()
	return nil
}
