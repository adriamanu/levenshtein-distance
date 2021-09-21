package levenshtein_distance

import (
	"log"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	t.Run("similar words requires no modification", func(t *testing.T) {
		a := "maison"
		b := "maison"
		lev := Levenshtein(a, b)
		if lev != 0 {
			log.Fatal("levenshtein distance should be equal to 0 when words are identical")
		}
	})

	t.Run("small modifications for words that look alike", func(t *testing.T) {
		a := "pull"
		b := "push"
		lev := Levenshtein(a, b)
		if lev != 2 {
			log.Fatal("2 modifications are required to go from 'pull' to 'push'")
		}
	})

	t.Run("fully modify a word", func(t *testing.T) {
		a := "zèbre"
		b := "anathème"
		lev := Levenshtein(a, b)
		if lev != 7 {
			log.Fatal("levenshtein distance should be equal to 7 while replacing 'zèbre' by 'anathème'")
		}
	})
}
