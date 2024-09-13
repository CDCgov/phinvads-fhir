package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// CodeSystemConcept represents a row from 'public.code_system_concept'.
type CodeSystemConcept struct {
	ID                              string         `json:"id"`                              // id
	Name                            string         `json:"name"`                            // name
	Codesystemoid                   string         `json:"codesystemoid"`                   // codesystemoid
	Conceptcode                     string         `json:"conceptcode"`                     // conceptcode
	Sdopreferreddesignation         string         `json:"sdopreferreddesignation"`         // sdopreferreddesignation
	Definitiontext                  sql.NullString `json:"definitiontext"`                  // definitiontext
	Precoordinatedflag              sql.NullBool   `json:"precoordinatedflag"`              // precoordinatedflag
	Status                          string         `json:"status"`                          // status
	Sdoconceptstatus                sql.NullString `json:"sdoconceptstatus"`                // sdoconceptstatus
	Supersededbycodesystemconceptid sql.NullString `json:"supersededbycodesystemconceptid"` // supersededbycodesystemconceptid
	Umlscui                         sql.NullString `json:"umlscui"`                         // umlscui
	Umlsaui                         sql.NullString `json:"umlsaui"`                         // umlsaui
	Isrootflag                      bool           `json:"isrootflag"`                      // isrootflag
	Isconceptflag                   bool           `json:"isconceptflag"`                   // isconceptflag
	Sdoconceptcreateddate           sql.NullTime   `json:"sdoconceptcreateddate"`           // sdoconceptcreateddate
	Sdoconceptrevisiondate          sql.NullTime   `json:"sdoconceptrevisiondate"`          // sdoconceptrevisiondate
	Statusdate                      time.Time      `json:"statusdate"`                      // statusdate
	Sdoconceptstatusdate            sql.NullTime   `json:"sdoconceptstatusdate"`            // sdoconceptstatusdate
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [CodeSystemConcept] exists in the database.
func (csc *CodeSystemConcept) Exists() bool {
	return csc._exists
}

// Deleted returns true when the [CodeSystemConcept] has been marked for deletion
// from the database.
func (csc *CodeSystemConcept) Deleted() bool {
	return csc._deleted
}

// Insert inserts the [CodeSystemConcept] to the database.
func (csc *CodeSystemConcept) Insert(ctx context.Context, db DB) error {
	switch {
	case csc._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case csc._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.code_system_concept (` +
		`id, name, codesystemoid, conceptcode, sdopreferreddesignation, definitiontext, precoordinatedflag, status, sdoconceptstatus, supersededbycodesystemconceptid, umlscui, umlsaui, isrootflag, isconceptflag, sdoconceptcreateddate, sdoconceptrevisiondate, statusdate, sdoconceptstatusdate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`)`
	// run
	logf(sqlstr, csc.ID, csc.Name, csc.Codesystemoid, csc.Conceptcode, csc.Sdopreferreddesignation, csc.Definitiontext, csc.Precoordinatedflag, csc.Status, csc.Sdoconceptstatus, csc.Supersededbycodesystemconceptid, csc.Umlscui, csc.Umlsaui, csc.Isrootflag, csc.Isconceptflag, csc.Sdoconceptcreateddate, csc.Sdoconceptrevisiondate, csc.Statusdate, csc.Sdoconceptstatusdate)
	if _, err := db.ExecContext(ctx, sqlstr, csc.ID, csc.Name, csc.Codesystemoid, csc.Conceptcode, csc.Sdopreferreddesignation, csc.Definitiontext, csc.Precoordinatedflag, csc.Status, csc.Sdoconceptstatus, csc.Supersededbycodesystemconceptid, csc.Umlscui, csc.Umlsaui, csc.Isrootflag, csc.Isconceptflag, csc.Sdoconceptcreateddate, csc.Sdoconceptrevisiondate, csc.Statusdate, csc.Sdoconceptstatusdate); err != nil {
		return logerror(err)
	}
	// set exists
	csc._exists = true
	return nil
}

// Update updates a [CodeSystemConcept] in the database.
func (csc *CodeSystemConcept) Update(ctx context.Context, db DB) error {
	switch {
	case !csc._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case csc._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.code_system_concept SET ` +
		`name = $1, codesystemoid = $2, conceptcode = $3, sdopreferreddesignation = $4, definitiontext = $5, precoordinatedflag = $6, status = $7, sdoconceptstatus = $8, supersededbycodesystemconceptid = $9, umlscui = $10, umlsaui = $11, isrootflag = $12, isconceptflag = $13, sdoconceptcreateddate = $14, sdoconceptrevisiondate = $15, statusdate = $16, sdoconceptstatusdate = $17 ` +
		`WHERE id = $18`
	// run
	logf(sqlstr, csc.Name, csc.Codesystemoid, csc.Conceptcode, csc.Sdopreferreddesignation, csc.Definitiontext, csc.Precoordinatedflag, csc.Status, csc.Sdoconceptstatus, csc.Supersededbycodesystemconceptid, csc.Umlscui, csc.Umlsaui, csc.Isrootflag, csc.Isconceptflag, csc.Sdoconceptcreateddate, csc.Sdoconceptrevisiondate, csc.Statusdate, csc.Sdoconceptstatusdate, csc.ID)
	if _, err := db.ExecContext(ctx, sqlstr, csc.Name, csc.Codesystemoid, csc.Conceptcode, csc.Sdopreferreddesignation, csc.Definitiontext, csc.Precoordinatedflag, csc.Status, csc.Sdoconceptstatus, csc.Supersededbycodesystemconceptid, csc.Umlscui, csc.Umlsaui, csc.Isrootflag, csc.Isconceptflag, csc.Sdoconceptcreateddate, csc.Sdoconceptrevisiondate, csc.Statusdate, csc.Sdoconceptstatusdate, csc.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [CodeSystemConcept] to the database.
func (csc *CodeSystemConcept) Save(ctx context.Context, db DB) error {
	if csc.Exists() {
		return csc.Update(ctx, db)
	}
	return csc.Insert(ctx, db)
}

// Upsert performs an upsert for [CodeSystemConcept].
func (csc *CodeSystemConcept) Upsert(ctx context.Context, db DB) error {
	switch {
	case csc._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.code_system_concept (` +
		`id, name, codesystemoid, conceptcode, sdopreferreddesignation, definitiontext, precoordinatedflag, status, sdoconceptstatus, supersededbycodesystemconceptid, umlscui, umlsaui, isrootflag, isconceptflag, sdoconceptcreateddate, sdoconceptrevisiondate, statusdate, sdoconceptstatusdate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, codesystemoid = EXCLUDED.codesystemoid, conceptcode = EXCLUDED.conceptcode, sdopreferreddesignation = EXCLUDED.sdopreferreddesignation, definitiontext = EXCLUDED.definitiontext, precoordinatedflag = EXCLUDED.precoordinatedflag, status = EXCLUDED.status, sdoconceptstatus = EXCLUDED.sdoconceptstatus, supersededbycodesystemconceptid = EXCLUDED.supersededbycodesystemconceptid, umlscui = EXCLUDED.umlscui, umlsaui = EXCLUDED.umlsaui, isrootflag = EXCLUDED.isrootflag, isconceptflag = EXCLUDED.isconceptflag, sdoconceptcreateddate = EXCLUDED.sdoconceptcreateddate, sdoconceptrevisiondate = EXCLUDED.sdoconceptrevisiondate, statusdate = EXCLUDED.statusdate, sdoconceptstatusdate = EXCLUDED.sdoconceptstatusdate `
	// run
	logf(sqlstr, csc.ID, csc.Name, csc.Codesystemoid, csc.Conceptcode, csc.Sdopreferreddesignation, csc.Definitiontext, csc.Precoordinatedflag, csc.Status, csc.Sdoconceptstatus, csc.Supersededbycodesystemconceptid, csc.Umlscui, csc.Umlsaui, csc.Isrootflag, csc.Isconceptflag, csc.Sdoconceptcreateddate, csc.Sdoconceptrevisiondate, csc.Statusdate, csc.Sdoconceptstatusdate)
	if _, err := db.ExecContext(ctx, sqlstr, csc.ID, csc.Name, csc.Codesystemoid, csc.Conceptcode, csc.Sdopreferreddesignation, csc.Definitiontext, csc.Precoordinatedflag, csc.Status, csc.Sdoconceptstatus, csc.Supersededbycodesystemconceptid, csc.Umlscui, csc.Umlsaui, csc.Isrootflag, csc.Isconceptflag, csc.Sdoconceptcreateddate, csc.Sdoconceptrevisiondate, csc.Statusdate, csc.Sdoconceptstatusdate); err != nil {
		return logerror(err)
	}
	// set exists
	csc._exists = true
	return nil
}

// Delete deletes the [CodeSystemConcept] from the database.
func (csc *CodeSystemConcept) Delete(ctx context.Context, db DB) error {
	switch {
	case !csc._exists: // doesn't exist
		return nil
	case csc._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.code_system_concept ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, csc.ID)
	if _, err := db.ExecContext(ctx, sqlstr, csc.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	csc._deleted = true
	return nil
}

// CodeSystemConceptByCodesystemoid retrieves a row from 'public.code_system_concept' as a [CodeSystemConcept].
//
// Generated from index 'code_system_concept_codesystemoid_idx'.
func CodeSystemConceptByCodesystemoid(ctx context.Context, db DB, codesystemoid string) ([]*CodeSystemConcept, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, codesystemoid, conceptcode, sdopreferreddesignation, definitiontext, precoordinatedflag, status, sdoconceptstatus, supersededbycodesystemconceptid, umlscui, umlsaui, isrootflag, isconceptflag, sdoconceptcreateddate, sdoconceptrevisiondate, statusdate, sdoconceptstatusdate ` +
		`FROM public.code_system_concept ` +
		`WHERE codesystemoid = $1`
	// run
	logf(sqlstr, codesystemoid)
	rows, err := db.QueryContext(ctx, sqlstr, codesystemoid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*CodeSystemConcept
	for rows.Next() {
		csc := CodeSystemConcept{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&csc.ID, &csc.Name, &csc.Codesystemoid, &csc.Conceptcode, &csc.Sdopreferreddesignation, &csc.Definitiontext, &csc.Precoordinatedflag, &csc.Status, &csc.Sdoconceptstatus, &csc.Supersededbycodesystemconceptid, &csc.Umlscui, &csc.Umlsaui, &csc.Isrootflag, &csc.Isconceptflag, &csc.Sdoconceptcreateddate, &csc.Sdoconceptrevisiondate, &csc.Statusdate, &csc.Sdoconceptstatusdate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &csc)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CodeSystemConceptByID retrieves a row from 'public.code_system_concept' as a [CodeSystemConcept].
//
// Generated from index 'code_system_concept_pkey'.
func CodeSystemConceptByID(ctx context.Context, db DB, id string) (*CodeSystemConcept, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, codesystemoid, conceptcode, sdopreferreddesignation, definitiontext, precoordinatedflag, status, sdoconceptstatus, supersededbycodesystemconceptid, umlscui, umlsaui, isrootflag, isconceptflag, sdoconceptcreateddate, sdoconceptrevisiondate, statusdate, sdoconceptstatusdate ` +
		`FROM public.code_system_concept ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	csc := CodeSystemConcept{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&csc.ID, &csc.Name, &csc.Codesystemoid, &csc.Conceptcode, &csc.Sdopreferreddesignation, &csc.Definitiontext, &csc.Precoordinatedflag, &csc.Status, &csc.Sdoconceptstatus, &csc.Supersededbycodesystemconceptid, &csc.Umlscui, &csc.Umlsaui, &csc.Isrootflag, &csc.Isconceptflag, &csc.Sdoconceptcreateddate, &csc.Sdoconceptrevisiondate, &csc.Statusdate, &csc.Sdoconceptstatusdate); err != nil {
		return nil, logerror(err)
	}
	return &csc, nil
}

// CodeSystem returns the CodeSystem associated with the [CodeSystemConcept]'s (Codesystemoid).
//
// Generated from foreign key 'code_system_concept_codesystemoid_fkey'.
func (csc *CodeSystemConcept) CodeSystem(ctx context.Context, db DB) (*CodeSystem, error) {
	return CodeSystemByOid(ctx, db, csc.Codesystemoid)
}
