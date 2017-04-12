
__retrieve_email_symbol () {
    git config --get emailprompt."$1".prompt
}

git_prompt_useremail () {
    if git rev-parse --is-inside-work-tree > /dev/null 2>&1; then
        echo "${GIT_AUTHOR_EMAIL:-$(command git config --get user.email 2>/dev/null)}"
    fi
}

git_prompt_useremail_symbol () {
    __email="$(git_prompt_useremail)"
    if [ -n "$__email" ]; then
        __symbol=$(__retrieve_email_symbol "$__email")
        if [ -n "$__symbol" ]; then
            echo "$__symbol"
        else
            echo "❔"
        fi
    fi
}
