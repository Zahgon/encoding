package iso8601

// ValidFlags is a bitset type used to configure the behavior of the Valid
// function.
type ValidFlags int

const (
	// Strict is a validation flag used to represent a string iso8601 validation
	// (this is the default).
	Strict ValidFlags = 0

	// AllowSpaceSeparator allows the presence of a space instead of a 'T' as
	// separator between the date and time.
	AllowSpaceSeparator ValidFlags = 1 << iota

	// AllowMissingTime allows the value to contain only a date.
	AllowMissingTime

	// AllowMissingSubsecond allows the value to contain only a date and time.
	AllowMissingSubsecond

	// AllowMissingTimezone allows the value to be missing the timezone
	// information.
	AllowMissingTimezone

	// AllowNumericTimezone allows the value to represent timezones in their
	// numeric form.
	AllowNumericTimezone

	// Flexible is a combination of all validation flag that allow for
	// non-strict checking of the input value.
	Flexible = AllowSpaceSeparator | AllowMissingTime | AllowMissingSubsecond | AllowMissingTimezone | AllowNumericTimezone
)

// Valid check value to verify whether or not it is a valid iso8601 time
// representation.
func Valid(value string, flags ValidFlags) bool {
	_ = "STUB: not implemented"

	// year
	return false
}

// month

// day

// date only

// separator

// hour

// minute

// second

// microsecond

// date and time

// timezone

// timezone hour

// timezone minute

func readDigits(value string, min, max int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func readByte(value string, c byte) (string, bool) { _ = "STUB: not implemented"; return "", false }

func isDigit(c byte) bool { _ = "STUB: not implemented"; return false }
