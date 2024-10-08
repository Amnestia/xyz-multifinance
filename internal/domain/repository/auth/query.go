package authrepo

const (
	auth = `
		SELECT
			id,
			nik,
			nik_index,
			password,
			pin,
			fullname,
			legal_name,
			date_of_birth,
			place_of_birth,
			salary,
			identity_photo,
			photo
		FROM
			consumer
		WHERE
			nik_index = ?
			AND deleted_at IS NULL
		LIMIT 1
	`

	insertNewAccount = `
		INSERT INTO consumer(
			nik,
  			nik_index,
  			password,
  			pin,
  			fullname,
  			legal_name,
  			date_of_birth,
  			place_of_birth,
  			salary,
  			identity_photo,
  			photo,
			created_by,
			updated_by
		) VALUES (
			:nik,
  			:nik_index,
  			:password,
  			:pin,
  			:fullname,
  			:legal_name,
  			:date_of_birth,
  			:place_of_birth,
  			:salary,
  			:identity_photo,
  			:photo,
			:fullname,
			:fullname
		)
	`

	insertNewPartner = `
		INSERT INTO partner(
			name,
			client_id,
			api_key,
			webhook,
			created_by,
			updated_by
		) VALUES (
			:name,
			:client_id,
			:api_key,
			:webhook,
			:name,
			:name
		)
	`

	getPartner = `
		SELECT
			id,
			name,
			client_id,
			api_key,
			webhook
		FROM
			partner
		WHERE
			client_id = ?
			AND deleted_at IS NULL
		LIMIT 1
	`
)
