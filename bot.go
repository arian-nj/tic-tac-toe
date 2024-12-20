package main

func bot_move(table *Table) bool {
	for _, rowCell := range table.Cells {
		for _, c := range rowCell {
			if c.Value == EmptyCell {
				c.Value = BotCell
				return true
			}
		}
	}
	return false
}
