/*
Based on https://github.com/ged/linguistics/blob/master/lib/linguistics/en/articles.rb

Copyright (c) 2003-20011, Michael Granger
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice,
  this list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the author/s, nor the names of the project's
  contributors may be used to endorse or promote products derived from this
  software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package english

import (
	"regexp"
)

var regexpNewlinePlusIndentation = regexp.MustCompile(`\n\s*`)

// This pattern codes the beginnings of all english words begining with a
// 'y' followed by a consonant. Any other y-consonant prefix therefore
// implies an abbreviation.
var A_y_cons = regexp.MustCompile(`(?i)^(y(?:b[lor]|cl[ea]|fere|gg|p[ios]|rou|tt))`)

// Exceptions to exceptions
var A_explicit_an = []*regexp.Regexp{
	regexp.MustCompile("(?i)euler"),
	regexp.MustCompile("(?i)houri?"),
	regexp.MustCompile("(?i)heir"),
	regexp.MustCompile("(?i)honest"),
	regexp.MustCompile("(?i)hono"),
}

// Returns the given word with a prepended indefinite article, unless
// +count+ is non-nil and not singular.
func IndefiniteArticleFor(word string) string {
	// Handle special cases
	for _, regex := range A_explicit_an {
		if regex.MatchString(word) {
			return "an"
		}
	}

	// Handle abbreviations (disabled because Go regexps don't have lookahead)
	// if A_abbrev.MatchString(word) {
	//	return "an"
	// }

	// "an A-frame"
	if regexp.MustCompile(`(?i)^[aefhilmnorsx][.-]`).MatchString(word) {
		return "an"
	}

	// "a G-string"
	if regexp.MustCompile(`(?i)^[a-z][.-]`).MatchString(word) {
		return "a"
	}

	// Handle consonants
	if regexp.MustCompile(`(?i)^[^aeiouy]`).MatchString(word) {
		return "a"
	}

	// Handle special vowel-forms
	if regexp.MustCompile(`(?i)^e[uw]`).MatchString(word) ||
		regexp.MustCompile(`(?i)^onc?e\b`).MatchString(word) ||
		regexp.MustCompile(`(?i)^uni([^nmd]|mo)`).MatchString(word) ||
		regexp.MustCompile(`(?i)^u[bcfhjkqrst][aeiou]`).MatchString(word) {
		return "a"
	}

	// Handle vowels
	if regexp.MustCompile(`(?i)^[aeiou]`).MatchString(word) {
		return "an"
	}

	// Handle y...
	// (before certain consonants implies (unnaturalized) "i.." sound)
	if A_y_cons.MatchString(word) {
		return "an"
	}

	// Otherwise, guess "a"
	return "a"
}
