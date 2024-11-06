package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type Box struct {
	X int
	Y int
	S int
}

func (b *Box) Contains(x, y int) bool {
	if x < b.X-b.S/2 {
		return false
	}
	if x > b.X+b.S/2 {
		return false
	}
	if y < b.Y-b.S/2 {
		return false
	}
	if y > b.Y+b.S/2 {
		return false
	}
	return true
}

type Selector struct {
	dirPath      string
	currentImage string
	Overrides    map[int]int
	BoxLabels    map[int][]Box
	countMap     map[int]map[int]int
}

func NewSelector(dirPath string) *Selector {
	return &Selector{
		dirPath:   dirPath,
		Overrides: make(map[int]int),
		countMap:  make(map[int]map[int]int),
		BoxLabels: make(map[int][]Box),
	}
}

func (s *Selector) Update(currentImage string, sam *MasksLoader, sem *MasksLoader) {
	s.currentImage = currentImage
	s.countMap = map[int]map[int]int{}
	s.Overrides = map[int]int{}
	s.BoxLabels = map[int][]Box{}
	// attemp to read overrides from disc.
	f, err := os.Open(fmt.Sprintf("%s/%s.json", s.dirPath, currentImage))
	if err == nil {
		decoder := json.NewDecoder(f)
		decoder.Decode(s)
	}
	defer f.Close()
	for i, row := range sam.LabelMatrix {
		for j, cell := range row {
			if s.countMap[cell] == nil {
				s.countMap[cell] = map[int]int{}
			}
			s.countMap[cell][sem.LabelMatrix[i][j]]++
		}
	}
}

func (s *Selector) AddBox(label int, box Box) {
	s.BoxLabels[label] = append(s.BoxLabels[label], box)
	raw, _ := json.Marshal(s)
	os.WriteFile(fmt.Sprintf("%s/%s.json", s.dirPath, s.currentImage),
		raw, os.ModePerm)
}

func (s *Selector) SetOverride(k, v int) {
	s.Overrides[k] = v
	raw, _ := json.Marshal(s)
	os.WriteFile(fmt.Sprintf("%s/%s.json", s.dirPath, s.currentImage),
		raw, os.ModePerm)
}

func reverse(boxes []Box) []Box {
	c := []Box{}
	c = append(c, boxes...)
	slices.Reverse(c)
	return c
}

func (s *Selector) Select(x, y int, c int) int {
	for i := 0; i < 19; i++ {
		if boxes, ok := s.BoxLabels[i]; ok {
			for _, box := range reverse(boxes) {
				if box.Contains(x, y) {
					return i
				}
			}
		}
	}
	if r, ok := s.Overrides[c]; ok {
		return r
	}
	var categoryCount int
	var mostCommonCategory int
	for category, count := range s.countMap[c] {
		if count >= categoryCount {
			categoryCount = count
			mostCommonCategory = category
		}
	}
	return mostCommonCategory
}
