package lib

import (
	"testing"

	"github.com/Jemmic/go-pcre2"
	u "github.com/sunshine69/golang-tools/utils"
)

func TestFindStringSubmatch(t *testing.T) {
	input := `This is sample phone: 12345`
	println(u.JsonDump(u.Must(Pcre2FindStringSubmatch(`phone: ([\d]+)`, input)), ""))
	ptn := pcre2.MustCompile(`phone: ([\d]+)`, 0)
	u.Assert(u.Must(Pcre2FindStringSubmatch(ptn, input))[1] == "12345", "OK", true)
}

func TestPcre2FindAllSubmatch(t *testing.T) {
	input := `This is sample phone: 12345 and phone: 56789`
	println(u.JsonDump(u.Must(Pcre2FindAllStringSubmatch(`phone: ([\d]+)`, input)), ""))
	ptn := pcre2.MustCompile(`phone: ([\d]+)`, 0)
	println(u.JsonDump(u.Must(Pcre2FindAllStringSubmatch(ptn, input)), ""))
	ob := u.Must(Pcre2FindAllSubmatch(`phone: ([\d]+)`, []byte(input)))
	os := [][]string{}
	for _, it := range ob {
		os = append(os, (u.SliceMap(it, func(b []byte) *string { s := string(b); return &s })))
	}
	println(u.JsonDump(os, ""))

}
