package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var textEng = `In Acid for the Children, Flea takes readers on a deeply
	personal and revealing tour of his formative years, spanning from
	Australia to the New York City suburbs to, finally, Los Angeles.
	Through hilarious anecdotes, poetical meditations, and occasional
	flights of fantasy, Flea deftly chronicles the experiences that forged
	him as an artist, a musician, and a young man. His dreamy, jazz-inflected
	prose makes the Los Angeles of the 1970s and 80s come to gritty, glorious
	life, including the potential for fun, danger, mayhem, or inspiration that
	lurked around every corner. It is here that young Flea, looking to escape
	a turbulent home, found family in a community of musicians, artists, and
	junkies who also lived on the fringe. He spent most of his time partying
	and committing petty crimes. But it was in music where he found a higher
	meaning, a place to channel his frustration, loneliness, and love. This
	left him open to the life-changing moment when he and his best friends,
	soul brothers, and partners-in-mischief came up with the idea to start
	their own band, which became the Red Hot Chili Peppers.`

var textReduced = `Клетки.
	Вас держат в коробке?
	Клетки.
	Связаны.
	Связаны.`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})

	t.Run("positive test of english text", func(t *testing.T) {
		expected := []string{
			"the",  // 10
			"and",  // 8
			"a",    // 6
			"to",   // 6
			"his",  // 5
			"of",   // 5
			"flea", // 3
			"he",   // 3
			"in",   // 2
			"that", // 2
		}
		require.Equal(t, expected, Top10(textEng))
	})
	t.Run("positive test of reduced text", func(t *testing.T) {
		expected := []string{
			"клетки",  // 2
			"связаны", // 2
			"в",       // 1
			"вас",     // 1
			"держат",  // 1
			"коробке", // 1
		}
		require.Equal(t, expected, Top10(textReduced))
	})
}
