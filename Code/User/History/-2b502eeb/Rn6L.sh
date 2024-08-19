git checkout --orphan new_start
rm -rf *
y
touch README.md
git add -A
git commit -m "reset repo"
git branch -D master
git branch -m master
git push -f origin master