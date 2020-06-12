package backtrack

import (
	"testing"
)

func TestRExp_Match(t *testing.T) {
	trues := map[string][]string{
		"ab": {"ab"},
		"?ab": {"eab", "xab"},
		"a?b": {"ab", "aeb", "alb"},
		"ab?": {"ab", "abd", "abe"},
		"??ab": {"ab", "sab", "exab"},
		"a??b": {"ab", "aeb", "axb", "asfb", "aexb"},
		"ab??": {"ab", "abd", "abpd"},
		"*ab": {"ab", "dab", "dfab", "31ab"},
		"a*b": {"ab", "adb", "addabfrwbb"},
		"ab*": {"ab", "abd", "abfasfd"},
	}

	falses := map[string][]string{
		"ab": {"ad"},
		"?ab": {"adab", "acab"},
		"a?b": {"ad", "aedb", "aldb"},
		"ab?": {"ad", "abdd", "abed"},
		"??ab": {"ad", "sadab", "exabd"},
		"a??b": {"ad", "gadb", "axdbd", "asfbd", "aexba"},
		"ab??": {"ad", "abdde", "aqbpq"},
		"*ab": {"ad", "dadb", "dfadadb", "31aqpb"},
		"a*b": {"ad", "adqerbd", "adrwbre"},
		"ab*": {"ad", "adbddfa", "adbfasfde"},
	}

	gotTrue := func(pattern string, texts []string) {
		r := NewRegularExpresstion(pattern)
		for _, text := range texts {
			if b := r.Match(text); !b {
				t.Errorf("expect: true, got: false. pattern: %s, text: %s", pattern, text)
			}
		}
	}

	gotFalse := func(pattern string, texts []string) {
		r := NewRegularExpresstion(pattern)
		for _, text := range texts {
			if b := r.Match(text); b {
				t.Errorf("expect: false, got: true. pattern: %s, text: %s", pattern, text)
			}
		}
	}

	for pattern, texts := range trues {
		gotTrue(pattern, texts)
	}

	for pattern, texts := range falses {
		gotFalse(pattern, texts)
	}
}
