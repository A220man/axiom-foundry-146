// Package release builds localized release notes from a single bilingual source.
//
// Upstream issue #2260: a release went out with English-only notes because the
// Chinese section was missing below a `<!-- lang:zh -->` marker. Apps set to
// Chinese then showed English notes. This package parses a notes document that
// contains an English section followed by a `<!-- lang:zh -->` marker and a
// Chinese section, and returns the section for the requested language.
package release

import "strings"

// langMarker separates the English section from the Chinese section.
const langMarker = "<!-- lang:zh -->"

// NoteLanguages lists the language tags this package understands, in display order.
var NoteLanguages = []string{"en", "zh"}

// ParseNotes splits a bilingual notes document into per-language sections.
//
// The document is expected to contain an English section, then the
// `<!-- lang:zh -->` marker, then a Chinese section. Text before the marker is
// returned under "en"; text after the marker is returned under "zh". If the
// marker is absent, the whole document is returned under "en" and "zh" is
// empty, mirroring the upstream failure mode so callers can detect it.
func ParseNotes(doc string) map[string]string {
	idx := strings.Index(doc, langMarker)
	if idx < 0 {
		return map[string]string{
			"en": strings.TrimSpace(doc),
			"zh": "",
		}
	}
	en := strings.TrimSpace(doc[:idx])
	zh := strings.TrimSpace(doc[idx+len(langMarker):])
	return map[string]string{"en": en, "zh": zh}
}

// NotesFor returns the notes for the given language tag.
//
// Unknown tags fall back to English. An empty Chinese section is returned as-is
// (empty string) so callers can warn rather than silently showing English.
func NotesFor(doc, lang string) string {
	sections := ParseNotes(doc)
	switch strings.ToLower(strings.TrimSpace(lang)) {
	case "zh":
		return sections["zh"]
	default:
		return sections["en"]
	}
}

// HasChinese reports whether the document contains a non-empty Chinese section.
func HasChinese(doc string) bool {
	return ParseNotes(doc)["zh"] != ""
}
