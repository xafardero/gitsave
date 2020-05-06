package backup

type Backup interface {
	// Does a backup
	Save()
}
