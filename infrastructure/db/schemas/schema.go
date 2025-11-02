package schemas

const CreateTableStmt = `
	CREATE TABLE IF NOT EXISTS transactions (
		id VARCHAR(32) PRIMARY KEY,
		status VARCHAR(12) NOT NULL
	)
`

const InsertTransactionStmt = `
	INSERT INTO transactions (id, status)
	VALUES (?, ?)
`

const GetTransactionByIDStmt = `
	SELECT id, status
		FROM transactions
		WHERE id = ?
`
const UpdateTransactionStatusByIdStmt = `
	UPDATE transactions
	SET status = ? 
	WHERE id = ?
`
