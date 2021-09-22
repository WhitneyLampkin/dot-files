
alias kube="kubectl"

alias pbcopy='xclip -selection clipboard'
alias pbpaste='xclip -selection clipboard -o'

alias swag-edit='docker run --rm -d -p 9000:8080 swaggerapi/swagger-editor'

alias commit-template='echo "
Subject: *Short title for change with Ticket Number*
Why was this change made:
*1-2 sentences for why this change was made*	
What issues does this change address:
- *Issue*: *describe fix*
Are there any side effects or issues to resolve:
- *Problem*
"'

DATE=$(date +%D)
alias snippet-template='echo "#snippet $DATE - $DATE
 
#achievements
- 
 
AKS SRE:
- 

Infra team:
- 

Dataplane team:
- 

Overlay team:
- 

CXPSRE:
- 

Other:
- 

"'


function urldecode() { : "${*//+/ }"; echo -e "${_//%/\\x}"; }

decode_jwt(){
     printf "%s" "$1" | base64 -d 2>/dev/null | jq .
}

cover () {
    t="/tmp/go-cover.$$.tmp"
    go test -coverprofile=$t $@ && go tool cover -html=$t && unlink $t
}

alias "list-taints"="kubectl get nodes -o json | jq '.items[].spec.taints'"

alias "go-main"="cat \"$HOME/template/main.go\""
