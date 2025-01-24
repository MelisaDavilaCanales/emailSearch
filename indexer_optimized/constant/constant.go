package constant

const (
	// DATE_FORMAT is the format used for parsing dates in the emails.
	DATE_FORMAT = "Mon, _2 Jan 2006 15:04:05 -0700 (MST)"

	//
	DATE_FORMAT_TO_LOGS = "2006-01-02 15:04:05"

	// EMAIL_REGEX is the regex used to clean some name fields.
	TAG_CONTENT_REGEX = "<[^>]*>"

	// EMAIL_WORD_REGEX is the regex used to clean some email fields.
	EMAIL_WORD_REGEX = "e-mail"

	// CHAR_REGEX is the regex used to clean some email fields.
	CHAR_REGEX = "<."
)
