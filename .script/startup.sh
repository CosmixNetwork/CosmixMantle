source ./.script/environment.sh

assetNode start >~/.cosmixMantle/Node/log &
sleep 10
assetClient rest-server --chain-id "$AM_CHAIN_ID">~/.cosmixMantle/Client/log &
echo "
Node and Client started up. For logs:
tail -f ~/.cosmixMantle/Node/log
tail -f ~/.cosmixMantle/Client/log
"
