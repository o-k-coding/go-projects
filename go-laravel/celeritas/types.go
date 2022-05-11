package celeritas

import "database/sql"

type initPaths struct {
	rootPath    string
	folderNames []string
}

// type cookieConfig struct {
// 	name string
// 	lifetime string
// 	persist string
// 	secure string
// 	domain string
// }

type Database struct {
	DataType string
	Pool     *sql.DB
}
