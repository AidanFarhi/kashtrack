#!/bin/bash

sqlite3 db/expense.db < db/sql/001-tables.sql
# sqlite3 db/expense.db < db/sql/002-seed.sql
# sqlite3 db/expense.db < db/sql/003-validate.sql
