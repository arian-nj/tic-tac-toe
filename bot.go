package main

func bot_move(table *Table) bool {
	for _, c := range table.Cells {
		if c.Value == EmptyCell {
			c.Value = BotCell
			return true
		}
	}
	return false
}
