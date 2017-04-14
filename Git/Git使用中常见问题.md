# Git使用中常见问题

1. 添加远程仓库时提示fatal: remote origin already exists

    原因: 存在已关联的远程仓库.

    解决方法:

    > git remote rm origin

2. 提示没有配置用户名和邮箱:

    解决方法:

    > git config --global user.email "email@example.com"

    > git config --global user.name "example_name"


3. 合并分支时的冲突:

    待续...