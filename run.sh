killall tbRank
chmod 755 tbRank
nohup ./tbRank start > log/tbRank_`date +%Y%m%d%H%M%S`.log 2>&1 &
