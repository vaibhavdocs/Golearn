package main

import "fmt"

/*

Traditionally we start with ASCII way of representing the English letters, symbols and controls character
It used 7 bits to represent a particular character, so we could have 128 representations.
But going ahead it faced a challenge where every country did not accept the this scheme since ASCII
only English letters. Others countries came up with different representation where it resembled ASCII till 127
and added their own character.

As you can already sense that every other non-English speaking and writing country will have its
own scheme/representation/standard.

After all this, Unicode came into existing, created by Go creators. Unicode was same as ASCII till 127
but after that it started incorporating all the letters, symbols and other characters in the world

ASCII code occupied maximum of 7 bits, but Unicode tool bytes into consideration, to be specific representation took
place in 1-4 bytes

ASCII code are represeted by 0xxxxxxx, where 0 says that it is an ASCII code

110xxxxx 10xxxxxx, higher bits in the first byte(Left to Right) tells the system that there are 2 bytes
Basically these invariable higher bits work as Headers, so you always know that there are exactly 2 bytes to be considered
or simply, 2 bytes make the character

110xxxxx 10xxxxxx - same theories for these numbers as well, where Headers tell How many bytes does this
10xxxxxx			character have ? 3

110xxxxx 10xxxxxx -  Ditto, 4 bytes
10xxxxxx 10xxxxxx

Now comes the UTF-8, which is nothing but encoding of Unicode Code points




Go's implementation


Reader (p []byte) (i int, e error)

Above implementation readers byte from the file and if the byte is lesser than 128 than it is easier to
decide since this it is an ASCII and take only a byte
But what if read byte is 198 or 165 then Reader does not know how may more bytes are required to
be read to form the character

Bufio
Reader(p []byte) (r rune,i int, e error)

File ---- [] byte of size N read from the file ---- DecodeRune() takes bytes from the []bytes and checks
whether code is a valid UTF-8 code or not if not then it is incremented by one byte and runes are read again if they
are valid UTF-8 codes, this goes on till 4 bytes otherwise it is an error (�).

*/
func main() {
	var x rune = 'A'
	fmt.Printf("%U", x)
}
