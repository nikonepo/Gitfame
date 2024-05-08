package main

import (
	"errors"
	"sort"
	"strings"
)

var (
	LINES   = SortType{name: "lines"}
	COMMITS = SortType{name: "commits"}
	FILES   = SortType{name: "files"}
)

func SortAuthors(authors []Author, sort1 SortType, sort2 SortType, sort3 SortType) []Author {
	sort.Slice(authors, func(i, j int) bool {
		var a, b int
		switch sort1 {
		case LINES:
			a, b = authors[i].lines, authors[j].lines

		case COMMITS:
			a, b = authors[i].commits, authors[j].commits

		case FILES:
			a, b = len(authors[i].files), len(authors[j].files)
		}

		if a == b {
			switch sort2 {
			case LINES:
				a, b = authors[i].lines, authors[j].lines

			case COMMITS:
				a, b = authors[i].commits, authors[j].commits

			case FILES:
				a, b = len(authors[i].files), len(authors[j].files)
			}

			if a == b {
				switch sort3 {
				case LINES:
					a, b = authors[i].lines, authors[j].lines

				case COMMITS:
					a, b = authors[i].commits, authors[j].commits

				case FILES:
					a, b = len(authors[i].files), len(authors[j].files)
				}

				if a == b {
					return strings.Compare(authors[i].name, authors[j].name) < 0
				} else {
					return a > b
				}
			} else {
				return a > b
			}

		} else {
			return a > b
		}
	})

	return authors
}

func Sort(authors []Author, sortType SortType) []Author {
	switch sortType {
	case COMMITS:
		return SortAuthors(authors, COMMITS, LINES, FILES)
	case LINES:
		return SortAuthors(authors, LINES, COMMITS, FILES)
	case FILES:
		return SortAuthors(authors, FILES, LINES, COMMITS)
	}

	return authors
}

func GetSortType(name string) (SortType, error) {
	switch name {
	case LINES.name:
		return LINES, nil
	case COMMITS.name:
		return COMMITS, nil
	case FILES.name:
		return FILES, nil
	default:
		return LINES, errors.New("illegal type")
	}
}
