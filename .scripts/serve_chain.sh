echo ' ::: SERVING LAVA CHAIN (ignite) :::'
(cd $LAVA && ignite chain serve -v -r) 2>&1 | grep -e lava_ -e ERR_ -e STARPORT] -e !
# (cd /go/lava && starport chain serve -v -r) 2>&1 | grep -e lava_ -e ERR_ -e STARPORT] -e !
