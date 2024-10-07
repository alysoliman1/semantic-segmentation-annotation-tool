package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type MasksLoader struct {
	dirPath      string
	currentImage string
	LabelMatrix  [][]int
	Labels       []int
}

func NewMasksLoader(dirPath string) *MasksLoader {
	return &MasksLoader{
		dirPath: dirPath,
	}
}

func (m *MasksLoader) SetCurrentImage(image string) {
	if image == m.currentImage {
		return
	}
	m.currentImage = image

	// Load label map from disc.
	f, _ := os.Open(fmt.Sprintf("%s/%s.csv", m.dirPath, image))
	reader := csv.NewReader(f)
	masks, _ := reader.ReadAll()
	m.LabelMatrix = nil
	labelSet := map[int]struct{}{}
	for _, rawRow := range masks {
		var row []int
		for _, rawLabel := range rawRow {
			label, _ := strconv.Atoi(rawLabel)
			labelSet[label] = struct{}{}
			row = append(row, label)
		}
		m.LabelMatrix = append(m.LabelMatrix, row)
	}
	m.Labels = nil
	for label := range labelSet {
		m.Labels = append(m.Labels, label)
	}
}
