// Package models contains generated code for schema 'public'.
package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// CodeSystem represents a row from 'public.code_system'.
type CodeSystem struct {
	Oid                           string         `json:"oid"`                           // oid
	ID                            string         `json:"id"`                            // id
	Name                          string         `json:"name"`                          // name
	Definitiontext                sql.NullString `json:"definitiontext"`                // definitiontext
	Status                        string         `json:"status"`                        // status
	Version                       string         `json:"version"`                       // version
	Versiondescription            sql.NullString `json:"versiondescription"`            // versiondescription
	Assigningauthorityversionname sql.NullString `json:"assigningauthorityversionname"` // assigningauthorityversionname
	Distributionsourceversionname sql.NullString `json:"distributionsourceversionname"` // distributionsourceversionname
	Distributionsourceid          string         `json:"distributionsourceid"`          // distributionsourceid
	Assigningauthorityid          string         `json:"assigningauthorityid"`          // assigningauthorityid
	Codesystemcode                string         `json:"codesystemcode"`                // codesystemcode
	Sourceurl                     sql.NullString `json:"sourceurl"`                     // sourceurl
	Hl70396identifier             string         `json:"hl70396identifier"`             // hl70396identifier
	Legacyflag                    bool           `json:"legacyflag"`                    // legacyflag
	Statusdate                    time.Time      `json:"statusdate"`                    // statusdate
	Acquireddate                  sql.NullTime   `json:"acquireddate"`                  // acquireddate
	Effectivedate                 sql.NullTime   `json:"effectivedate"`                 // effectivedate
	Expirydate                    sql.NullTime   `json:"expirydate"`                    // expirydate
	Assigningauthorityreleasedate sql.NullTime   `json:"assigningauthorityreleasedate"` // assigningauthorityreleasedate
	Distributionsourcereleasedate sql.NullTime   `json:"distributionsourcereleasedate"` // distributionsourcereleasedate
	Sdocreatedate                 sql.NullTime   `json:"sdocreatedate"`                 // sdocreatedate
	Lastrevisiondate              time.Time      `json:"lastrevisiondate"`              // lastrevisiondate
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [CodeSystem] exists in the database.
func (cs *CodeSystem) Exists() bool {
	return cs._exists
}

// Deleted returns true when the [CodeSystem] has been marked for deletion
// from the database.
func (cs *CodeSystem) Deleted() bool {
	return cs._deleted
}

// Insert inserts the [CodeSystem] to the database.
func (cs *CodeSystem) Insert(ctx context.Context, db DB) error {
	switch {
	case cs._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case cs._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.code_system (` +
		`oid, id, name, definitiontext, status, version, versiondescription, assigningauthorityversionname, distributionsourceversionname, distributionsourceid, assigningauthorityid, codesystemcode, sourceurl, hl70396identifier, legacyflag, statusdate, acquireddate, effectivedate, expirydate, assigningauthorityreleasedate, distributionsourcereleasedate, sdocreatedate, lastrevisiondate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23` +
		`)`
	// run
	logf(sqlstr, cs.Oid, cs.ID, cs.Name, cs.Definitiontext, cs.Status, cs.Version, cs.Versiondescription, cs.Assigningauthorityversionname, cs.Distributionsourceversionname, cs.Distributionsourceid, cs.Assigningauthorityid, cs.Codesystemcode, cs.Sourceurl, cs.Hl70396identifier, cs.Legacyflag, cs.Statusdate, cs.Acquireddate, cs.Effectivedate, cs.Expirydate, cs.Assigningauthorityreleasedate, cs.Distributionsourcereleasedate, cs.Sdocreatedate, cs.Lastrevisiondate)
	if _, err := db.ExecContext(ctx, sqlstr, cs.Oid, cs.ID, cs.Name, cs.Definitiontext, cs.Status, cs.Version, cs.Versiondescription, cs.Assigningauthorityversionname, cs.Distributionsourceversionname, cs.Distributionsourceid, cs.Assigningauthorityid, cs.Codesystemcode, cs.Sourceurl, cs.Hl70396identifier, cs.Legacyflag, cs.Statusdate, cs.Acquireddate, cs.Effectivedate, cs.Expirydate, cs.Assigningauthorityreleasedate, cs.Distributionsourcereleasedate, cs.Sdocreatedate, cs.Lastrevisiondate); err != nil {
		return logerror(err)
	}
	// set exists
	cs._exists = true
	return nil
}

// Update updates a [CodeSystem] in the database.
func (cs *CodeSystem) Update(ctx context.Context, db DB) error {
	switch {
	case !cs._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case cs._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.code_system SET ` +
		`id = $1, name = $2, definitiontext = $3, status = $4, version = $5, versiondescription = $6, assigningauthorityversionname = $7, distributionsourceversionname = $8, distributionsourceid = $9, assigningauthorityid = $10, codesystemcode = $11, sourceurl = $12, hl70396identifier = $13, legacyflag = $14, statusdate = $15, acquireddate = $16, effectivedate = $17, expirydate = $18, assigningauthorityreleasedate = $19, distributionsourcereleasedate = $20, sdocreatedate = $21, lastrevisiondate = $22 ` +
		`WHERE oid = $23`
	// run
	logf(sqlstr, cs.ID, cs.Name, cs.Definitiontext, cs.Status, cs.Version, cs.Versiondescription, cs.Assigningauthorityversionname, cs.Distributionsourceversionname, cs.Distributionsourceid, cs.Assigningauthorityid, cs.Codesystemcode, cs.Sourceurl, cs.Hl70396identifier, cs.Legacyflag, cs.Statusdate, cs.Acquireddate, cs.Effectivedate, cs.Expirydate, cs.Assigningauthorityreleasedate, cs.Distributionsourcereleasedate, cs.Sdocreatedate, cs.Lastrevisiondate, cs.Oid)
	if _, err := db.ExecContext(ctx, sqlstr, cs.ID, cs.Name, cs.Definitiontext, cs.Status, cs.Version, cs.Versiondescription, cs.Assigningauthorityversionname, cs.Distributionsourceversionname, cs.Distributionsourceid, cs.Assigningauthorityid, cs.Codesystemcode, cs.Sourceurl, cs.Hl70396identifier, cs.Legacyflag, cs.Statusdate, cs.Acquireddate, cs.Effectivedate, cs.Expirydate, cs.Assigningauthorityreleasedate, cs.Distributionsourcereleasedate, cs.Sdocreatedate, cs.Lastrevisiondate, cs.Oid); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [CodeSystem] to the database.
func (cs *CodeSystem) Save(ctx context.Context, db DB) error {
	if cs.Exists() {
		return cs.Update(ctx, db)
	}
	return cs.Insert(ctx, db)
}

// Upsert performs an upsert for [CodeSystem].
func (cs *CodeSystem) Upsert(ctx context.Context, db DB) error {
	switch {
	case cs._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.code_system (` +
		`oid, id, name, definitiontext, status, version, versiondescription, assigningauthorityversionname, distributionsourceversionname, distributionsourceid, assigningauthorityid, codesystemcode, sourceurl, hl70396identifier, legacyflag, statusdate, acquireddate, effectivedate, expirydate, assigningauthorityreleasedate, distributionsourcereleasedate, sdocreatedate, lastrevisiondate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23` +
		`)` +
		` ON CONFLICT (oid) DO ` +
		`UPDATE SET ` +
		`id = EXCLUDED.id, name = EXCLUDED.name, definitiontext = EXCLUDED.definitiontext, status = EXCLUDED.status, version = EXCLUDED.version, versiondescription = EXCLUDED.versiondescription, assigningauthorityversionname = EXCLUDED.assigningauthorityversionname, distributionsourceversionname = EXCLUDED.distributionsourceversionname, distributionsourceid = EXCLUDED.distributionsourceid, assigningauthorityid = EXCLUDED.assigningauthorityid, codesystemcode = EXCLUDED.codesystemcode, sourceurl = EXCLUDED.sourceurl, hl70396identifier = EXCLUDED.hl70396identifier, legacyflag = EXCLUDED.legacyflag, statusdate = EXCLUDED.statusdate, acquireddate = EXCLUDED.acquireddate, effectivedate = EXCLUDED.effectivedate, expirydate = EXCLUDED.expirydate, assigningauthorityreleasedate = EXCLUDED.assigningauthorityreleasedate, distributionsourcereleasedate = EXCLUDED.distributionsourcereleasedate, sdocreatedate = EXCLUDED.sdocreatedate, lastrevisiondate = EXCLUDED.lastrevisiondate `
	// run
	logf(sqlstr, cs.Oid, cs.ID, cs.Name, cs.Definitiontext, cs.Status, cs.Version, cs.Versiondescription, cs.Assigningauthorityversionname, cs.Distributionsourceversionname, cs.Distributionsourceid, cs.Assigningauthorityid, cs.Codesystemcode, cs.Sourceurl, cs.Hl70396identifier, cs.Legacyflag, cs.Statusdate, cs.Acquireddate, cs.Effectivedate, cs.Expirydate, cs.Assigningauthorityreleasedate, cs.Distributionsourcereleasedate, cs.Sdocreatedate, cs.Lastrevisiondate)
	if _, err := db.ExecContext(ctx, sqlstr, cs.Oid, cs.ID, cs.Name, cs.Definitiontext, cs.Status, cs.Version, cs.Versiondescription, cs.Assigningauthorityversionname, cs.Distributionsourceversionname, cs.Distributionsourceid, cs.Assigningauthorityid, cs.Codesystemcode, cs.Sourceurl, cs.Hl70396identifier, cs.Legacyflag, cs.Statusdate, cs.Acquireddate, cs.Effectivedate, cs.Expirydate, cs.Assigningauthorityreleasedate, cs.Distributionsourcereleasedate, cs.Sdocreatedate, cs.Lastrevisiondate); err != nil {
		return logerror(err)
	}
	// set exists
	cs._exists = true
	return nil
}

// Delete deletes the [CodeSystem] from the database.
func (cs *CodeSystem) Delete(ctx context.Context, db DB) error {
	switch {
	case !cs._exists: // doesn't exist
		return nil
	case cs._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.code_system ` +
		`WHERE oid = $1`
	// run
	logf(sqlstr, cs.Oid)
	if _, err := db.ExecContext(ctx, sqlstr, cs.Oid); err != nil {
		return logerror(err)
	}
	// set deleted
	cs._deleted = true
	return nil
}

// CodeSystemByID retrieves a row from 'public.code_system' as a [CodeSystem].
//
// Generated from index 'code_system_id_key'.
func CodeSystemByID(ctx context.Context, db DB, id string) (*CodeSystem, error) {
	// query
	const sqlstr = `SELECT ` +
		`oid, id, name, definitiontext, status, version, versiondescription, assigningauthorityversionname, distributionsourceversionname, distributionsourceid, assigningauthorityid, codesystemcode, sourceurl, hl70396identifier, legacyflag, statusdate, acquireddate, effectivedate, expirydate, assigningauthorityreleasedate, distributionsourcereleasedate, sdocreatedate, lastrevisiondate ` +
		`FROM public.code_system ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	cs := CodeSystem{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&cs.Oid, &cs.ID, &cs.Name, &cs.Definitiontext, &cs.Status, &cs.Version, &cs.Versiondescription, &cs.Assigningauthorityversionname, &cs.Distributionsourceversionname, &cs.Distributionsourceid, &cs.Assigningauthorityid, &cs.Codesystemcode, &cs.Sourceurl, &cs.Hl70396identifier, &cs.Legacyflag, &cs.Statusdate, &cs.Acquireddate, &cs.Effectivedate, &cs.Expirydate, &cs.Assigningauthorityreleasedate, &cs.Distributionsourcereleasedate, &cs.Sdocreatedate, &cs.Lastrevisiondate); err != nil {
		return nil, logerror(err)
	}
	return &cs, nil
}

// CodeSystemByOid retrieves a row from 'public.code_system' as a [CodeSystem].
//
// Generated from index 'code_system_pkey'.
func CodeSystemByOid(ctx context.Context, db DB, oid string) (*CodeSystem, error) {
	// query
	const sqlstr = `SELECT ` +
		`oid, id, name, definitiontext, status, version, versiondescription, assigningauthorityversionname, distributionsourceversionname, distributionsourceid, assigningauthorityid, codesystemcode, sourceurl, hl70396identifier, legacyflag, statusdate, acquireddate, effectivedate, expirydate, assigningauthorityreleasedate, distributionsourcereleasedate, sdocreatedate, lastrevisiondate ` +
		`FROM public.code_system ` +
		`WHERE oid = $1`
	// run
	logf(sqlstr, oid)
	cs := CodeSystem{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, oid).Scan(&cs.Oid, &cs.ID, &cs.Name, &cs.Definitiontext, &cs.Status, &cs.Version, &cs.Versiondescription, &cs.Assigningauthorityversionname, &cs.Distributionsourceversionname, &cs.Distributionsourceid, &cs.Assigningauthorityid, &cs.Codesystemcode, &cs.Sourceurl, &cs.Hl70396identifier, &cs.Legacyflag, &cs.Statusdate, &cs.Acquireddate, &cs.Effectivedate, &cs.Expirydate, &cs.Assigningauthorityreleasedate, &cs.Distributionsourcereleasedate, &cs.Sdocreatedate, &cs.Lastrevisiondate); err != nil {
		return nil, logerror(err)
	}
	return &cs, nil
}
