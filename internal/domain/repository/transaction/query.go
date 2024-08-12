package transactionrepo

var (
	insertNewTransaction = `
		INSERT INTO credit_transaction(
			contract_number,
			asset_name,
			consumer_id,
			partner_id,
			otr,
			admin_fee,
			total_installment,
			interest,
			created_by,
			updated_by
		) VALUES (
			:contract_number,
			:asset_name,
			:consumer_id,
			:partner_id,
			:otr,
			:admin_fee,
			:total_installment,
			:interest,
			:partner_id,
			:partner_id
		)
	`
	insertNewPayment = `
		INSERT INTO credit_payment(
			transaction_id,
			consumer_id,
			total_amount,
			monthly_amount,
			duration,
			interest,
			status,
			created_by,
			updated_by
		) VALUES (
			:transaction_id,
			:consumer_id,
			:total_amount,
			:monthly_amount,
			:duration,
			:interest,
			:status,
			:created_by,
			:updated_by
		)
	`

	insertNewPaymentInstallment = `
		INSERT INTO credit_payment_installment(
			payment_id,
			amount,
			status,
			created_by,
			updated_by
		) VALUES (
			:payment_id,
			:amount,
			:status,
			:created_by,
			:updated_by
		)
	`
)

var (
	getLimit = `
		SELECT
			id,
			consumer_id,
			duration,
			amount
		FROM
			credit_limit
		WHERE
			consumer_id = ?
			AND duration = ?
			AND deleted_at IS NULL
		LIMIT 1
	`

	getOngoingPayment = `
		SELECT
			id,
			transaction_id,
			consumer_id,
			total_amount,
			monthly_amount,
			duration,
			interest,
			status
		FROM
			credit_payment
		WHERE
			consumer_id = ?
			AND duration = ?
			AND status = ?
			AND deleted_at IS NULL
	`

	getOngoingPaymentInstallment = `
		SELECT
			id,
			payment_id,
			amount,
			status
		FROM
			credit_payment_installment
		WHERE
			status = :status
			AND payment_id IN (:payment_id)
			AND deleted_at IS NULL
	`
)
