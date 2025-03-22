package app

// Init is a dummy command because config init is
// called for all commands. In the future it will
// probably initialize the remove
func (a *App) Init() error {
	return nil
}
