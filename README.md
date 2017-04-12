# git_prompt_useremail
> zsh plugin to add prompt reminders for git user.email

Do you git commit under multiple email addresses depending on the situation?
(For example different addresses for work and open source projects.) Tend to
forget which one you have active at a given time, leading to having to amend
commits to fix it?  This is for you.

I personally use this to have an emoji in my prompt to remind me which email
will be recorded when I git commit.

## Usage

The shell function `git_prompt_useremail` will echo the email address your
next commit will be attributed to (checking both the `$GIT_AUTHOR_EMAIL`
environment variable and the local `git-config` cascade for `user.email`.)

More importantly, the `git_prompt_useremail_symbol` function will echo to you a
handy user definable emoji/icon/unicode-glyph for quick recognizeability!

To define the symbols, add entries to your git configuration (`~/.gitconfig` or
also see [git-config] for details of the many options there), with the format:

    [emailprompt "bwayne73@hotmail.com"]
        prompt = 🥂
    [emailprompt "bruce@wayne-enterprises.biz"]
        prompt = 👔
    [emailprompt "batman@justice-league.usa.gov"]
        prompt = 🦇

I personally like to put this one in my `$RPROMPT`.

    export RPROMPT='$(git_prompt_useremail_symbol) '

_Note that extra space at the end makes emoji in the RPROMPT look better in the
macOS terminal. The single quotation marks are important though, so that the
execution happens every time $RPROMPT is evaluated rather than when you first
set it._

<img alt="screenshot" 
     src="https://cloud.githubusercontent.com/assets/40650/23048542/63ea7ec2-f46b-11e6-8b7c-dc7102911feb.png"
     width="430" />

## Installation

TODO

[git-config]: https://git-scm.com/docs/git-config
