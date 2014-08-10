package media

import(
  "testing"
)

func TestNewResourceURL(t *testing.T) {
  expected := "http://example.com/movie.mp4"
  actual := NewResource(expected).URL()

  if expected != actual {
    t.Fatalf("invalid URL. Want %s but got %s", expected, actual)
  }
}
