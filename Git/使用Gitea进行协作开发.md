# 使用Gitea进行协作开发

## Git和Gitea的关系

SVN:

![](http://www.liaoxuefeng.com/files/attachments/001384860735706fd4c70aa2ce24b45a8ade85109b0222b000/0)

Git

![](http://www.liaoxuefeng.com/files/attachments/0013848607465969378d7e6d5e6452d8161cf472f835523000/0)




Git服务器

![](http://on51si7u9.bkt.clouddn.com/Git.png)


## 进行团队协作主要使用的功能:

- 组织

- 团队

- 分支

  - Git flow

  - 分支合并与冲突

- issue(非Git原生功能)

- Pull Request(PR,Merge)


## Review

## 持续集成

- hooks

## Git flow

- ### 长期分支:

    - master branch

    - develop branch

- ### 短期分支:

    - feature branch
    - hotfix branch

## 分支合并与冲突

创建分支的使用就是为了最后的合并.

合并的过程中分支之间很有可能存在冲突的情况导致合并不了.


## Gitea备份功能

Gitea服务基于Git,但脱离于Git文件.


代码文件的安全性不取决于Gitea,取决于存放git文件的文件系统.

> /opt/sunrun/data/gitea/repo


gitea dump导出数据库数据,并压缩Git文件.


安全性的关键:定期将Git文件做冗余,存放于多个位置.

