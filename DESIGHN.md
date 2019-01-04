# アプリ名: osake

# 概要
お酒の小売店が商品の在庫管理を行う、という想定の商品管理アプリケーション

# 機能一覧
* 管理者ログイン／ログアウト機能
* 商品検索／一覧表示機能
* 商品追加機能
* 商品詳細更新機能
* 商品削除機能

# 画面遷移
![画面遷移図](https://imgur.com/a/Mo8EjSj)
# 内部設計
![内部設計図](https://imgur.com/a/guZkCg6)

# 機能詳細
## 管理者ログイン機能
* ログインに成功した際、セッションクッキー（"auth"）を生成し、ログインIDを登録
* 各ハンドラでクッキーを確認。クッキーが未取得またはログインIDが空文字の時に、ログイン画面にリダイレクトす
* クッキーはブラウザを閉じた際に削除される

## 商品検索／一覧表示機能
* 検索文字列 ⇔ 商品名 から商品を検索し情報を表示

# データテーブル
## 管理者テーブル（SUPER_VISOR）
| No | 論理名 | 物理名 | データ型(桁数) | PK | NOT NULL | 備考 |
| :- | :- | :- | :- | :- | :- | :- |
| 1 | 管理者ID | super_visor_id | VARCHAR(30) | 〇 | | |
| 2 | パスワード | super_visor_password | VARCHAR(20) | | 〇 | |
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


# アーキテクチャ構成
* go による Webサーバ
* postgres による DB サーバ
## Web サーバ

## DB サーバ
