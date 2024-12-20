package table

func BotMove(t *Table) bool {
	for _, c := range t.Cells {
		if c.Value == EmptyCell {
			c.Value = BotCell
			return true
		}
	}
	return false
}
