package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
)

func TestMonth(t *testing.T) {
	// current is Friday
	fixt := []Fixture{
		// past/last
		{"это нужно было сделать 3 марта", 42, "3 март", 1368 * time.Hour},
		{"напиши мне двадцатого декабря, договоримся", 20, "двадцатого декабря", 8376 * time.Hour},
		{"тридцать первое декабря", 0, "тридцать первое декабря", 8640 * time.Hour},
		{"июнь", 0, "июнь", 3648 * time.Hour},
		{"июня", 0, "июня", 3648 * time.Hour},
		{"4 января", 0, "4 января", -48 * time.Hour},
		{"7е марта", 0, "7е март", 1464 * time.Hour},
		{"1го сентября", 0, "1го сентября", 5736 * time.Hour},
	}

	w := when.New(nil)

	w.Add(ru.Month(rules.Override))

	ApplyFixtures(t, "ru.Month", w, fixt)
}
