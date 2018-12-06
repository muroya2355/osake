# 構成
* go による Webサーバ
* postgres による DB サーバ
## Web サーバ

## DB サーバ
### postgres の設定
* postgres 10.6
* インストール先：C:\Program Files\PostgreSQL\10
* データ保存先：C:\Program Files\PostgreSQL\10\data
* スーパーユーザ／パスワード：postgres／postgres
* ポート：5432

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
| 5 | 型番 | model_number | VARCHAR(30) | | 〇 | |
| 6 | スペック | specs | TEXT |　| | |
| 7 | 販売価格 | indicated_price | INTEGER | | | |
| 8 | 仕入れ値 | purchase_price | NUMERIC | | | |
| 9 | 在庫 | stock | INTEGER | | | |
| 10 | 削除フラグ | is_deleted | BOOLEAN | | 〇 | |
| 11 | 更新者ID | update_super_visor_id | VARCHAR(30) | | 〇 | |
| 12 | 更新日時 | update_date | TIMESTAMP | | 〇 | |
| 13 | 更新バージョンID | update_version_id | BIGINT | | 〇 |


# 機能
## 機能一覧
* 管理者ログイン機能
* 商品検索／一覧表示機能
* 商品追加機能
* 商品詳細更新機能

## ログイン機能
1. ブラウザから http://localhost:8080/login.html にアクセスしてログイン画面を表示する
1. ログイン画面でユーザIDとパスワードを入力し、ログインボタンを押下
1. 入力されたユーザID を基に SUPER_VISOR テーブルを検索
1. パスワードが正当か確認
1. 正当な場合はログイン成功画面を表示、失敗の場合は再度ログイン画面を表示

* login.html -> POST /authenticate -> loginsuccessful.html

## 商品検索／一覧表示機能

## 商品追加機能

## 商品詳細更新機能

