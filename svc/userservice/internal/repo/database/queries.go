package database

// INSERT
const (
	insertUserLoginCred = `
		INSERT INTO USER_CRED(EMAIL,PASSWORD) VALUES ($1,$2) RETURNING ID;	
`
	insertUserInfo = `
		INSERT INTO USER_INFO(ID,EMAIL,FULL_NAME,PHONE_NUMBER) VALUES ($1,$2,$3,$4) RETURNING ID;
`
)

// UPDATE
const (
	updateUserInfo = `
		UPDATE USER_INFO
		SET EMAIL = $1,
			FULL_NAME = $2,
			PHONE_NUMBER = $3
		WHERE ID = $4;
`
)

// DELETE
const (
	deleteUserInfo = `
		DELETE FROM USER_INFO
		WHERE ID = $1;
`
)

// SELECT
const (
	getUserCredByEmail = `
		SELECT ID,EMAIL,PASSWORD FROM USER_CRED WHERE EMAIL = $1;
`
	getUserInfoByEmail = `
		SELECT ID,EMAIL,FULL_NAME,PHONE_NUMBER FROM USER_INFO WHERE EMAIL = $1;
`
)
