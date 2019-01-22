# アプリ名: osake

# 概要
お酒の小売店が商品の在庫管理を行う、という想定の商品管理アプリケーション

https://www.osaketen.com

ID : a , パスワード : password1 でログインできる

ローカルマシンへのインストール方法は local-dev ブランチを参照のこと
# 機能一覧
* 管理者ログイン／ログアウト機能
* 商品検索／一覧表示機能
* 商品追加機能
* 商品詳細更新機能
* 商品削除機能

# 画面遷移
<img src="./images/osake.png" width=70%>

# アーキテクチャ構成

## Web3層

Dockerコンテナを使って ↓ を作成する
* Apache によるリバースプロキシ
* goバイナリ による Webサーバ（開発環境・実行環境）
* postgreSQL による DB サーバ

![コンテナ構成図](./images/docker.png)


## ログ分析基盤

Dockerコンテナを使って ↓ を作成する
* ログ収集：fluentd
* ログ格納：elasticsearch
* 可視化：Kibana

![ログ分析基盤構成図](./images/logger.png)