# git bash
# PS1="\[\e[37;40m\][\[\e[32;40m\]termux:\[\e[35;40m\]\W\[\e[0m\]]\\$ "

export LANG="en_US.UTF-8"


# ls
alias ls='ls --color=auto'
alias l='ls -CF'
alias la='ls -a'
alias ll='ls -l'
alias lsa='ls -la'

# elegant
alias dk='docker'
alias dka='docker exec -it'
alias dir='ll'
alias pss='ps -a|grep'

# disk
for i in {a..z}
do
    alias $i:='cd /'$i
    alias ${i^}:='cd /'${i^}
done

