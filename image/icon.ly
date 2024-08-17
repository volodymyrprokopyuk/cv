\version "2.25.18"

% lilypond -f svg icon.ly

\paper {
  #(set-paper-size '(cons (* 8 mm) (* 8 mm)))
  top-margin = -1.5
  bottom-margin = 0
  left-margin = 0
  right-margin = 0
  oddFooterMarkup = ##f
}

\markup \serif \bold \magnify #1.4 \rotate #30 CV
