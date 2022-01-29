package palindromic_test

import (
	"testing"

	"leetcode/palindromic"

	"github.com/stretchr/testify/assert"
)

func TestFirstPalindrome(t *testing.T) {
	t.Parallel()

	words := []string{
		"knywrurkwbrtpalvuuzbczcwzpdqibcwwyflwiddixemsrwblupyerjgvcpavdjfhmujitcsmdbvhxw",
		"ovkeowhqvwlndtpxdnimgietvjsqydbuuwmxkfxxgn",
		"damomwtjugmsrfyvytaheg",
		"bngqatscosdakdwjz",
		"jwzcowuantgqlzjrzgpapcugxvviltrhmcwijtpqapmxjfcilrsmsbeffphcxmaozlczncoxxjmuqijhidxqinhywrtah",
		"ujvoejixvaioikkwzxgtmkchckrigfejjpheiiehpjjefgirkchckmtgxzwkkioiavxijeovju",
		"kacjvcouuigbhydrryaperxvjetwsycmnlkxujaqngjhhotqskclquklxsozfryfxwiksstmropcdvhgsnosgvltqo",
		"qrbsdxxolwzmyltproznfgyydxkkejwdiwpvfzvjoxqvwguoerhclytzvolymbxummwsoqtttyttik",
		"tkekt",
		"esrshrlfoihhjrouleucwpeubygivoatrfraphgwpvtkanwu",
		"kwzrfelynncvsrwvaukiinhjdydmlimggjldhflygemzwnjizzlsuqwahsufwmwhfjkjpngdfsudyavtogoaqzknpew",
		"sdgpcnvsbzxhyjt",
	}

	if !assert.Equal(t,
		"ujvoejixvaioikkwzxgtmkchckrigfejjpheiiehpjjefgirkchckmtgxzwkkioiavxijeovju",
		palindromic.FirstPalindrome(words)) {
		t.Fail()
	}
}

func BenchmarkFirstPalindrome10000(b *testing.B) {
	words := []string{
		"knywrurkwbrtpalvuuzbczcwzpdqibcwwyflwiddixemsrwblupyerjgvcpavdjfhmujitcsmdbvhxw",
		"ovkeowhqvwlndtpxdnimgietvjsqydbuuwmxkfxxgn",
		"damomwtjugmsrfyvytaheg",
		"bngqatscosdakdwjz",
		"jwzcowuantgqlzjrzgpapcugxvviltrhmcwijtpqapmxjfcilrsmsbeffphcxmaozlczncoxxjmuqijhidxqinhywrtah",
		"ujvoejixvaioikkwzxgtmkchckrigfejjpheiiehpjjefgirkchckmtgxzwkkioiavxijeovju",
		"kacjvcouuigbhydrryaperxvjetwsycmnlkxujaqngjhhotqskclquklxsozfryfxwiksstmropcdvhgsnosgvltqo",
		"qrbsdxxolwzmyltproznfgyydxkkejwdiwpvfzvjoxqvwguoerhclytzvolymbxummwsoqtttyttik",
		"tkekt",
		"esrshrlfoihhjrouleucwpeubygivoatrfraphgwpvtkanwu",
		"kwzrfelynncvsrwvaukiinhjdydmlimggjldhflygemzwnjizzlsuqwahsufwmwhfjkjpngdfsudyavtogoaqzknpew",
		"sdgpcnvsbzxhyjt",
	}

	for i := 0; i < b.N; i++ {
		_ = palindromic.FirstPalindrome(words)
	}
}
