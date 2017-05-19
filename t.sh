# maybe more powerful
# for mac (sed for linux is different)
dir=`echo ${PWD##*/}`
grep "weixin-x" * -R | grep -v Godeps | awk -F: '{print $1}' | sort | uniq | xargs sed -i '' "s#weixin-x#$dir#g"
mv weixin-x.ini $dir.ini

