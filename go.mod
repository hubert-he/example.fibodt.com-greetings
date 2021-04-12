module example.fibodt.com/greetings

go 1.15

replace (
	example.fibodt.com/greetings => ../greetings
	example.fibodt.com/translate => github.com/hubert-he/translate v0.0.0-20210412071231-a42662b5a9f1
)

require (
	example.fibodt.com/translate v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.3
)
