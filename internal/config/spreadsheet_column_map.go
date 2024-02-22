package config

import "github.com/360EntSecGroup-Skylar/excelize"

type SpreadsheetColumnMap string

func (scm SpreadsheetColumnMap) Index() int {
	return excelize.TitleToNumber(string(scm))
}
