bin/git-prompt-useremail-symbol: *.go go.mod go.sum
	mkdir -p bin
	go build -o bin/git-prompt-useremail-symbol .

bench: bin/git-prompt-useremail-symbol
	hyperfine --warmup 3 'source git-prompt-useremail.plugin.zsh; eval git_prompt_useremail_symbol' 'bin/git-prompt-useremail-symbol'
