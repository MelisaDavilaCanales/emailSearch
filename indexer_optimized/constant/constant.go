package constant

const (
	// DATE_FORMAT is the format used for parsing dates in the emails.
	DATE_FORMAT = "Mon, _2 Jan 2006 15:04:05 -0700 (MST)"

	//
	DATE_FORMAT_TO_LOGS = "2006-01-02 15:04:05"

	// TAG_CONTENT_REGEX is the regex used to clean some name fields.
	TAG_CONTENT_REGEX = "<[^>]*>"

	// TAG_CONTENT_REGEX it's the regex that identifies the atrribute href in the html tags.
	HREF_REGEXEP = `\s*href="[^"]*"`

	// IMPORTATNT_REGEXP is the regex used to clean the !important attribute in the css.
	IMPORTATNT_REGEXP = `\s*!important`

	// PREFIXES_AND_SYMBOLS_REGEXP  is the regex used to clean unwanted Chars e.g. e-mail, <email.>, <., <'.' and >
	PREFIXES_AND_SYMBOLS_REGEXP = `(?i)e-mail|<email.>|<\.\s*|<'?'\s*|\s*>`
)
