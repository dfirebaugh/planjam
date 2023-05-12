package plan

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/dfirebaugh/planjam/pkg/strfmt"
	"gopkg.in/yaml.v3"
)

type Lane struct {
	Label    string             `yaml:"label"`
	Features []FeatureReference `yaml:"features"`
}

type Board struct {
	Lanes           []Lane `yaml:"lanes"`
	ShowFeatureIDs  bool   `yaml:"-"`
	ShowLaneNumbers bool   `yaml:"-"`
}

// NewBoardFromYAML parses the given YAML data and returns a new Board.
func GetBoard(data []byte) (Board, error) {
	var board Board
	var parsedLanes []Lane
	err := yaml.Unmarshal(data, &parsedLanes)

	var formattedLanes []Lane

	for _, l := range parsedLanes {
		formattedLanes = append(formattedLanes, fmtLane(l))
	}

	board.Lanes = formattedLanes
	return board, err
}

func fmtLabel(label string) string {
	return strings.ReplaceAll(label, " ", "_")
}

func fmtLane(lane Lane) Lane {
	lane.Label = fmtLabel(lane.Label)
	return lane
}

func (b Board) Table() *strfmt.Table {
	table := strfmt.NewTable()
	columnIndex := 0
	for laneIndex, l := range b.Lanes {
		var laneLabel string
		if b.ShowLaneNumbers {
			laneLabel = fmt.Sprintf(" [%d] ", laneIndex)
		}
		laneLabel += l.Label
		table.AppendToColumn(columnIndex, laneLabel)
		for _, i := range l.Features {
			FeatureLabel := ""
			if b.ShowFeatureIDs {
				FeatureLabel += fmt.Sprintf(" [%s] ", i.ID)
			}
			FeatureLabel += i.Label
			table.AppendToColumn(columnIndex, FeatureLabel)
		}
		columnIndex++
	}

	return table
}

func (b Board) ASCIIDocTable() string {
	return b.Table().ASCIIDocTable()
}

func (b Board) String() string {
	return b.Table().String()
}

func (b Board) Write(w io.Writer) error {
	raw, err := yaml.Marshal(b.Lanes)
	if err != nil {
		return err
	}

	_, err = w.Write(raw)
	return err
}

func (b Board) GetLane(index int) Lane {
	if index >= len(b.Lanes) {
		return Lane{}
	}
	return b.Lanes[index]
}

func (b Board) GetFeatureRef(id string) FeatureReference {
	for _, l := range b.Lanes {
		for _, i := range l.Features {
			if i.ID != id {
				continue
			}
			if i.ID == id {
				return i
			}
		}
	}
	return FeatureReference{}
}

func (b Board) CountFeatures() int {
	count := 0
	for _, l := range b.Lanes {
		count += len(l.Features)
	}
	return count
}

func (b Board) Move(id string, laneLabel string) Board {
	Feature := b.GetFeatureRef(id)
	var board Board

	laneExists := false
	for _, l := range b.Lanes {
		if l.Label == laneLabel {
			laneExists = true
		}
	}

	if !laneExists {
		return b
	}

	for _, l := range b.Lanes {
		newLane := Lane{
			Features: []FeatureReference{},
			Label:    l.Label,
		}
		for _, iss := range l.Features {
			if iss.ID == id {
				continue
			}
			newLane.Features = append(newLane.Features, iss)
		}

		if l.Label == laneLabel {
			newLane.Features = append(newLane.Features, Feature)
		}
		board.Lanes = append(board.Lanes, newLane)
	}
	return board
}

func (b Board) GetLaneByLabel(label string) Lane {
	for _, l := range b.Lanes {
		if l.Label != label {
			continue
		}
		return l
	}

	return Lane{}
}

func (b Board) AddLane(label string) Board {
	l := b.GetLaneByLabel(label)
	if l.Label == label {
		println("lane already exists")
		return b
	}
	b.Lanes = append(b.Lanes, Lane{
		Label: label,
	})

	return b
}

func (b Board) AddFeature(label string) Board {
	b.Lanes[0].Features = append(b.Lanes[0].Features, FeatureReference{
		ID:    strconv.Itoa(b.CountFeatures()),
		Label: fmtLabel(label),
	})
	return b
}

func (b Board) AddField(FeatureLabel string, fieldLabel string) Board {
	return b
}

func (b *Board) RemoveFeature(id string) {
	var lanes []Lane

	for _, l := range b.Lanes {
		var lane Lane
		var features []FeatureReference
		for _, f := range l.Features {
			if f.ID == id {
				continue
			}

			features = append(features, f)
		}
		lane.Label = l.Label
		lane.Features = features

		lanes = append(lanes, lane)
	}

	b.Lanes = lanes
}

func (b Board) MoveLane(position int) Board {
	return b
}
