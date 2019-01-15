package crypto

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	"github.com/c3systems/go-substrate/common/u8util"
)

func TestNewSHA256(t *testing.T) {
	for i, tt := range []struct {
		in  []byte
		out string
	}{
		{[]byte(""), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{[]byte("abc"), "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{[]byte("hello"), "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := NewSHA256(tt.in)
			if hex.EncodeToString(result[:]) != tt.out {
				t.Errorf("want %v; got %v", tt.out, hex.EncodeToString(result[:]))
			}
		})
	}
}

func TestNewBlake2b256(t *testing.T) {
	for i, tt := range []struct {
		in  []byte
		out string
	}{
		{[]byte(""), "0e5751c026e543b2e8ab2eb06099daa1d1e5df47778f7787faab45cdf12fe3a8"},
		{[]byte("abc"), "bddd813c634239723171ef3fee98579b94964e3bb1cb3e427262c8c068d52319"},
		{[]byte("hello"), "324dcf027dd4a30a932c441f365a25e86b173defa4b8e58948253471b81b72cf"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := NewBlake2b256(tt.in)
			if hex.EncodeToString(result[:]) != tt.out {
				t.Errorf("want %v; got %v", tt.out, hex.EncodeToString(result[:]))
			}
		})
	}
}

func TestNewBlake2b512(t *testing.T) {
	for i, tt := range []struct {
		in  []byte
		out string
	}{
		{[]byte(""), "786a02f742015903c6c6fd852552d272912f4740e15847618a86e217f71f5419d25e1031afee585313896444934eb04b903a685b1448b755d56f701afe9be2ce"},
		{[]byte("abc"), "ba80a53f981c4d0d6a2797b69f12f6e94c212f14685ac4b74b12bb6fdbffa2d17d87c5392aab792dc252d5de4533cc9518d38aa8dbf1925ab92386edd4009923"},
		{[]byte("hello"), "e4cfa39a3d37be31c59609e807970799caa68a19bfaa15135f165085e01d41a65ba1e1b146aeb6bd0092b49eac214c103ccfa3a365954bbbe52f74a2b3620c94"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := NewBlake2b512(tt.in)
			if hex.EncodeToString(result[:]) != tt.out {
				t.Errorf("want %v; got %v", tt.out, hex.EncodeToString(result[:]))
			}
		})
	}
}

func TestNewBlake2b256Sig(t *testing.T) {
	for i, tt := range []struct {
		key  []byte
		data []byte
		out  string
	}{
		{nil, []byte("abc"), "bddd813c…68d52319"},
		{[]byte{4, 5, 6}, []byte{1, 2, 3}, "af0e60f4…8a714a8f"},
		{[]byte("abc"), nil, "7a78f945…73b8f07b"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := NewBlake2b256Sig(tt.key, tt.data)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(u8util.ToHex(result[:], 512/8, false), tt.out) {
				t.Error(u8util.ToHex(result[:], 512/8, false), tt.out)
			}
		})
	}
}

func TestNewBlake2b512Sig(t *testing.T) {
	for i, tt := range []struct {
		key  []byte
		data []byte
		out  string
	}{
		{nil, []byte("abc"), "ba80a53f…d4009923"},
		{[]byte{4, 5, 6}, []byte{1, 2, 3}, "7ed2dd42…2ec9362e"},
		{[]byte("abc"), nil, "91cc35fc…f47b00e5"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := NewBlake2b512Sig(tt.key, tt.data)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(u8util.ToHex(result[:], 512/8, false), tt.out) {
				t.Error(u8util.ToHex(result[:], 512/8, false), tt.out)
			}
		})
	}
}
