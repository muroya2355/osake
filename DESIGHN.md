# アプリ名: osake

# 概要
お酒の小売店が商品の在庫管理を行う、という想定の商品管理アプリケーション

# インストール方法
1. Docker または Docker Toolbox をインストール、起動する

2. git リポジトリをクローン、作業フォルダに移動
```
$ git clone https://github.com/muroya2355/osake.git
$ cd osake
$ ls
```
&emsp;&emsp;&nbsp; docker-compose.yml があることを確認

3. Docker イメージをビルド
```
$ docker-compose build
$ docker image ls
```
&emsp;&emsp;&nbsp; osake_osake, osake_osake-dev, osake_postgres, osake_httpd が作成されていることを確認

4. Docker コンテナを起動
```
$ docker-compose up
```
&emsp;&emsp;&nbsp; 別のターミナルを立ち上げて
```
$ docker ps -a
```
&emsp;&emsp;&nbsp; osake_* イメージに対応する、4つのコンテナが起動 (Up) していることを確認

5. ブラウザからコンテナにアクセス
- Docker をインストールした場合\
→ http://localhost:10080/
- Docker Toolbox をインストールした場合 \
→ ` $ docker-machile ls ` で仮想マシンのIPアドレスを確認（例:192.168.99.100）\
→ http://192.168.99.100:8080/

# 機能一覧
* 管理者ログイン／ログアウト機能
* 商品検索／一覧表示機能
* 商品追加機能
* 商品詳細更新機能
* 商品削除機能

# データテーブル
## 管理者テーブル（SUPER_VISOR）
| No | 論理名 | 物理名 | データ型(桁数) | PK | NOT NULL | 備考 |
| :- | :- | :- | :- | :- | :- | :- |
| 1 | 管理者ID | super_visor_id | VARCHAR(30) | 〇 | | |
| 2 | パスワード | super_visor_password | VARCHAR(20) | | 〇 | |

## 商品（GOODS）
| No | 論理名 | 物理名 | データ型(桁数) | PK | NOT NULL | 備考 |
| :- | :- | :- | :- | :- | :- | :- |
| 1 | 商品ID | goods_id | BIGINT | 〇 | | 固有の管理番号 |
| 2 | 商品名 | goods_name | VARCHAR(100) | | 〇 | |
| 3 | 分類ID | class_id   | INTEGER | | 〇 | |
| 4 | メーカーID | maker_id | INTEGER | | 〇 | |
| 5 | 販売価格 | indicated_price | INTEGER | | | |
| 6 | 仕入れ値 | purchase_price | NUMERIC | | | |
| 7 | 在庫 | stock | INTEGER | | | |
| 8 | 更新者ID | update_super_visor_id | VARCHAR(30) | | 〇 | |
| 9 | 更新日時 | update_date | TIMESTAMP | | 〇 | |
| 10 | 更新バージョンID | update_version_id | BIGINT | | 〇 |

## 分類（CLASS）
| No | 論理名 | 物理名 | データ型(桁数) | PK | NOT NULL | 備考 |
| :- | :- | :- | :- | :- | :- | :- |
| 1 | 分類ID | class_id | BIGINT | 〇 | | 固有の管理番号 |
| 2 | 分類名 | class_name | VARCHAR(100) | | 〇 | |

## メーカー（MAKER）
| No | 論理名 | 物理名 | データ型(桁数) | PK | NOT NULL | 備考 |
| :- | :- | :- | :- | :- | :- | :- |
| 1 | メーカーID | class_id | BIGINT | 〇 | | 固有の管理番号 |
| 2 | メーカー名 | class_name | VARCHAR(100) | | 〇 | |

# 画面遷移
![画面遷移図](./images/osake.png)
# 内部設計
![内部設計図](./images/osake_detail.png)

# ディレクトリ階層
```
./go
│
├─main.go
│
├─controller
│    ├─addgoods.go
│    ├─authenticate.go
│    ├─deletegoods.go
│    ├─goodsdetail.go
│    ├─inputgoods.go
│    └─updategoods.go
│
├─model
│    ├─class.go
│    ├─goods.go
│    ├─maker.go
│    └─user.go
│
├─view
│     ├─goodsdetail.html
│     ├─goodslist.html
│     ├─inputgoods.html
│     └─login.html
│
└─utils
     └─utils.go
```

# 機能詳細
## 管理者ログイン機能
* ログインに成功した際、セッションクッキー（"auth"）を生成し、ログインIDを登録
* 各ハンドラでクッキーを確認する。
* クッキーが未取得またはログインIDが空文字の時に、ログイン画面にリダイレクトする
* クッキーはブラウザを閉じた際に削除される

## 商品検索／一覧表示機能
* 検索文字列 ⇔ 商品名 から商品を検索し情報を表示


# アーキテクチャ構成

Dockerコンテナを使って ↓ を作成する
* Apache によるリバースプロキシ
* goバイナリ による Webサーバ（開発環境・実行環境）
* postgreSQL による DB サーバ

![コンテナ構成図](./images/docker.png)

## Apache

## golang

## PostgreSQL
