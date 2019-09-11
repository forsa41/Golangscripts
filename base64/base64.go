package main

import (
  "fmt"
  "flag"
  "os"
  "io/ioutil"
  "unicode"
  "regexp"
  "strings"
  "encoding/base64"
)

var asciikarakter = &unicode.RangeTable{    //https://golang.org/pkg/unicode/#RangeTable
  R16: []unicode.Range16{
    {0x0020, 0x007F, 1},
  },

}

func main() {
  flag.Parse()

  dosya := flag.Arg(0)

  if dosya == ""{ // dosya kullanimi
    fmt.Fprintln(os.Stderr, "kullanim: base64 <dosya adi>")
    return
  }

  f, err := os.Open(dosya)
  if err != nil {   //Hata kontrolu
    fmt.Fprintf(os.Stderr, "%s\n", err)
    return
  }

  b, err := ioutil.ReadAll(f)
  if err != nil {   // dosya ici okunmasi
    fmt.Fprintf(os.Stderr, "%s\n", err)
    return
  }

  icerik := string(b)

  re := regexp.MustCompile("[A-Za-z0-9+/][a-zA-z0-9+/]+={0,2}")
  esles := re.FindAllString(icerik, -1)

  if esles == nil {
    return
    }
    for _, m := range esles {

      if len(m) < 7 {
        continue
      }

      if (len(m)-1)%4 != 0{
        continue
        }

        decb, _ := base64.StdEncoding.DecodeString(m[1:])
        decoded := string(decb)
        if decoded == ""{
          continue
        }

        decoded = strings.Replace(decoded, "\n", " ", -1)

        nonascii := false

        for _, r := range decoded {

          if !unicode.Is(asciikarakter, r) {
            nonascii = true
            break
          }
         }
         if nonascii {
           continue
         }
         fmt.Println(decoded)
  }

}
