package models

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.

type SnippetModel struct {
	DB *pgx.Conn
}

// This will insert a new snippet into the database.

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {

	// Write the SQL statement we want to execute. I've split it over two lines
	// for readability (which is why it's surrounded with backquotes instead
	// of normal double quotes).
	stmt := `INSERT INTO scrolls (title, content, created, expires) 
	VALUES($1, $2,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP + $3 * INTERVAL '1 DAY') RETURNING id`

	var id int

	// Use the Exec() method on the embedded connection pool to execute the
	// statement. The first parameter is the SQL statement, followed by the
	// title, content and expiry values for the placeholder parameters. This
	// method returns a sql.Result type, which contains some basic
	// information about what happened when the statement was executed.
	err := m.DB.QueryRow(context.Background(), stmt, title, content, expires).Scan(&id)

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// This will return a specific snippet based on its id.

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	// Write the SQL statement we want to execute. Again, I've split it over two
	// lines for readability.
	stmt := `SELECT id, title, content, created, expires FROM scrolls
	WHERE expires > CURRENT_TIMESTAMP AND id = $1`
	// Use the QueryRow() method on the connection pool to execute our
	// SQL statement, passing in the untrusted id variable as the value for the
	// placeholder parameter. This returns a pointer to a sql.Row object which
	// holds the result from the database.
	row := m.DB.QueryRow(context.Background(), stmt, id)
	// Initialize a pointer to a new zeroed Snippet struct.
	s := &Snippet{}
	// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Snippet struct. Notice that the arguments
	// to row.Scan are *pointers* to the place you want to copy the data into,
	// and the number of arguments must be exactly the same as the number of
	// columns returned by your statement.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a
		// sql.ErrNoRows error. We use the errors.Is() function check for that
		// error specifically, and return our own ErrNoRecord error
		// instead (we'll create this in a moment).
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	// If everything went OK then return the Snippet object.
	return s, nil
}

// This will return the 10 most recently created snippets.

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
