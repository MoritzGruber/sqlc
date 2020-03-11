package catalog

import (
	"errors"

	"github.com/kyleconroy/sqlc/internal/sql/ast"
	sqlerr "github.com/kyleconroy/sqlc/internal/sql/errors"
)

func (c *Catalog) alterTable(stmt *ast.AlterTableStmt) error {
	var implemented bool
	for _, item := range stmt.Cmds.Items {
		switch cmd := item.(type) {
		case *ast.AlterTableCmd:
			switch cmd.Subtype {
			case ast.AT_AddColumn:
				implemented = true
			case ast.AT_AlterColumnType:
				implemented = true
			case ast.AT_DropColumn:
				implemented = true
			case ast.AT_DropNotNull:
				implemented = true
			case ast.AT_SetNotNull:
				implemented = true
			}
		}
	}
	if !implemented {
		return nil
	}
	_, table, err := c.getTable(stmt.Table)
	if err != nil {
		return err
	}

	for _, cmd := range stmt.Cmds.Items {
		switch cmd := cmd.(type) {
		case *ast.AlterTableCmd:
			idx := -1

			// Lookup column names for column-related commands
			switch cmd.Subtype {
			case ast.AT_AlterColumnType,
				ast.AT_DropColumn,
				ast.AT_DropNotNull,
				ast.AT_SetNotNull:
				for i, c := range table.Columns {
					if c.Name == *cmd.Name {
						idx = i
						break
					}
				}
				if idx < 0 && !cmd.MissingOk {
					return sqlerr.ColumnNotFound(table.Rel.Name, *cmd.Name)
				}
				// If a missing column is allowed, skip this command
				if idx < 0 && cmd.MissingOk {
					continue
				}
			}

			switch cmd.Subtype {

			case ast.AT_AddColumn:
				for _, c := range table.Columns {
					if c.Name == cmd.Def.Colname {
						return sqlerr.ColumnExists(table.Rel.Name, c.Name)
					}
				}
				table.Columns = append(table.Columns, &Column{
					Name:      cmd.Def.Colname,
					Type:      *cmd.Def.TypeName,
					IsNotNull: cmd.Def.IsNotNull,
				})

			case ast.AT_AlterColumnType:
				table.Columns[idx].Type = *cmd.Def.TypeName
				// table.Columns[idx].IsArray = isArray(d.TypeName)

			case ast.AT_DropColumn:
				table.Columns = append(table.Columns[:idx], table.Columns[idx+1:]...)

			case ast.AT_DropNotNull:
				table.Columns[idx].IsNotNull = false

			case ast.AT_SetNotNull:
				table.Columns[idx].IsNotNull = true

			}
		}
	}

	return nil
}
func (c *Catalog) createTable(stmt *ast.CreateTableStmt) error {
	ns := stmt.Name.Schema
	if ns == "" {
		ns = c.DefaultSchema
	}
	schema, err := c.getSchema(ns)
	if err != nil {
		return err
	}
	if _, _, err := schema.getTable(stmt.Name); err != nil {
		if !errors.Is(err, sqlerr.NotFound) {
			return err
		}
	} else if stmt.IfNotExists {
		return nil
	}
	tbl := Table{Rel: stmt.Name}
	for _, col := range stmt.Cols {
		tbl.Columns = append(tbl.Columns, &Column{
			Name:      col.Colname,
			Type:      *col.TypeName,
			IsNotNull: col.IsNotNull,
		})
	}
	schema.Tables = append(schema.Tables, &tbl)
	return nil
}

func (c *Catalog) dropTable(stmt *ast.DropTableStmt) error {
	for _, name := range stmt.Tables {
		ns := name.Schema
		if ns == "" {
			ns = c.DefaultSchema
		}
		schema, err := c.getSchema(ns)
		if errors.Is(err, sqlerr.NotFound) && stmt.IfExists {
			continue
		} else if err != nil {
			return err
		}

		_, idx, err := schema.getTable(name)
		if errors.Is(err, sqlerr.NotFound) && stmt.IfExists {
			continue
		} else if err != nil {
			return err
		}

		schema.Tables = append(schema.Tables[:idx], schema.Tables[idx+1:]...)
	}
	return nil
}