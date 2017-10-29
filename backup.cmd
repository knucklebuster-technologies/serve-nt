git add --all
git commit -m "setting changes into history"
git push
restic backup -r G:\RESTIC_BACKUP\ C:\work\go\src\github.com\qawarrior\serve-nt\
restic forget -r G:\RESTIC_BACKUP\ --keep-last 2 --prune