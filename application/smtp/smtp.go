package smtp

type OrderSMTP interface {
	SendEmail(
		recipient, subject, message string,
	) error
}
