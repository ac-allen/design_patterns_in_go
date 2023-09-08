package main

import (
  "fmt"
  "strings"
  "unicode"
)

type TextRange struct {
  Start, End int
  Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
  return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
  plainText string
  formatting []*TextRange
}

func (b *BetterFormattedText) String() string {
  sb := strings.Builder{}

  for i := 0; i < len(b.plainText); i++ {
    c := b.plainText[i]
    for _, r := range b.formatting {
      if r.Covers(i) && r.Capitalize {
        c = uint8(unicode.ToUpper(rune(c)))
      }
    }
    sb.WriteRune(rune(c))
  }

  return sb.String()
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
  return &BetterFormattedText{plainText: plainText}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
  r := &TextRange{start, end, false, false, false}
  b.formatting = append(b.formatting, r)
  return r
}



func main() {
  text := "This is a brave new world"

  bft := NewBetterFormattedText(text)
  bft.Range(16, 19).Capitalize = true // new
  fmt.Println(bft.String())
}