module github.com/openHacking/hello

go 1.17

replace github.com/openHacking/search => ../search

require github.com/openHacking/search v0.0.0-00010101000000-000000000000

require github.com/kljensen/snowball v0.6.0 // indirect
