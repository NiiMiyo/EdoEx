package edopro

import (
	"database/sql"

	"edoex/embedfiles"
	"edoex/environment"
	"edoex/logger"

	_ "modernc.org/sqlite"
)

// Creates expansion-name.cdb
func WriteToCdb() error {
	logger.Verbosef("Connecting to database '%s'", environment.EdoproDatabaseBuildPath())
	db, err := sql.Open("sqlite", environment.EdoproDatabaseBuildPath())
	if err != nil {
		return err
	}
	defer db.Close()

	logger.Verbose("Creating database tables")
	db.Exec(embedfiles.CreateTablesScriptEdopro)

	for _, c := range environment.Cards {
		logger.Verbosef("Inserting card '%s' (%d)", c.Name, c.Id)
		cdb := c.ToDb()

		logger.Verbose("Preparing insert")
		stmt, err := db.Prepare(
			"insert into datas(id, ot, alias, setcode, type, atk, def, level, race, attribute, category) values (?,?,?,?,?,?,?,?,?,?,?)",
		)
		if err != nil {
			return err
		}

		logger.Verbose("Inserting card")
		_, err = stmt.Exec(cdb.Id, cdb.Ot, cdb.Alias, cdb.Setcode, cdb.Type, cdb.Atk, cdb.Def, cdb.Level, cdb.Race, cdb.Attribute, cdb.Category)
		if err != nil {
			return err
		}

		logger.Verbose("Preparing texts")
		stmt, err = db.Prepare("insert into texts(id, name, desc, str1, str2, str3, str4, str5, str6, str7, str8, str9, str10, str11, str12, str13, str14, str15, str16) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			return err
		}

		logger.Verbose("Inserting texts")
		_, err = stmt.Exec(cdb.Id, cdb.Name, cdb.Desc, cdb.Strings[0], cdb.Strings[1], cdb.Strings[2], cdb.Strings[3], cdb.Strings[4], cdb.Strings[5], cdb.Strings[6], cdb.Strings[7], cdb.Strings[8], cdb.Strings[9], cdb.Strings[10], cdb.Strings[11], cdb.Strings[12], cdb.Strings[13], cdb.Strings[14], cdb.Strings[15])
		if err != nil {
			return err
		}
	}

	return nil
}
