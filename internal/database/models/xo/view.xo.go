package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// View represents a row from 'public.view'.
type View struct {
	Name            string         `json:"name"`            // name
	ID              string         `json:"id"`              // id
	Descriptiontext string         `json:"descriptiontext"` // descriptiontext
	Status          string         `json:"status"`          // status
	Viewnotes       sql.NullString `json:"viewnotes"`       // viewnotes
	Statusdate      time.Time      `json:"statusdate"`      // statusdate
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [View] exists in the database.
func (v *View) Exists() bool {
	return v._exists
}

// Deleted returns true when the [View] has been marked for deletion
// from the database.
func (v *View) Deleted() bool {
	return v._deleted
}

// Insert inserts the [View] to the database.
func (v *View) Insert(ctx context.Context, db DB) error {
	switch {
	case v._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case v._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.view (` +
		`name, id, descriptiontext, status, viewnotes, statusdate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)`
	// run
	logf(sqlstr, v.Name, v.ID, v.Descriptiontext, v.Status, v.Viewnotes, v.Statusdate)
	if _, err := db.ExecContext(ctx, sqlstr, v.Name, v.ID, v.Descriptiontext, v.Status, v.Viewnotes, v.Statusdate); err != nil {
		return logerror(err)
	}
	// set exists
	v._exists = true
	return nil
}

// Update updates a [View] in the database.
func (v *View) Update(ctx context.Context, db DB) error {
	switch {
	case !v._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case v._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.view SET ` +
		`name = $1, descriptiontext = $2, status = $3, viewnotes = $4, statusdate = $5 ` +
		`WHERE id = $6`
	// run
	logf(sqlstr, v.Name, v.Descriptiontext, v.Status, v.Viewnotes, v.Statusdate, v.ID)
	if _, err := db.ExecContext(ctx, sqlstr, v.Name, v.Descriptiontext, v.Status, v.Viewnotes, v.Statusdate, v.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [View] to the database.
func (v *View) Save(ctx context.Context, db DB) error {
	if v.Exists() {
		return v.Update(ctx, db)
	}
	return v.Insert(ctx, db)
}

// Upsert performs an upsert for [View].
func (v *View) Upsert(ctx context.Context, db DB) error {
	switch {
	case v._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.view (` +
		`name, id, descriptiontext, status, viewnotes, statusdate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, descriptiontext = EXCLUDED.descriptiontext, status = EXCLUDED.status, viewnotes = EXCLUDED.viewnotes, statusdate = EXCLUDED.statusdate `
	// run
	logf(sqlstr, v.Name, v.ID, v.Descriptiontext, v.Status, v.Viewnotes, v.Statusdate)
	if _, err := db.ExecContext(ctx, sqlstr, v.Name, v.ID, v.Descriptiontext, v.Status, v.Viewnotes, v.Statusdate); err != nil {
		return logerror(err)
	}
	// set exists
	v._exists = true
	return nil
}

// Delete deletes the [View] from the database.
func (v *View) Delete(ctx context.Context, db DB) error {
	switch {
	case !v._exists: // doesn't exist
		return nil
	case v._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.view ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, v.ID)
	if _, err := db.ExecContext(ctx, sqlstr, v.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	v._deleted = true
	return nil
}

// ViewByID retrieves a row from 'public.view' as a [View].
//
// Generated from index 'view_pkey'.
func ViewByID(ctx context.Context, db DB, id string) (*View, error) {
	// query
	const sqlstr = `SELECT ` +
		`name, id, descriptiontext, status, viewnotes, statusdate ` +
		`FROM public.view ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	v := View{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&v.Name, &v.ID, &v.Descriptiontext, &v.Status, &v.Viewnotes, &v.Statusdate); err != nil {
		return nil, logerror(err)
	}
	return &v, nil
}
