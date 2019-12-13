# LearningGitAndGo
这是一个学习go和git的仓库，里面会有一些关于go语言学习的小项目和笔记。
2019.12.13：今天在push文件到远程仓库时，碰到了rejected错误。通过翻阅资料，发现是由于本地仓库不是最新的（必须和远程仓库版本同步，github上的README.md文件没有下载下来）。通过在git中输入git pull --rebase origin master将最新版本拉到本地，这样就可以从本地上传文件到远程了
