package _535_encode_decode_url

import "testing"

func TestEncodeDecode(t *testing.T) {
	codec := Constructor()
	longURL := "https://leetcode.com/problems/design-tinyurl"

	shortURL := codec.encode(longURL)

	if got := codec.decode(shortURL); got != longURL {
		t.Fatalf(
			"decode(%q) = %q; want %q",
			shortURL,
			got,
			longURL,
		)
	}
}

func TestEncodeReturnsSequentialUniqueURLs(t *testing.T) {
	codec := Constructor()

	first := codec.encode("https://example.com/first")
	second := codec.encode("https://example.com/second")
	third := codec.encode("https://example.com/third")

	if first != "http://tinyurl.com/1" {
		t.Errorf(
			"first short URL = %q; want %q",
			first,
			"http://tinyurl.com/1",
		)
	}

	if second != "http://tinyurl.com/2" {
		t.Errorf(
			"second short URL = %q; want %q",
			second,
			"http://tinyurl.com/2",
		)
	}

	if third != "http://tinyurl.com/3" {
		t.Errorf(
			"third short URL = %q; want %q",
			third,
			"http://tinyurl.com/3",
		)
	}
}

func TestSeveralURLsAreDecodedIndependently(t *testing.T) {
	codec := Constructor()

	longURLs := []string{
		"https://example.com",
		"https://example.com/path?q=go#section",
		"http://localhost:8080/users/42",
	}

	shortURLs := make([]string, len(longURLs))

	for i, longURL := range longURLs {
		shortURLs[i] = codec.encode(longURL)
	}

	for i, shortURL := range shortURLs {
		if got := codec.decode(shortURL); got != longURLs[i] {
			t.Errorf(
				"decode(%q) = %q; want %q",
				shortURL,
				got,
				longURLs[i],
			)
		}
	}
}

func TestEncodingSameURLCreatesNewShortURL(t *testing.T) {
	codec := Constructor()
	longURL := "https://example.com/repeated"

	first := codec.encode(longURL)
	second := codec.encode(longURL)

	if first == second {
		t.Fatalf(
			"two encode calls returned the same URL %q",
			first,
		)
	}

	if got := codec.decode(first); got != longURL {
		t.Errorf(
			"decode(first) = %q; want %q",
			got,
			longURL,
		)
	}

	if got := codec.decode(second); got != longURL {
		t.Errorf(
			"decode(second) = %q; want %q",
			got,
			longURL,
		)
	}
}

func TestDecodeUnknownURL(t *testing.T) {
	codec := Constructor()

	if got := codec.decode("http://tinyurl.com/999"); got != "" {
		t.Errorf(
			"decode(unknown URL) = %q; want an empty string",
			got,
		)
	}
}

func TestCodecInstancesHaveIndependentState(t *testing.T) {
	firstCodec := Constructor()
	secondCodec := Constructor()

	firstURL := firstCodec.encode("https://example.com/a")
	secondURL := secondCodec.encode("https://example.com/b")

	if firstURL != "http://tinyurl.com/1" ||
		secondURL != "http://tinyurl.com/1" {
		t.Fatalf(
			"independent codecs returned %q and %q; both must start with id 1",
			firstURL,
			secondURL,
		)
	}

	if got := firstCodec.decode(firstURL); got != "https://example.com/a" {
		t.Errorf("first codec returned %q", got)
	}

	if got := secondCodec.decode(secondURL); got != "https://example.com/b" {
		t.Errorf("second codec returned %q", got)
	}
}
