package _535_encode_decode_url

import "strconv"

type Codec struct {
	hashMap map[string]string
	id      int
}

func Constructor() Codec {
	return Codec{
		hashMap: make(map[string]string),
		id:      0,
	}
}

// Encodes a URL to a shortened URL.
func (cc *Codec) encode(longUrl string) string {
	cc.id++
	shortUrl := "http://tinyurl.com/" + strconv.Itoa(cc.id)

	cc.hashMap[shortUrl] = longUrl

	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (cc *Codec) decode(shortUrl string) string {
	return cc.hashMap[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
